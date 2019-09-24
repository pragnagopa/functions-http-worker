package main
import (
    "fmt"
    "log"
    "net/http"
    "time"
    "encoding/json"
)
type Profile struct {
    Name    string
    Hobbies []string
  }
func queueTriggerHandler(w http.ResponseWriter, r *http.Request) {
    for k, v := range r.Header {
        fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
    }
    queueItemArray := r.Header["Myqueueitem"]
    queueItem := ""
    for _, element := range queueItemArray {
        queueItem = queueItem + element
    }
    w.Write([]byte("Go server returning Myqueueitem: " + queueItem))
  }
  
  func httpTriggerHandler(w http.ResponseWriter, r *http.Request) {
    t := time.Now()
    fmt.Println(t.Month())
    fmt.Println(t.Day())
    fmt.Println(t.Year())
    //w.Write([]byte("Hello World from go worker:pgopa"))
    profile := Profile{"Alex", []string{"snowboarding", "programming"}}

  js, err := json.Marshal(profile)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
  }

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/HttpTrigger", httpTriggerHandler)
    mux.HandleFunc("/QueueTrigger", queueTriggerHandler)
    log.Println("Go server Listening...")
    log.Fatal(http.ListenAndServe(":8090", mux))
}