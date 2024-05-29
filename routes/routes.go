package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go_api_pocketbase_template/types"

	"github.com/gorilla/mux"
)

// pocketbase url
var pb_DB = "http://127.0.0.1:8090"
var pb_endpoint = "/api/collections/posts/records/"

// for testing the golang and mux server
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w , "Hello from the GOLANG and MUX")
}

// for getting the List of posts
func GetPosts(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get(pb_DB+pb_endpoint)
	if err != nil {log.Fatal(err)}

	defer resp.Body.Close()

	// decoding JSON data into struct
	var listUsers types.ListUsers
	_ = json.NewDecoder(resp.Body).Decode(&listUsers)
	fmt.Println("GET request for list of Posts")

	// Encoding struct data to JSON and passing it to 'w'
  	w.Header().Set("Content-Type", "application/json")
  	json.NewEncoder(w).Encode(&listUsers)	
}


// For creating a post
func CreatePost(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Post(pb_DB+pb_endpoint,"application/json",r.Body)
	if err != nil {log.Fatal(err)}
	defer resp.Body.Close()

	// decoding JSON data into struct
	var post types.Post
	_ = json.NewDecoder(resp.Body).Decode(&post)
	fmt.Println("POST request for new Post")

	// Encoding struct data to JSON and passing it to 'w'
	w.Header().Set("Content-Type", "application/json")
  	json.NewEncoder(w).Encode(&post)	

}

// For getting a single Post
func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	resp, err := http.Get(pb_DB+pb_endpoint+params["id"])
	if err != nil {log.Fatal(err)}
	defer resp.Body.Close()

	// decoding JSON data into struct
	var post types.Post
	_ = json.NewDecoder(resp.Body).Decode(&post)
	fmt.Println("GET request for a single Post")

	// Encoding struct data to JSON and passing it to 'w'
  	w.Header().Set("Content-Type", "application/json")
  	json.NewEncoder(w).Encode(&post)
}

// For updating a single Post
func UpdatePost(w http.ResponseWriter, r *http.Request) {


	params := mux.Vars(r)

	// create new HTTP PATCH request with JSON payload
	req, err := http.NewRequest("PATCH", pb_DB+pb_endpoint+params["id"], r.Body)
	if err != nil {panic(err)}

	// set content-type header to JSON
	req.Header.Set("Content-Type", "application/json")

	// create HTTP client and execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {panic(err)}
	defer resp.Body.Close()

	fmt.Println("PATCH request for updating")

	// decoding JSON data into struct
	var post types.Post
	_ = json.NewDecoder(resp.Body).Decode(&post)

	// Encoding struct data to JSON and passing it to 'w'
	w.Header().Set("Content-Type", "application/json")
  	json.NewEncoder(w).Encode(&post)	
	
}

// for deleting post
func DeletePost(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	
    // create a new DELETE request
    req, err := http.NewRequest("DELETE",pb_DB+pb_endpoint+params["id"] , nil)
    if err != nil {panic(err)}

    // send the request
	// create a new HTTP client
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {panic(err)}
    defer resp.Body.Close()


	fmt.Println("DELETE request has been made")

	// Passing the JSON to 'w'
  	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w,resp.Body)
}