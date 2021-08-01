package main

import (
	"os"
	"fmt"
	"strconv"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

const SERVER_ID int;

func handleRequests(PORT int) {
	fmt.Printf("listening for requests on port %d ...\n", PORT);

	myRouter := mux.NewRouter().StrictSlash(true);
	myRouter.HandleFunc("/", homePage).Methods("GET");
	myRouter.HandleFunc("/api/send/{msg}", sendMessage).Methods("POST");
	myRouter.HandleFunc("/api/receive", receiveMessage).Methods("POST");

	log.Fatal( http.ListenAndServe( fmt.Sprintf(":%d", PORT), myRouter) );
}

func handleSend(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r);
	msg := vars["msg"];
	// fmt.Fprintf(w, msg);
}

func handleReceive(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r);
	msg := vars["msg"];
	fmt.Printf("Server %d received message: %s\n", SERVER_ID, msg);
}

func main() {
	var PORT int;
	if len(os.Args) > 2 {
		p, err := strconv.Atoi(os.Args[2]);
		if err != nil {
			log.New(os.Stderr, "", 0).Printf("invalid port number: %s\n", os.Args[1]);
			os.Exit(1);
		}
		PORT = p;
	} else {
		PORT = 8080;
		fmt.Printf("no port number provided by command line, will use default %d\n", PORT);
	}
	SERVER_ID = PORT; // maybe come up with a better server id convention, or not. this will do fine for now

	handleRequests(PORT);
}
