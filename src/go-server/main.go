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
	/// myRouter.HandleFunc("/api/sendMessage", sendMessage).Methods("POST");
	/// myRouter.HandleFunc("/api/receiveMessage", receiveMessage).Methods("POST");
	myRouter.HandleFunc("/api/ping", ping).Methods("POST");
	myRouter.HandleFunc("/api/pong", pong).Methods("POST");

	log.Fatal ( http.ListenAndServe( fmt.Sprintf(":%d", PORT), myRouter) );
}

type PiPoMessage struct {
	Msg string `json:"msg"`
	Counter int `json:"counter"`
}

func send(msgMap map[string]string, url string) {
	json_data, err := json.Marshal(msgMap)
	if err  != nil { log.Fatal(err ) }
	_, err2 := http.Post(url, "application/json", bytes.NewBuffer(json_data))
	if err2 != nil { log.Fatal(err2) }
}

func PipoParseBody( writer http.ResponseWriter, request *http.Request ) PiPoMessage {
	var ppm PiPoMessage

	b, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		http.Error(writer, err.Error(), 500)
	}
	err = json.Unmarshal(b, &ppm)
	if err != nil {
		http.Error(writer, err.Error(), 500)
	}

	return ppm
}

func ping ( writer http.ResponseWriter, request *http.Request ) {
	ppm := PipoParseBody(writer, request)
	if ppm.Counter < 0 {
		__nodePrint("done!")
		return
	}
	__nodePrint(ppm.Msg)
	time.Sleep(500 * time.Millisecond)

	var pingMsg PiPoMessage;
	pingMsg.Msg = "ping"
	pingMsg.Counter = ppm.Counter-1;

	json_data, err := json.Marshal(pingMsg)
	if err != nil { log.Fatal(err) }
	_, err2 := http.Post("http://localhost:8081/api/pong", "application/json", bytes.NewBuffer(json_data))
	if err2 != nil { log.Fatal(err2) }
}
func pong ( writer http.ResponseWriter, request *http.Request ) {
	ppm := PipoParseBody(writer, request)
	if ppm.Counter < 0 {
		__nodePrint("done!")
		return
	}
	__nodePrint(ppm.Msg)
	time.Sleep(500 * time.Millisecond)

	var pongMsg PiPoMessage;
	pongMsg.Msg = "pong"
	pongMsg.Counter = ppm.Counter-1;

	json_data, err := json.Marshal(pongMsg)
	if err != nil { log.Fatal(err) }
	_, err2 := http.Post("http://localhost:8080/api/ping", "application/json", bytes.NewBuffer(json_data))
	if err2 != nil { log.Fatal(err2) }
}

// func sendMessage (
//   writer http.ResponseWriter,
//   request *http.Request
// ) {
// 
// 	body := map[string]string{"filename":"hello.txt","textBody":".hellow0rld."} // this will be replaced by the more complex onion message, not just two filename parameters
// 	json_data, err := json.Marshal(body)
// 	if err != nil {
// 		log.Fatal(err);
// 	}
// 
// 	// msg := json.Marshal({"filename": "hello.txt", "content": "hello a b c x y z lmno q e r ee"})
// 	// fmt.Fprintf(writer, msg);
// 	Url, + := url.Parse("http://localhost:8081") // make this URL dynamically chosen, based on the whole dimensions and message information of the onion message
// 	Url.Path += "/api/receiveMessage"
// 	request, _ := http.Post(urL, "application/json", bytes.NewBuffer(json_data))
// 
// 	//- get a body of json data in this post request
// 
// 	res, err := http.DefaultClient.Do(request)
// 	if err != nil {
// 		fmt.Println(err)
// 		return "error happened"
// 	}
// }
// 
// func receiveMessage (
//   writer http.ResponseWriter,
//   request *http.Request
// ) {
// 	__nodePrint("Received message at " + string(time.Now()))
// 	fmt.Println("Here's the body: ")
// 	body := request.Body
// 	fmt.Println ( body )
// }

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
