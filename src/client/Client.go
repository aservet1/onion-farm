package main

import (
	"os"
	"fmt"
	"log"
	"bufio"
	//"net/url"
	"bytes"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
	"math/rand"
	"regexp"
);

type RelayMessage struct {
	Msg string `json:"msg"`
	RelayIndex int `json:"relayIndex"`
	NodeURLs []string `json:"nodeURLs"`
}

func RandomSubset(list []string, n int) []string {
	rand.Seed(time.Now().UnixNano())
	var newlist []string
	for i := 0; i < n; i++ {
		newlist = append(newlist, list[rand.Intn(n-0)+0])
	}
	return newlist
}

func CreateRelayMessage(text string, networkurls []string) RelayMessage {
	var message RelayMessage
	message.RelayIndex = 0
	message.NodeURLs = RandomSubset(networkurls, 3)
	message.Msg = text
	return message
}

func ReadURLFile(filename string) []string {
	bytesRead, _ := ioutil.ReadFile(os.Args[1])
	fileContent := string(bytesRead)
	re := regexp.MustCompile(`/#.*$/`)
	fileContent = re.ReplaceAllString(fileContent, "")
	URLs := strings.Split(fileContent, "\n")
	return URLs
}

func ReadMessageFromInput() string {
	fmt.Print("enter the message you want to send: ")
	keyboard := bufio.NewReader(os.Stdin)
	text, _ := keyboard.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s network-url-file\n", os.Args[0])
		os.Exit(1)
	}
	message := CreateRelayMessage(ReadMessageFromInput(), ReadURLFile(os.Args[1]))
	firstURL := message.NodeURLs[0] + "/api/relay" //todo: do actual url builder
	json_data, _ := json.Marshal(message)
	fmt.Println("sending across relay:",message.NodeURLs)
	response, err := http.Post(firstURL, "application/json", bytes.NewBuffer(json_data))
	if err != nil { log.Fatal(err) }
	fmt.Println("response:",response)
}

