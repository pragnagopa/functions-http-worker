package main
import (
    "fmt"
    "log"
    "net/http"
    "time"
    "encoding/json"
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

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/HttpTrigger", httpTriggerHandler)
    mux.HandleFunc("/QueueTrigger", queueTriggerHandler)
    log.Println("Go server Listening...")
    log.Fatal(http.ListenAndServe(":8090", mux))
}