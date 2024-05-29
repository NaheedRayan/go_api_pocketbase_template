package main

import (
	"flag"
	"fmt"
	"go_api_pocketbase_template/routes"
	"net/http"

	"github.com/gorilla/mux"
)




func main() {

	// for argument passing
	// go run main.go --listenAddr :2000
	listenAddr := flag.String("listenAddr" , ":3000" , "the server address")
	flag.Parse()// Parse the command-line flags

	

	router := mux.NewRouter()
	router.HandleFunc("/", routes.Hello).Methods("GET")
	router.HandleFunc("/posts", routes.GetPosts).Methods("GET")
	router.HandleFunc("/posts", routes.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{id}", routes.GetPost).Methods("GET")
	router.HandleFunc("/posts/{id}", routes.UpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", routes.DeletePost).Methods("DELETE")



	fmt.Println("Server running at port", *listenAddr)
	// http.ListenAndServe(":8000", router)
	http.ListenAndServe(*listenAddr, router)



}