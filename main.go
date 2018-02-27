package main

import (
    "fmt"
    "net/http"
    "os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
 fmt.Fprintf(w, "Welcome to the HomePage!")
}

func homePage_two(w http.ResponseWriter, r *http.Request) {
 fmt.Fprintf(w,"Welcome to the HomePageTwo!")
}

func homePage_three(w http.ResponseWriter, r *http.Request) {
 fmt.Fprintf(w,"Welcome to the HomePageThree!")
}

func main() {
    fmt.Println("listening...")
         http.HandleFunc("/a", homePage)
         http.HandleFunc("/b", homePage_two)
         http.HandleFunc("/c", homePage_three)
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
        panic(err)
    }
}
