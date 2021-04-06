package main

import (
	"log"
	"net/http"
	"encoding/json"
	"fmt"
	"bytes"
	"io/ioutil"
)

const HOST_NAME = "http://localhost:6969/"

func main () {
	var name string
	
	fmt.Print("enter name : ")
	fmt.Scanln(&name)
	
	reqBody, err := json.Marshal (map[string]string {
		"name" : name, 
	})

	if err != nil {
		log.Fatal(err)
	}

	response , err := http.Post("http://localhost:6969/login","application/json",bytes.NewBuffer(reqBody))

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		print(err)
	}
	fmt.Println(string(body))
}