package main

import (
	"os"
	"fmt"
	"log"
	"bytes"
	"net/url"
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/gorilla/mux"
);

var NODE_ID int; // find out how to make this const but dynamically defined (ie defined by command line args at runtime, but unchangeable after that)

type RelayMessage struct {
	Msg string `json:"msg"`
	RelayIndex int `json:relayIndex`
	NodeURLs []string `json:nodeURLs`
}

func _nodePrint(msg string) {
	fmt.Printf("Node %d: %s\n", NODE_ID, msg)
}

func handleRequests(PORT int) {
	fmt.Printf("listening for requests on port %d ...\n", PORT);

	myRouter := mux.NewRouter().StrictSlash(true);
	myRouter.HandleFunc("/api/relay", Relay).Methods("POST");

	log.Fatal ( http.ListenAndServe( fmt.Sprintf(":%d", PORT), myRouter) );
}

func ParseRelayMessage( writer http.ResponseWriter, request *http.Request ) RelayMessage {
	var relaymsg RelayMessage

	b, err:= ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return relaymsg
	}
	err = json.Unmarshal(b, &relaymsg)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return relaymsg
	}

	return relaymsg
}

func Relay (writer http.ResponseWriter, request *http.Request) {
	relayMsg := ParseRelayMessage(writer, request)
	if relayMsg.RelayIndex >= len(relayMsg.NodeURLs) {
		_nodePrint("Received message: " + relayMsg.Msg)
		return
	}

	_nodePrint("relaying message to " + relayMsg.NodeURLs[relayMsg.RelayIndex] + "...")
	NextURL, _ := url.Parse(relayMsg.NodeURLs[relayMsg.RelayIndex])
	NextURL.Path += "/api/relay"

	relayMsg.RelayIndex = relayMsg.RelayIndex + 1
	json_data, err := json.Marshal(relayMsg)
	if err != nil {
		log.Fatal(err)
	}

	//todo: actually read and work with the response to this
	response, err2 := http.Post(NextURL.String(), "application/json", bytes.NewBuffer(json_data))
	if err2 != nil {
		log.Fatal(err)
	}

	fmt.Fprint(writer, response)
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
