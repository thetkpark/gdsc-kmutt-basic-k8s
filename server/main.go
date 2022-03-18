package main

import (
	"fmt"
	"log"
	"os"
	"time"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/caarlos0/env/v6"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/joho/godotenv/autoload"
	"github.com/thetkpark/gdsc-kmutt-basic-k8s/config"
	"github.com/thetkpark/gdsc-kmutt-basic-k8s/todo"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "github.com/thetkpark/gdsc-kmutt-basic-k8s/docs"
)

// @title Todo API
// @version 1.0
// @description This is a exmaple of Todo API for K8S traning
// @BasePath /
func main() {
	// Create Zap logger
	l, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalln("Failed to create Zap logger")
	}
	logger := l.Sugar()

	// Parse ENV
	dbConfig := config.DB{}
	var dbDialect gorm.Dialector
	if err := env.Parse(&dbConfig, env.Options{RequiredIfNoDef: true}); err != nil {
		logger.Info("Failed to parse Database ENV. Use sqlite as DB")
		dbDialect = sqlite.Open("todo.db")
	} else {
		logger.Info("Use MySQL as DB")
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
		dbDialect = mysql.Open(dsn)
	}

	// Open GORM DB
	db, err := gorm.Open(dbDialect)
	if err != nil {
		logger.Errorw("Unable to open GORM DB", "error", err.Error())
		os.Exit(1)
	}

	// Run auto migration in db
	if err := db.AutoMigrate(&todo.Todo{}); err != nil {
		logger.Errorw("Unable to auto migrate", "error", err.Error())
		os.Exit(1)
	}

	// Create Todo handler
	todoHandler := todo.NewHandler(db, logger)

	// Create fiber app
	app := fiber.New()

	app.Use(cors.New())
	app.Use(fiberLogger.New())

	// Test Route
	app.Get("/api/ping", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message":   "pong",
			"timestamp": time.Now(),
			"version":   2,
		})
	})

	app.Get("/api/todos", todoHandler.ListTodos)
	app.Post("/api/todo", todoHandler.CreateTodo)
	app.Patch("/api/todo/:id", todoHandler.FinishedTodo)
	app.Delete("/api/todo/:id", todoHandler.DeleteTodo)

	app.Get("/swagger/*", swagger.HandlerDefault)

	logger.Info("Version 2")
	// Start listening the request
	if err := app.Listen(":5050"); err != nil {
		logger.Errorw("Failed to start listening the request", "error", err.Error())
		os.Exit(1)
	}
}
