package main

import (
	"encoding/json"
	"log"
	"net/http"

	"context"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/get", getHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	queryValues := r.URL.Query()
	// query string
	q := queryValues.Get("q")

	// Create a context
	// Assign to empty context
	ctx := context.WithValue(context.Background(), "q", q)

	// Create a context base on original request
	c := context.WithValue(r.Context(), "q", q)
	// Assign to original request
	r = r.WithContext(c)

	// Print q
	log.Print("===Message===")
	log.Printf("Incoming q = %s", q)
	log.Printf("ctx q = %s", ctx.Value("q"))
	log.Printf("ctx inside original request q = %s", r.Context().Value("q"))
	log.Println("===End===")
	msg := map[string]interface{}{
		"message": "Example working",
		"q":       ctx.Value("q"),
	}
	json.NewEncoder(w).Encode(&msg)
}
