package routes

import (
	"../controllers"
	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App){
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

}