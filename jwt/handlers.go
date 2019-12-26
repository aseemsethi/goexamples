/*
 * curl --header "Content-Type: application/json"   --request POST   --data '{"username":"user1","password":"password1"}' localhost:8000/signin

 *  curl -v --cookie "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIxIiwiZXhwIjoxNTc3MzI3MjEzfQ.aiK0U4igyUCoqblvkXRoUILr7zJ6kSCGV-rjQQCKOE4" localhost:8000/welcome
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
	// A two-value assignment tests for the existence of a key in users db
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

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome called")
	c, err := r.Cookie("token")
	if (err != nil) {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			fmt.Println("Cookie not found")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := c.Value  // Get the JWT token
	claims := &Claims{} // Init claims structure
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println("JWT Signature invalid")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		fmt.Println("JWT invalid")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Finally, return the welcome message to the user, along with their
	// username given in the token
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))

}

