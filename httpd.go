package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	db     = map[string]interface{}{}
	dbLock sync.Mutex
)

type Entry struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

// type MathRequest struct {
// 	Op    string  `json:"op"`
// 	Left  float64 `json:"left"`
// 	Right float64 `json:"right"`
// }
//
// type MathResponse struct {
// 	Error  string  `json: "error"`
// 	Result float64 `json: "result"`
// }

func sendResponse(entry *Entry, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(entry); err != nil {
		log.Printf("error encoding %+v - %s", entry, err)
	}
}

func kvPostHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	entry := &Entry{}
	if err := dec.Decode(entry); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dbLock.Lock()
	defer dbLock.Unlock()
	db[entry.Key] = entry.Value

	sendResponse(entry, w)
}

func kvGetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[4:]

	dbLock.Lock()
	defer dbLock.Unlock()
	value, ok := db[key]
	if !ok {
		http.Error(w, fmt.Sprintf("Key %q not found.", key), http.StatusNotFound)
		return
	}

	entry := &Entry{
		Key:   key,
		Value: value,
	}
	sendResponse(entry, w)
}

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello! Gophers.")
// }
//
// func mathHandler(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
//
// 	dec := json.NewDecoder(r.Body)
// 	req := &MathRequest{}
// 	if err := dec.Decode(req); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
//
// 	resp := &MathResponse{}
// 	switch req.Op {
// 	case "+":
// 		resp.Result = req.Left + req.Right
// 	case "-":
// 		resp.Result = req.Left - req.Right
// 	case "*":
// 		resp.Result = req.Left * req.Right
// 	case "/":
// 		if req.Right == 0.0 {
// 			resp.Error = "divided by zero"
// 		} else {
// 			resp.Result = req.Left / req.Right
// 		}
// 	default:
// 		resp.Error = fmt.Sprintf("unknown Operation: %s", req.Op)
// 	}
//
// 	w.Header().Set("Content-Type", "application/json")
// 	if resp.Error != "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 	}
//
// 	enc := json.NewEncoder(w)
// 	if err := enc.Encode(resp); err != nil {
// 		log.Printf("can`t encode %v - %s", resp, err)
// 	}
// }

func main() {
	// http.HandleFunc("/hello", helloHandler)
	// http.HandleFunc("/math", mathHandler)
	http.HandleFunc("/db", kvPostHandler)
	http.HandleFunc("/db/", kvGetHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
