	package main

import (
	"log"
	"net/http"
)

type (
	MessagePoster struct {
		
	}
)

func main () {
	http.HandleFunc ("/post",PostMessage)
	http.HandleFunc ("/login",Login)

	log.Print("\u001B[32mstarting server\u001B[0m")
	if err := http.ListenAndServe(":6969",nil); err != nil {
		log.Fatal(err)
	}
	
}

func PostMessage (writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		if err := request.ParseForm(); err != nil {
			log.Fatal(err)
		}

	log.Printf( "Post from website! r.PostFrom = %v\n", request.PostForm)
		
	} else {
		panic("only post method is allowed")
	}
}

func Login (writer http.ResponseWriter, request *http.Request) {
	log.Print("OK")
	if request.Method == "POST" {
		if err := request.ParseForm(); err != nil {
			log.Fatal(err)
		}
	
	log.Printf( "Post from website! r.PostFrom = %v\n", request.PostForm)
			
	} else {
		panic("only post method is allowed")
	}
}