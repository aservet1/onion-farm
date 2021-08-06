package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"bytes"
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/gorilla/mux"
);

var NODE_ID int; // find out how to make this const but dynamically defined (ie defined by command line args at runtime, but unchangeable after that)

func __nodePrint(msg string) {
	fmt.Printf("Node %d: %s\n", NODE_ID, msg)
}

func handleRequests(PORT int) {
	fmt.Printf("listening for requests on port %d ...\n", PORT);

	myRouter := mux.NewRouter().StrictSlash(true);
	// myRouter.HandleFunc("/", homePage).Methods("GET");
	myRouter.HandleFunc("/api/sendMessage", sendMessage).Methods("POST");
	myRouter.HandleFunc("/api/receiveMessage", receiveMessage).Methods("POST");

	log.Fatal ( http.ListenAndServe( fmt.Sprintf(":%d", PORT), myRouter) );
}

func sendMessage (
  writer http.ResponseWriter,
  request *http.Request
) {

	body := map[string]string{"filename":"hello.txt","textBody":".hellow0rld."}
	json_data, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err);
	}

	Url, + := url.Parse("http://localhost:8081") // make this URL dynamically chosen, based on the whole dimensions and message information of the onion message
	Url.Path += "/api/receiveMessage"
	request, _ := http.Post(urL, "application/json", bytes.NewBuffer(json_data))

	//- get a body of json data in this post request

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return "error happened"
	}
}

func receiveMessage (
  writer http.ResponseWriter,
  request *http.Request
) {
	__nodePrint("Received message at " + string(time.Now()))
	fmt.Println("Here's the body: ")
	body := request.Body
	fmt.Println ( body )
}

func main() {
	var PORT int;
	if len(os.Args) > 1 {
		p, err := strconv.Atoi(os.Args[1]);
		if err != nil {
			log.New(os.Stderr, "", 0).Printf("invalid port number: %s\n", os.Args[1]);
			os.Exit(1);
		}
		PORT = p;
	} else {
		PORT = 8080;
		fmt.Printf("no port number provided by command line, will use default %d\n", PORT);
	}
	NODE_ID = PORT; // maybe come up with a better node id convention, or not. this will do fine for now
	handleRequests(PORT);
}
