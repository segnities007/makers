package main

import (
	"fmt"
	"log"
	"os"

	makerHttp "server/logics/http"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
  "server/models"
  "server/logics/db"
)



func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Using environment variables.")
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbCharset := os.Getenv("DB_CHARSET")
	dbParseTime := os.Getenv("DB_PARSE_TIME")
	dbLoc := os.Getenv("DB_LOC")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
		dbCharset,
		dbParseTime,
		dbLoc,
	)

	var dbErr error
	db.DB, dbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if dbErr != nil {
		log.Fatalf("Failed to connect to database: %v", dbErr)
	}

	db.DB.AutoMigrate(
		&models.User{},
		&models.UserInfo{},
	  &models.Group{},
		&models.Image{},
	  &models.Message{},
		&models.DirectMessage{},
		&models.Tag{},
	)

	routing(e)
	e.Logger.Fatal(e.Start(":1323"))
}

func routing(e *echo.Echo){
	makerHttp.POSTs(e)
	makerHttp.GETs(e)
	makerHttp.PUTs(e)
	makerHttp.DELETEs(e)
}
