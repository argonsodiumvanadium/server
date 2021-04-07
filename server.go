	package main

import (
	"log"
	"net/http"
	"math/rand"
	"strings"
)

type (
	LogInManager struct {
		members []*Member
		colorMap map[string]*Member
	}

	Member struct {
		color string
		name string
		URL string
	}
)

var (
	colors = [ ... ]string{"\u001B[91m","\u001B[92m","\u001B[96m","\u001B[95m"}
	loginManager = &LogInManager{members : make([]*Member,0),colorMap : make(map[string]*Member)}
)

const (
	RESET = "\u001B[0m"
)

func main () {
	http.HandleFunc ("/post",PostMessage)

	http.Handle ("/login",loginManager)

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

	msg := request.PostForm["message"][0]
	name := request.PostForm["name"][0]

	msg = ParseMessage(msg,name)

	log.Println("msg recv",msg)

	for i , member := range loginManager.members {
		data := map[string][]string{"message": {msg}}
		_, err := http.PostForm(member.URL,data)

		if err != nil {
			log.Println("\u001B[91m"+member.name+" left\u001B[0m")
			loginManager.RemoveMemberAt(i)
		}
	}


		
	} else {
		panic("only post method is allowed")
	}
}

func (self *LogInManager) RemoveMemberAt (i int) {
	self.members[len(self.members)-1], self.members[i] = self.members[i], self.members[len(self.members)-1]
	self.members = self.members[:len(self.members)-1]
}

func ParseMessage (msg,name string) (message string) {
	color := loginManager.colorMap[name].color

	message = color + name + RESET + " : " + msg + "\n"

	return
}

func (self *LogInManager) ServeHTTP (writer http.ResponseWriter, request *http.Request) {
	log.Print("user entered")
	if request.Method == "POST" {
		if err := request.ParseForm(); err != nil {
			log.Fatal(err)
		}
	
	newUser := request.PostForm["name"][0]

	writer.Header().Add("Content-Type", "application/json")
	ip := GetIP(request)

	log.Println(ip)

	member := &Member {color : colors[rand.Intn(len(colors))],name : newUser,URL : "http://" + ip + ":4242/broadcast"}

	self.members = append(self.members, member)
	self.colorMap[newUser] = member

	} else {
		panic("only post method is allowed")
	}
}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return strings.Split(r.RemoteAddr,":")[0]
}