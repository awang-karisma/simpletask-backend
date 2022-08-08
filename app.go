package main

import (
	"database/sql"
	"fmt"
	"os"
	"simpletask-backend/handlers"
	"strings"

	_ "github.com/glebarez/go-sqlite"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	server := echo.New()
	if os.Getenv("APP_CORS_URL") == "" {
		server.Logger.Fatal("APP_CORS_URL env not set")
	}
	corsHost := strings.Split(os.Getenv("APP_CORS_URL"), ",")
	fmt.Printf("CORS : %s", os.Getenv("APP_CORS_URL"))
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: corsHost,
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	db := initDB("storage.db")
	migrate(db)
	server.Static("/", "public")

	server.GET("/api/tasks", handlers.GetTasks(db))
	server.POST("/api/tasks", handlers.PutTask(db))
	server.PUT("/api/tasks/:id", handlers.EditTask(db))
	server.DELETE("/api/tasks/:id", handlers.DeleteTask(db))

	server.Logger.Fatal(server.Start(":8000"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite", filepath)
	if err != nil {
		panic(err)
	}

	if db == nil {
		panic(err)
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
CREATE TABLE IF NOT EXISTS tasks(
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	name VARCHAR NOT NULL,
	detail VARCHAR NOT NULL,
	assignee VARCHAR NOT NULL,
	due VARCHAR,
	status INTEGER 
);`
	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}
