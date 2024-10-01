package main

import (
	init_db "gocommerce/src/database"
	admin_routes "gocommerce/src/routes"
	"log"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	db := init_db.InitDB(host, user, password, dbname, port)

	store := gorm.NewStore(db, true, []byte(os.Getenv("APP_KEY")))

	router.Use(sessions.Sessions("sessions", store))

	router.LoadHTMLGlob("src/templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	admin_routes.AdminRoutes(router)

	router.Run(":8080")
}
