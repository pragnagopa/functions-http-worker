package main

import (
    "fmt"
    "log"
    "net/http"
)

func queueTriggerHandler(w http.ResponseWriter, r *http.Request) {
     //fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
    //Iterate over all header fields
    for k, v := range r.Header {
        fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
    }

    //fmt.Fprintf(w, "Host = %q\n", r.Host)
    //fmt.Fprintf(w, "RemoteAddr= %q\n", r.RemoteAddr)
    //Get value for a specified token
    fmt.Fprintf(w, "Go HttpServer of Myqueueitem", r.Header["Myqueueitem"])
    w.Write([]byte("HTTP status code returned!"))
  }
  
  func httpTriggerHandler(w http.ResponseWriter, r *http.Request) {
    queryParamName := r.URL.Query().Get("name")
	w.Write([]byte("Hello World from go worker:" + queryParamName))
  }

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/invoke/httpTriggerHandler", httpTriggerHandler)
    mux.HandleFunc("/invoke/queueTriggerHandler", queueTriggerHandler)
    log.Println("Go server Listening...")
    log.Fatal(http.ListenAndServe(":8000", mux))
}