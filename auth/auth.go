package auth

import (
	"log"
	"net/http"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/joho/godotenv"
)

var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("godotenv:", err)
	}

	// set token header
	token := jwt.New(jwt.SigningMethodHS256)

	// set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = "54546557354"
	claims["name"] = "taro"
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, _ := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))

	w.Write([]byte(tokenString))
})

var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(t *jwt.Token) (interface{}, error) {
		err := godotenv.Load(".env")
		if err != nil {
			log.Println("godotenv:", err)
		}
		return []byte(os.Getenv("SIGNINGKEY")), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
