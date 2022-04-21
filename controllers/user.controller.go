package controllers

import (
	"encoding/json"
	"fmt"
	"golang-crud/database"
	"golang-crud/models"
	"log"
	"net/http"
)

// @desc -> if user endpoint
func UserHome(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "user" {
		w.Write([]byte("Not Authorized"))
		return
	}
	w.Write([]byte("Welcome ,user"))
}

// @desc -> User Register End Point
func Register(w http.ResponseWriter, r *http.Request) {
	connection := database.Database
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		var err error
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	var dbUser models.User
	connection.Where("email = ?", user.Email).First(&dbUser)

	// check email is already registered or not
	if dbUser.Email != "" {
		var err error
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	user.Password, err = GenerareHashPassword(user.Password)
	if err != nil {
		log.Fatalln("Error in Paddword Hashing")
	}
	connection.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}
func Login(w http.ResponseWriter, r *http.Request) {
	connection := database.Database
	var authDetails models.Authentication
	err := json.NewDecoder(r.Body).Decode(&authDetails)
	if err != nil {
		var err error
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	var authUser models.User

	connection.Where("email= ?", authDetails.Email).First(&authUser)
	if authUser.Email == "" {
		var err error
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	// check := CheckPasswordHash(authDetails.Password, authUser.Password)
	check := CompareHashAndPlainPassword(authDetails.Password, authUser.Password)
	if !check {
		var err error
		fmt.Println("Username or Password is incorrect")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	validToken, err := GenerateJWT(authUser.Email, authUser.Role)
	if err != nil {
		var err error
		fmt.Println("Failed to generate token")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	var token models.Token
	token.Email = authUser.Email
	token.Role = authUser.Role
	token.TokenString = validToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}
