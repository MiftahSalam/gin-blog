package common

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {
	dbPort, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		LogE.Fatal("Error while loading DATABASE_PORT. Err: ", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"), dbPort)
	enableLog, err := strconv.Atoi(os.Getenv("DATABASE_LOGGING"))
	if err != nil {
		LogE.Fatal("Error while loading DATABASE_LOGGING. Err: ", err)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(enableLog)),
	})
	if err != nil {
		LogE.Fatal("db init error: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		LogE.Fatal("get db instance error: ", err)
	}
	sqlDB.SetConnMaxIdleTime(10)

	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
