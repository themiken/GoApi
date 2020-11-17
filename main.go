
// main.go
package main

import (    
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Miken api works")
    fmt.Println("Test")
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
    handleRequests()
}