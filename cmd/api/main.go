package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/muhammadaghits/golang-todo-restful-api/internal/config"
	"github.com/muhammadaghits/golang-todo-restful-api/internal/database"
	"github.com/muhammadaghits/golang-todo-restful-api/internal/handlers"
	"github.com/muhammadaghits/golang-todo-restful-api/internal/middleware"
)

func main() {
	var cfg *config.Config
	var err error
	cfg, err = config.Load()

	if err != nil {
		log.Fatal("Failed to load configuration: ", err)
	}

	var pool *pgxpool.Pool
	pool, err = database.Connect(cfg.DatabaseURL)

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	defer pool.Close()

	var router *gin.Engine = gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "Todo API is running well!",
			"status":   "Success",
			"database": "Connected",
		})
	})

	router.POST("/api/auth/register", handlers.RegisterHandler(pool))
	router.POST("/api/auth/login", handlers.LoginHandler(pool, cfg))

	protected := router.Group("/api/todos")
	protected.Use(middleware.AuthMiddleware(cfg))
	{
		protected.POST("", handlers.CreateTodoHandler(pool))
		protected.GET("", handlers.GetAllTodoHandler(pool))
		protected.GET("/:id", handlers.GetTodoByIDHandler(pool))
		protected.PUT("/:id", handlers.UpdateTodoHandler(pool))
		protected.DELETE("/:id", handlers.DeleteTodoHandler(pool))
	}

	router.GET("/protected-test", middleware.AuthMiddleware(cfg), handlers.TestProtectedHandler())

	router.Run(":" + cfg.Port)
}
