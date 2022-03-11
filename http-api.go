package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Test struct {
	Data string
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/test", test)

	http.ListenAndServe(":5000", router)
}

func test(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("This is the test"))
	json.NewEncoder(w).Encode(struct {
		ID string
	}{ID: fmt.Sprintf("%d", time.Now().UnixNano())})
}
