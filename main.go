package main

import (
	"fmt"
	"news_index/apps"
	testingnews "news_index/testing_news"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db := apps.ConnectToDatabase()
	adminRoute := app.Group("/news")
	adminRoute.Post("/sign-up", func(c *fiber.Ctx) error {
		admin := testingnews.Admins{}
		c.BodyParser(&admin)

		if err := db.Create(&admin).Error; err != nil {
			return c.Status(400).JSON(fiber.Map{"Message: ": "Sign-up Gagal", "Error: ": err})
		}

		response := testingnews.Admins{}
		if err := db.Model(&response).Where("id = ?", admin.ID).Take(&response).Error; err != nil {
			return c.Status(400).JSON(fiber.Map{"Message: ": "Get By Id Gagal", "Error: ": err})
		}
		return c.Status(200).JSON(fiber.Map{
			"Data: ": response,
		})
	})

	adminRoute.Get("/:id", func(c *fiber.Ctx) error {
		adminID := c.Params("id")
		response := testingnews.Admins{}
		if err := db.Model(&response).Preload("Blogs").Preload("Blogs.Category").Where("id = ?", adminID).Take(&response).Error; err != nil {
			return c.Status(400).JSON(fiber.Map{"Message: ": "Get By Id Gagal", "Error: ": err})
		}

		var blogs []testingnews.ResponseBlog
		for _, data := range response.Blogs {
			var blog testingnews.ResponseBlog
			blog.ID = data.ID
			blog.Title = data.Title
			blog.TextBlog = data.TextBlog
			if len(data.Category) > 0 {
				blog.Category = data.Category[0].Category
			}
			blogs = append(blogs, blog)
		}

		adminResp := testingnews.AdminResponse{
			ID:       response.ID,
			Username: response.Username,
			Blogs:    blogs,
		}
		fmt.Print(adminResp)
		return c.Status(200).JSON(fiber.Map{
			"Data: ": adminResp,
		})
	})

	adminRoute.Put("/:id", func(c *fiber.Ctx) error {
		adminID := c.Params("id")
		admin := testingnews.Admins{}
		c.BodyParser(&admin)
		response := testingnews.Admins{}
		if err := db.Model(&response).Where("id = ?", adminID).Joins("Categoy").Take(&response).Error; err != nil {
			return c.Status(400).JSON(fiber.Map{"Message: ": "Get By Id Gagal", "Error: ": err})
		}

		response.Password = admin.Password
		response.Username = admin.Username

		return c.Status(200).JSON(fiber.Map{
			"Data: ": response,
		})
	})
	adminRoute.Delete("/:id", func(c *fiber.Ctx) error {
		adminID := c.Params("id")
		response := testingnews.Admins{}
		if err := db.Model(&response).Where("id = ?", adminID).Delete(&response).Error; err != nil {
			return c.Status(400).JSON(fiber.Map{"Message: ": "Get By Id Gagal", "Error: ": err})
		}
		return c.Status(200).JSON(fiber.Map{
			"Data: ": response,
		})
	})

	caregoriesRoute := app.Group("/news/category")
	caregoriesRoute.Post("/create", func(c *fiber.Ctx) error {
		category := testingnews.Categories{}
		c.BodyParser(&category)
		if err := db.Model(&testingnews.Categories{}).Create(&category).Error; err != nil {
			return c.Status(400).JSON(fiber.Map{"Error: ": err, "Message: ": "Create Gagal"})
		}

		response := testingnews.Categories{}
		if err := db.Model(&response).Where("category = ?", category.Category).Take(&response).Error; err != nil {
			return c.Status(400).JSON(fiber.Map{"Message: ": "Get category Gagal", "Error: ": err})
		}

		
		return c.Status(200).JSON(fiber.Map{"Status: ": "ok", "Data: ": response})
	})

	blogsRoute := app.Group("/news/blogs")
	blogsRoute.Post("/create", func(c *fiber.Ctx) error {
		blog := testingnews.Blogs{}
		c.BodyParser(&blog)
		if err := db.Model(&testingnews.Blogs{}).Create(&blog).Error; err != nil {
			return c.Status(400).JSON(fiber.Map{"Error: ": err, "Message: ": "Create Gagal"})
		}

		if err := db.Table("blogs_categories").Create(map[string]interface{}{
			"blog_id":     blog.ID,
			"category_id": blog.CategoryID,
		}).Error; err != nil {
			return c.Status(400).JSON(fiber.Map{"Error: ": err, "Msg: ": "error create many to many"})
		}

		response := testingnews.Blogs{}
		if err := db.Model(&response).Preload("Category").Take(&response, "id = ?", blog.ID).Error; err != nil {
			return c.Status(400).JSON(fiber.Map{"Error: ": err, "Msg: ": "Error get data many to many"})
		}

		var category string
		if len(response.Category) > 0 {
			category = response.Category[0].Category
		} else {
			category = ""
		}
		return c.Status(200).JSON(fiber.Map{"Data: ": testingnews.ResponseBlog{
			ID:       response.ID,
			Title:    response.Title,
			TextBlog: response.TextBlog,
			Category: category,
		}})

	})

	app.Listen(":3000")
}
