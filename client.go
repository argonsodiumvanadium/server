package main

import (
	"log"
	"net/http"
	"fmt"
	"io/ioutil"
	"bufio"
	"strings"
	"os"
)

const HOST_NAME = "http://localhost:6969/"

func main () {	
	var ip string
	
	fmt.Print("enter IP : ")
	fmt.Scanln(&ip)


	var name string
	
	fmt.Print("enter name : ")
	fmt.Scanln(&name)

	data := map[string][]string{"name": {name}}

	response , err := http.PostForm("http://"+ip+":6969/login",data)
	
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		print(err)
	}

	fmt.Println(string(body))

	go Listen()

	for {
		fmt.Print(" : ")

		reader := bufio.NewReader(os.Stdin)
		// ReadString will block until the delimiter is entered
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			return
		}

		// remove the delimeter from the string
		msg = strings.TrimSuffix(msg, "\n")

		data := map[string][]string{"message": {msg},"name" : {name}}

		_ , err = http.PostForm("http://localhost:6969/post",data)
		if err != nil {
			log.Fatal(err)
		}

	}
}

func Listen () {
	http.HandleFunc ("/broadcast",GetMessage)

	log.Print("\u001B[92mchat joined\u001B[0m")
	if err := http.ListenAndServe(":4242",nil); err != nil {
		log.Fatal(err)
	}
}

func GetMessage (writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		if err := request.ParseForm(); err != nil {
			log.Fatal(err)
		}

	log.Print(request.PostForm["message"][0])

	} else {
		panic("only post method is allowed")
	}
}