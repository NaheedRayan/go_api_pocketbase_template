## `GO` +`MUX` + `POCKETBASE` Template
This is a `GO` API template with MUX where the backend database is `POCKETBASE`. This is just for testing purpose.

## Basic Structure of the app

```python
go_api_pocketbase_template
|
|--pocketbase_0.22.12_linux_amd64  # The Pocketbase executable
|--routes
|   |--routes.go                   # Go file for defining routes
|
|--types
|   |--types.go                    # Go file for defining types
|
|--go.mod                          # Go module file
|--go.sum                          # Go dependencies checksum file
|--main_test.go                    # Testing file
|--main.go                         # Main entry point of the Go application
|--README.md                       # Documentation file

```




## Usage

### 1. Running the pocketbase server.
```
cd pocketbase_0.22.12_linux_amd64 && ./pocketbase serve
```

Click on the Admin UI: http://127.0.0.1:8090/_/

`Username`:`joker@gmail.com`

`Password`:`0123456789`


``Note:`` You can also download `pocketbase_0.22.12_linux_amd64` from https://pocketbase.io/docs/ and extract it on the file structure provided above.

### 2. Running the GOLANG REST api server.
```
go run main.go
```

You can also specify port number
```
go run main.go --listenAddr :2000
```







## important commands
```go
go mod init go_api_pocketbase_template
go mod tidy
```


## Endpoints
Use `postman` or `thunder client` for testing the endpoints
Sure, here are all the endpoints summarized in six lines:

- **GET** `/` - Check server status (`routes.Hello`) - `http://localhost:3000/`
- **GET** `/posts` - Retrieve all posts (`routes.GetPosts`) - `http://localhost:3000/posts`
- **POST** `/posts` - Create a new post (`routes.CreatePost`) - `http://localhost:3000/posts`
- **GET** `/posts/{id}` - Retrieve a post by ID (`routes.GetPost`) - `http://localhost:3000/posts/{id}`
- **PUT** `/posts/{id}` - Update a post by ID (`routes.UpdatePost`) - `http://localhost:3000/posts/{id}`
- **DELETE** `/posts/{id}` - Delete a post by ID (`routes.DeletePost`) - `http://localhost:3000/posts/{id}`


## Config

You can config the `pocketbase` url and endpoint on `routes/routes.go` on line 15,16.






## Running the Application and Tests
To run your application with a custom listening address:

```bash
go run main.go --listenAddr :2000
```

To run the tests locally:
```bash
go test -v ./...
```


Testing the main app
```bash
go test -v main_test.go
```

Running the pocketbase server.
```
cd pocketbase_0.22.12_linux_amd64 && ./pocketbase serve
```

Testing the pocketbase server
```bash
go test -v server_test.go
```

