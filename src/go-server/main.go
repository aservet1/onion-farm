package main

import (
	"os"
	"fmt"
	"strconv"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux" // locally installed library, i think was pulled with wget at some point (important dependency to note for deploying purposes)
)

type ErrorStruct struct {
	ErrorStr string `json:"error"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*");
	vars := mux.Vars(r);
	name := vars["name"];
	msg := fmt.Sprintf("hello, %s!", name);
	json.NewEncoder(w).Encode(msg);
}

func handleRequests(PORT int) {
	myRouter := mux.NewRouter().StrictSlash(true);
	myRouter.HandleFunc("/api/{name}", homePage).Methods("GET");
	log.Fatal( http.ListenAndServe( fmt.Sprintf(":%d", PORT), myRouter) );
}

func main() {
	PORT := 8080;
	if len(os.Args) == 1 {
		fmt.Printf("no port number provided by command line, will use default %d\n", PORT);
	} else {
		PORT, _ = strconv.Atoi(os.Args[1]);
		// if err != nil {
		// 	log.New(os.Stderr, "", 0).Printf("invalid port number: %s\n", os.Args[1]);
		// }
	}

	fmt.Printf("listening for requests on port %d ...\n", PORT);
	handleRequests(PORT);
}
