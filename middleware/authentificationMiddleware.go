package middleware

import (
	"github.com/dgrijalva/jwt-go"
	customHTTP "github.com/hellojebus/go-envoz-api/http"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func VerifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Authentication failure")
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := VerifyToken(tokenString)
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error verifying JWT token: " + err.Error())
			return
		}

		//pass userId claim to req
		//todo: find a better way to convert the claim to string
		userId := strconv.FormatFloat(claims.(jwt.MapClaims)["user_id"].(float64), 'g', 1, 64)
		r.Header.Set("userId", userId)
		next.ServeHTTP(w, r)
	})
}

func extractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecret := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})
	if err != nil {
		return nil, false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} //else
	log.Printf("Invalid JWT Token")
	return nil, false

}
