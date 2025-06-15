package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/welcome", welcomeHandler)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/user" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}
	if err:=r.ParseForm(); err!=nil {
		fmt.Fprintln(w,"Parse form failed",err)
		return
	}
	fmt.Fprintln(w, "POST request successful")
	username:=r.FormValue("username")
	addr:=r.FormValue("addr")
	fmt.Fprintln(w,"Name =",username)
	fmt.Fprintln(w,"Address =",addr)
	

}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/welcome" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}
	/* trunk-ignore(golangci-lint2/errcheck) */
	fmt.Fprintln(w, "Welcome to the static website")
}
