package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"golang.org/x/crypto/bcrypt"
)

// @desc -> Genreate Hash password from plain password
func GenerareHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// @desc -> Compare Hash(stored) and newly entered  password
func CompareHashAndPlainPassword(password, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}

// @desc -> Generate Jwt Token
func GenerateJWT(email, role string) (string, error) {
	var secretKey = []byte("1234")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Printf("Failed TO Generate Jwt Key %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

// @desc -> check wheather user is authenticated or not
func IsAuthenticated(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] == nil {
			var err error
			json.NewEncoder(w).Encode(err)
			return
		}
		var secretKey = []byte("1234")
		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing token.")
			}
			return secretKey, nil
		})
		if err != nil {
			var err error
			json.NewEncoder(w).Encode(err)
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "admin" {
				r.Header.Set("Role", "admin")
				return
			} else if claims["role"] == "user" {
				r.Header.Set("Role", "user")
				handler.ServeHTTP(w, r)
				return
			}
		}

		json.NewEncoder(w).Encode(err)

	} // end handler sub func

} // end main IsAUthenticated Function
