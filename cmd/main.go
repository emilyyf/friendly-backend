package main

import (
	conn "friendly-backend/internal/db/connection"
	"friendly-backend/internal/handlers"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db, err := conn.OpenConnection()
	if err != nil {
		panic(err)
	}

	conn.RunMigrations(db)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:5173"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.POST("/register", handlers.CreateUserHandler)
	r.POST("/login", handlers.LoginHandler)

	authorized := r.Group("/", handlers.AuthMiddleware())
	{
		authorized.GET("/profile", handlers.ProfileHandler)
	}

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
