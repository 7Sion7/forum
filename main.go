package main

import (
	"fmt"
	"forum/models"
	r "forum/routes"
	"log"
	"net/http"

	//"os/exec"

	_ "github.com/mattn/go-sqlite3"
)

// func init() {
// 	c.Tpl = template.Must(template.ParseGlob("templates/*.html"))
// }

func main() {
	models.InitDB()

	mux := http.NewServeMux()
	r.SetUpRoutes(mux)
	
	fmt.Println("Serving on Port -> http://localhost:8888")
	if err := http.ListenAndServe(":8888", mux); err != nil {
		log.Fatalf("Failure on Listening and Serving: %v", err)
	}

}
