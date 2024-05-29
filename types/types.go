package types

// for individual post
type Post struct {
	ID            string    `json:"id"`
	CollectionID  string    `json:"collectionId"`
	CollectionName string   `json:"collectionName"`
	Created       string 	`json:"created"`
	Updated       string 	`json:"updated"`
	Title         string    `json:"title"`
	Body          string    `json:"body"`
}

// for list of posts
// ListUsers represents the root structure of your JSON.
type ListUsers struct {
	Page       int     `json:"page"`
	PerPage    int     `json:"perPage"`
	TotalItems int     `json:"totalItems"`
	TotalPages int     `json:"totalPages"`
	Items      []Item  `json:"items"`
}

// Item represents the individual items within the JSON.
type Item struct {
	Body            string    `json:"body"`
	CollectionID   	string    `json:"collectionId"`
	CollectionName 	string    `json:"collectionName"`
	Created       	string    `json:"created"`
	ID       		string    `json:"id"`
	Title 			string    `json:"title"`
	Updated         string    `json:"updated"`
	
}
