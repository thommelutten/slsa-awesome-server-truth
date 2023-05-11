package main

import (
    "fmt"
    "net/http"
)

func awesome(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "Who's awesome? You're awesome!\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

    http.HandleFunc("/awesome", awesome)
    http.HandleFunc("/headers", headers)

    http.ListenAndServe(":8090", nil)
}