package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/invokeFunction", handler)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
    //Iterate over all header fields
    //for k, v := range r.Header {
    //    fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
    //}

    //fmt.Fprintf(w, "Host = %q\n", r.Host)
    //fmt.Fprintf(w, "RemoteAddr= %q\n", r.RemoteAddr)
    //Get value for a specified token
    fmt.Fprintf(w, "Go HttpServer of Myqueueitem", r.Header["Myqueueitem"])
    w.Write([]byte("HTTP status code returned!"))
}