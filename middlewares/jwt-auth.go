package middlewares

import (
	"net/http"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(os.Getenv("SECRET_JWT")),
		SigningMethod: "HS256",
	})
}

func CreateTokenAdmin(adminID uuid.UUID, username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["adminID"] = adminID
	claims["username"] = username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))

}

func IsAdmin(c echo.Context) (uuid.UUID, error) {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return uuid.Nil, echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}
	claims := user.Claims.(jwt.MapClaims)
	if claims["admin"] != true {
		return uuid.Nil, echo.NewHTTPError(401, "Unauthorized")
	}

	adminID := claims["adminID"].(string)
	uid, err := uuid.Parse(adminID)
	if err != nil {
		return uuid.Nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	return uid, nil
}

func CreateTokenUser(userID uuid.UUID, nis string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = userID
	claims["nis"] = nis
	claims["admin"] = false
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}

func IsUser(c echo.Context) (uuid.UUID, error) {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return uuid.Nil, echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}
	claims := user.Claims.(jwt.MapClaims)
	if claims["admin"] != false {
		return uuid.Nil, echo.NewHTTPError(401, "Unauthorized")
	}

	userID := claims["userID"].(string)
	uid, err := uuid.Parse(userID)
	if err != nil {
		return uuid.Nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	return uid, nil
}
