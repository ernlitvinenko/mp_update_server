package updates

import "github.com/gofiber/fiber/v2"

func InitRoutes(router fiber.Router) {
	group := router.Group("/updates")
	// TODO Write your endpoints here

	group.Post("/create-app", CreateApplication)
	group.Post("/upload", UploadApp)
	group.Get("/list-app", ListApplications)
	group.Post("/create", AddVersion)
}
