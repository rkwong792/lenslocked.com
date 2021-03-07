package main

import (
	"fmt"
	"net/http"
)

//Takes in a ResponseWriter - this is where we write our response to the user

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func main() {
	//built in router where we want to map the request. No matter what path they give us, send to the handlerFunc
	http.HandleFunc("/", handlerFunc)

	//localhost:3000, port 3000
	//starts up the server
	http.ListenAndServe(":3000", nil) //when the second parameter, the handler, is nil, use the DefaultServeMux
}
