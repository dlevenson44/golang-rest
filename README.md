## Golang REST API
###### Goals
- Create a server that will start a server that will listen to requests
- Will hook server up to a DB, add ability to CRUD within the DB

###### Notes
- `http.ListenAndServe` will be used to start server, tell it to listen for new HTTP requests, then serve them using handler functions-- ie `http.ListenAndServe(":3333", nil)`
  - first argument specifies port to listen on-- without specifying this, it will listen on all available ports on your machine
- `http.HandleFunc` tells the server which function to call to handle a request to the server
- `func handlerName(w http.ResponseWriter, r *http.Request)` is the func signature defined as `http.HandlerFunc`
  - When request is made to server, it sets these two values with info about the request made and then calls the handler func with those values
  - `http.ResponseWriter` value is used to control the response info being written back to the client that made the request
    - such as body of response, status code, error messages, etc.
  - `http.Request` value is used to get info about the request that came into the server
    - such as request method, body, query params, headers, etc.
  - 

###### Resources
- [Creating Go Server/Endpoints](https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go)
