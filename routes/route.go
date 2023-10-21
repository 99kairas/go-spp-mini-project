package routes

import (
	"go-spp/controllers"
	"go-spp/middlewares"
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func Routes(e *echo.Echo, db *gorm.DB) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/register/admin", controllers.RegisterAdminController)
	e.POST("/login/admin", controllers.LoginAdminController)
	e.POST("/login/student", controllers.LoginStudentController)

	// STUDENT

	authJWT := e.Group("/admin", middlewares.JWTMiddleware())
	authJWT.Use(echojwt.JWT([]byte(os.Getenv("SECRET_KEY"))))
	authJWT.POST("/register/student", controllers.RegisterStudentController)
	// e.POST("/admin/register/student", controllers.RegisterStudentController)
}
