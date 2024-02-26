package apps

import (
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToDatabase() *gorm.DB {
	// err := godotenv.Load(".env")
	// helpers.PanicIfError(err, "error at load env")
	// dbUserName := os.Getenv("DB_USERNAME")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbPort := os.Getenv("DB_PORT")
	// dbName := os.Getenv("DB_NAME")
	// sslMode := "disable"
	// timeZone := os.Getenv("DB_TIME_ZONE")

	// dns := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", dbUserName, dbPassword, dbName, dbPort, sslMode, timeZone)

	dns := "user=postgres password=anangs port=5432 dbname=news sslmode=disable TimeZone=Asia/Jakarta"

	db, _ := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dns,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxIdleTime(60 * time.Minute)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db
}
