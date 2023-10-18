package middlewares

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

func CreateToken(adminID uuid.UUID, username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["adminID"] = adminID
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	byteSecret := []byte(os.Getenv("SECRET_JWT"))
	return token.SignedString(byteSecret)
}
