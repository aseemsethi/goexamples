/*
 * Learnings from https://github.com/sohamkamani/jwt-go-example/blob/master/handlers.go
 */
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/signin", Signin)
	//http.HandleFunc("/welcome", Welcome)
	//http.HandleFunc("/signin", Signin)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
