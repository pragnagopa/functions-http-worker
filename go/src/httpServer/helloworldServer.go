package main
import (
    "fmt"
    "log"
    "net/http"
    "time"
    "encoding/json"
    "os"
    "flag"
)
type ReturnValue struct {
    Data string
}
type InvokeResponse struct {
  Outputs  map[string]interface{}
	Logs []string
	ReturnValue
}

type InvokeRequest struct {
    Data  map[string]interface{}
    Metadata  map[string]interface{}
}

func queueTriggerHandler(w http.ResponseWriter, r *http.Request) {
    var invokeReq InvokeRequest
    d := json.NewDecoder(r.Body)
    decodeErr := d.Decode(&invokeReq)
    if decodeErr != nil {
    // bad JSON or unrecognized json field
    http.Error(w, decodeErr.Error(), http.StatusBadRequest)
    return
    }
    fmt.Println("The JSON data is:invokeReq metadata......")
    fmt.Println(invokeReq.Metadata)
    fmt.Println("The JSON data is:invokeReq data......")
    fmt.Println(invokeReq.Data)
    

   returnValue := ReturnValue{Data:"HelloWorld"}
	outputs := make(map[string]interface{})
    outputs["output"] = "output from go"
    outputs["output2"] = map[string]interface{}{
    "home": "123-466-799",
    "office": "564-987-654",
    }
    invokeResponse := InvokeResponse{outputs, []string{"test log1", "test log2"},returnValue}

  js, err := json.Marshal(invokeResponse)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}
  
  func httpTriggerHandler(w http.ResponseWriter, r *http.Request) {
    t := time.Now()
    fmt.Println(t.Month())
    fmt.Println(t.Day())
    fmt.Println(t.Year())
    ua := r.Header.Get("User-Agent")
    fmt.Printf("user agent is: %s \n", ua)
    invocationid := r.Header.Get("X-Azure-Functions-InvocationId")
    fmt.Printf("invocationid is: %s \n", invocationid)
          
    //w.Write([]byte("Hello World from go worker:pgopa"))
    returnValue := ReturnValue{Data:"return val"}
	outputs := make(map[string]interface{})
    outputs["output"] = "Mark Taylor"
    outputs["output2"] = map[string]interface{}{
    "home": "123-466-799",
    "office": "564-987-654",
    }
    invokeResponse := InvokeResponse{outputs, []string{"test log1", "test log2"},returnValue}

  js, err := json.Marshal(invokeResponse)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
  }

  func simpleHttpTriggerHandler(w http.ResponseWriter, r *http.Request) {
    t := time.Now()
    fmt.Println(t.Month())
    fmt.Println(t.Day())
    fmt.Println(t.Year())
    ua := r.Header.Get("User-Agent")
    fmt.Printf("user agent is: %s \n", ua)
    invocationid := r.Header.Get("X-Azure-Functions-InvocationId")
    fmt.Printf("invocationid is: %s \n", invocationid)

    queryParams := r.URL.Query()

    for k, v := range queryParams {
      fmt.Println("k:", k, "v:", v)
    }
          
   w.Write([]byte("Hello World from go worker:pgopa"))
  }

func main() {
  argsWithProg := os.Args
  argsWithoutProg := os.Args[1:]
  //arg := os.Args[3]
  fmt.Println(argsWithProg)
  fmt.Println(argsWithoutProg)
  
  //args := flag.Args()
  var httpInvokerPort string
    flag.StringVar(&httpInvokerPort, "httpInvokerPort", "", "Usage")

    flag.Parse()

    fmt.Println(httpInvokerPort)
    mux := http.NewServeMux()
    mux.HandleFunc("/HttpTrigger", httpTriggerHandler)
    mux.HandleFunc("/QueueTrigger", queueTriggerHandler)
    mux.HandleFunc("/SimpleHttpTrigger", simpleHttpTriggerHandler)
    mux.HandleFunc("/SimpleHttpTriggerWithReturn", simpleHttpTriggerHandler)
    log.Println("Go server Listening...on httpInvokerPort:", httpInvokerPort)
    log.Fatal(http.ListenAndServe(":" + httpInvokerPort, mux))
}