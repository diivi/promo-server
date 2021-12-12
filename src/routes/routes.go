package routes

import (
	"promo/src/controllers"
	"promo/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	admin := api.Group("admin")
	admin.Post("/register", controllers.Register)
	admin.Post("/login", controllers.Login)

	adminAuthenticated := admin.Use(middleware.IsAuthenticated)
	adminAuthenticated.Post("/user",controllers.GetAuthUser)
	adminAuthenticated.Post("/logout", controllers.Logout)
}
