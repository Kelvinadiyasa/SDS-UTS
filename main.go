package main

import (
	"uts/database"
	"uts/models"
	"uts/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Menghubungkan ke database
	database.Connect()

	app := fiber.New()

	// Migrasi model User (jika belum dilakukan)
	database.DB.AutoMigrate(&models.User{})

	// Routes
	app.Post("/insert", route.InsertData)
	app.Get("/getData", route.GetAllData)
	app.Get("/getDataUser/:id_user", route.GetUserByid)
	app.Get("/delete/:id_user", route.Delete)
	app.Put("/update/:id_user", route.Update)

	app.Listen(":3000")
}
