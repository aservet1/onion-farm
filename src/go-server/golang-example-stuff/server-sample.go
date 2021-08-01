package main

import (
	"os"
	"fmt"
	"strconv"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func handleRequests(PORT int) {
	fmt.Printf("listening for requests on port %d ...\n", PORT);
	myRouter := mux.NewRouter().StrictSlash(true);
	myRouter.HandleFunc("/", homePage).Methods("GET");
	myRouter.HandleFunc("/api/{name}", helloPage).Methods("GET");
	log.Fatal( http.ListenAndServe( fmt.Sprintf(":%d", PORT), myRouter) );
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*");
	fmt.Fprintf(w,	"<body style='background-color: black'>" +
						"<h1 style='color: blue'> my server </h1>" +
					"</body>");
}

func helloPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r);
	name := vars["name"];
	msg := fmt.Sprintf("<body style='background-color: black'><h1 style='color: blue'>hello, %s!</h1></body>", name);
	fmt.Fprintf(w, msg);
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

	handleRequests(PORT);
}
