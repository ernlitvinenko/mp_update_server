package auth

import "github.com/gofiber/fiber/v2"

func InitRoutes(router fiber.Router) {
	group := router.Group("/auth")

	group.Get("/me", UserInfo)
	group.Post("/login", Login)
}
