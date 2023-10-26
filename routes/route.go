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

	// FOR ADMIN
	authJWT := e.Group("/admin", middlewares.JWTMiddleware())
	authJWT.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))
	authJWT.POST("/register/student", controllers.RegisterStudentController)
	authJWT.GET("/student/:id", controllers.GetStudentIDController)
	// e.POST("/admin/register/student", controllers.RegisterStudentController)

	// SPP
	e.POST("/admin/spp", controllers.CreateSPPController)
}
