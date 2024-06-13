package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/tassanai001/go-fiber-crm-basic/database"
	"github.com/tassanai001/go-fiber-crm-basic/lead"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()
}
