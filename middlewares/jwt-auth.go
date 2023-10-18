package middlewares

import (
	"go-spp/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningMethod: "HS256",
	SigningKey:    []byte(os.Getenv("SECRET_JWT")),
})

func CreateAdminToken(adminID uuid.UUID, username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["adminID"] = adminID
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	byteSecret := []byte(os.Getenv("SECRET_JWT"))
	return token.SignedString(byteSecret)
}

func CreateStudentToken(studentID uuid.UUID, nis string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["studentID"] = studentID
	claims["nis"] = nis
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	byteSecret := []byte(os.Getenv("SECRET_JWT"))
	return token.SignedString(byteSecret)
}

func IsAdmin(c echo.Context) (int, error) {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return 0, echo.NewHTTPError(401, "Unauthorized")
	}

	admin := &models.Admin{}
	claims := user.Claims.(jwt.MapClaims)
	if claims["username"] != admin.Username {
		return 0, echo.NewHTTPError(401, "Unauthorized")
	}
	adminID := int(claims["adminID"].(float64))

	return adminID, nil
}
