/*
 * curl --header "Content-Type: application/json"   --request POST   --data '{"username":"user1","password":"password1"}' localhost:8000/signin
 */
package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"encoding/json"
	"net/http"
	"time"
)

var jwtKey = []byte("my_secret_key")

var users = map[string]string{
	"user1":"password1",
	"user2":"password2",
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Signin(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	fmt.Println("Signin Called...")
	err := json.NewDecoder(r.Body).Decode(&creds)
	if (err != nil) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Signin - decode data in request failed")
		return
	}
	expectedPassword, ok := users[creds.Username]
	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("Signin - Unauthoried")
		return
	}
	fmt.Println("Signin - pwd ok")
	expirationTime := time.Now().Add(5*time.Minute)
	claims := &Claims {
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if (err != nil) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Signin - could not sign JWT")
		return
	}
	fmt.Println("JWT: ", tokenString)
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: tokenString,
		Expires: expirationTime,
	})
}

