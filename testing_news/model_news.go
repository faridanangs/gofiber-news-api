package testingnews

type Admins struct {
	ID       string  `gorm:"column:id" json:"id"`
	Username string  `gorm:"column:username" json:"username"`
	Password string  `gorm:"column:password" json:"password"`
	Blogs    []Blogs `gorm:"foreignKey:admin_id;references:id"`
}

type Categories struct {
	ID       int    `gorm:"column:id" json:"id"`
	Category string `gorm:"column:category" json:"category"`
}

type Blogs struct {
	ID         string       `gorm:"column:id" json:"id"`
	Title      string       `gorm:"column:title" json:"title"`
	TextBlog   string       `gorm:"column:text_blog" json:"text_blog"`
	AdminID    string       `gorm:"column:admin_id" json:"admin_id"`
	CategoryID int          `gorm:"-" json:"category_id"`
	Category   []Categories `gorm:"many2many:blogs_categories;foreignKey:id;joinForeignKey:blog_id;references:id;joinReferences:category_id"`
}

func (a *Admins) TableName() string {
	return "admins"
}

func (a *Categories) TableName() string {
	return "categories"
}

func (a *Blogs) TableName() string {
	return "blogs"
}

type ResponseBlog struct {
	ID       string
	Title    string
	TextBlog string
	Category string
}

type AdminResponse struct {
	ID       string         `json:"id"`
	Username string         `json:"username"`
	Blogs    []ResponseBlog `json:"blogs"`
}
