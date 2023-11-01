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
	authJWT.GET("/student", controllers.GetAllStudentController)
	authJWT.POST("/spp", controllers.CreateSPPController)
	authJWT.POST("/payment/student", controllers.AdminCreatePaymentByStudentIDController)
	authJWT.POST("/payment", controllers.AdminCreatePaymentAllStudentController)
	authJWT.GET("/payment", controllers.GetAllPaymentsController)
	authJWT.GET("/payment/photo", controllers.GetPaymentsWithPhotoController)
	authJWT.GET("/payment/details/:id", controllers.GetPaymentByIDController)
	authJWT.PUT("/payment/approve", controllers.ApprovePaymentController)
	authJWT.PUT("/payment/reject", controllers.RejectPaymentController)

	// FOR STUDENT
	authStudent := e.Group("/student", middlewares.JWTMiddleware())
	authStudent.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))
	authStudent.PUT("/password/change", controllers.UpdatePasswordController)
	authStudent.PUT("/profile", controllers.UpdateProfileController)
	authStudent.POST("/upload", controllers.UploadPaymentPhotoController)
}
