package middlewares

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return mid.JWTWithConfig(mid.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(os.Getenv("SECRET_JWT")),
	})
}

func CreateTokenAdmin(adminID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["adminID"] = adminID
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))

}

func ExtractAdminToken(e echo.Context) (int, error) {
	user := e.Get("user").(*jwt.Token)
	if !user.Valid {
		return 0, echo.NewHTTPError(401, "Unauthorized")
	}
	claims := user.Claims.(jwt.MapClaims)
	if claims["admin"] != true {
		return 0, echo.NewHTTPError(401, "Unauthorized")
	}
	adminID := int(claims["adminID"].(float64))

	return adminID, nil
}

func CreateTokenUser(userId uuid.UUID) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["admin"] = false
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))

}

func ExtractUserToken(e echo.Context) (int, error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return int(userId), nil
	}
	return 0, fmt.Errorf("token invalid")
}
