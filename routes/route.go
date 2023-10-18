package routes

import (
	"go-spp/controllers"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func Routes(e *echo.Echo, db *gorm.DB) {
	e.POST("/register/admin", controllers.RegisterAdminController)
	e.POST("/login/admin", controllers.LoginAdminController)
}
