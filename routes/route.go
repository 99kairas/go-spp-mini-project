package routes

import (
	"go-spp/controllers"
	"go-spp/middlewares"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

func Routes(e *echo.Echo, db *gorm.DB) {
	e.POST("/register/admin", controllers.RegisterAdminController)
	e.POST("/login/admin", controllers.LoginAdminController)
	e.POST("/register/student", controllers.RegisterStudentController)
	e.POST("/login/student", controllers.LoginStudentController)

	// STUDENT

	authJWT := e.Group("/admin", middlewares.IsLoggedIn)
	authJWT.Use(middleware.JWT([]byte(os.Getenv("SECRET_KEY"))))
	authJWT.POST("/register/student", controllers.RegisterStudentController)
}
