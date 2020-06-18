package main

import (
	"fmt"
	"net/http"
)

var chats map[string](chan string)

func main() {
	chats = make(map[string](chan string))
	fmt.Println("Server is up")
	http.HandleFunc("/send", Send)
	http.HandleFunc("/receive", Receive)
	http.ListenAndServe(":8080", nil)
}

func showChannels() {
	for k, v := range chats {
		fmt.Printf("name: %v, channel: %v, size: %v\n", k, v, len(v))
	}
}

func Send(w http.ResponseWriter, r *http.Request) {
	chatName := r.URL.Query()["channel"][0]
	if chats[chatName] == nil {
		fmt.Printf("Creating channel: %v\n", chatName)
		chats[chatName] = make(chan string, 10)
	}
	text := r.URL.Query()["text"][0]
	chats[chatName] <- text
	fmt.Printf("Sent [%s:%s]\n", chatName, text)
	fmt.Fprintf(w, "Sent [%s:%s]\n", chatName, text)
	showChannels()
}

func Receive(w http.ResponseWriter, r *http.Request) {
	chatName := r.URL.Query()["channel"][0]
	if chats[chatName] == nil {
		fmt.Printf("Creating channel: %v\n", chatName)
		chats[chatName] = make(chan string, 10)
	}
	showChannels()
	fmt.Printf("Waiting for input in channel %s, %v\n", chatName, chats[chatName])
	select {
	case p := <-chats[chatName]:
		fmt.Fprintf(w, "Received: %s", p)
	default:
		fmt.Fprintf(w, "Nothin'")
	}
}
