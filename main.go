package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello world!!\n")

       /*  fmt.Println("row1")
        fmt.Print("row2")
        fmt.Println("row3")
        fmt.Print("row4") */
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
