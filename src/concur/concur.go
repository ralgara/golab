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
	http.HandleFunc("/", Home)
	http.ListenAndServe(":8080", nil)
}

func showChannels() {
	for k, v := range chats {
		fmt.Printf("name: %v, channel: %v, size: %v\n", k, v, len(v))
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<html>
		<body>
			<h2>Go Chat</h4>
			<h4>Send</h4>
			<form action="/send">
				Channel: <input name="channel" type="text"/>
				Text: <input name="text" type="text"/>
				<input type="submit"/>
			</form>
			<br/>
			<h4>Receive</h4>
			<form action="/receive">
				Channel: <input name="channel" type="text"/>
				<input type="submit"/>
			</form>
		</body>
	</body>
	`)
}

func GetChannel(chatName string) chan string {
	if chats[chatName] == nil {
		fmt.Printf("Creating channel: %v\n", chatName)
		chats[chatName] = make(chan string, 10)
	}
	return chats[chatName]
}

func Send(w http.ResponseWriter, r *http.Request) {
	chatName := r.URL.Query()["channel"][0]
	channel := GetChannel(chatName)
	text := r.URL.Query()["text"][0]
	channel <- text
	caption := fmt.Sprintf("Sent [%s:%s]\n", chatName, text)
	fmt.Printf(caption)
	fmt.Fprintf(w, caption)
	showChannels()
}

func Receive(w http.ResponseWriter, r *http.Request) {
	chatName := r.URL.Query()["channel"][0]
	channel := GetChannel(chatName)
	showChannels()
	fmt.Printf("Waiting for input in channel %s, %v\n", chatName, channel)
	select {
	case p := <-channel:
		fmt.Fprintf(w, "Received: %s", p)
	default:
		fmt.Fprintf(w, "Nothin'")
	}
}
