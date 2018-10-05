package main

import (
	"fmt"
	"net/http"

	"encoding/json"
	"github.com/gorilla/mux"
	"log"
)

type Check struct {
	ID               string
	Token            string
	ExpectedInterval int
	GracePeriod      int
}

type Payload struct {
	Stuff Data
}
type Data struct {
	Actions Actions
}

type Actions map[string]string
type Wrapper map[string]Actions

func defaultHandler(w http.ResponseWriter, r *http.Request) {

	Wrapper := make(map[string]Actions)
	Actions := make(map[string]string)

	Actions["/check"] = "Checkin from your host"
	Actions["/status"] = "View status of all checks"
	Actions["/"] = "dis."

	Wrapper["actions"] = Actions

	jData, err := json.Marshal(Wrapper)
	if err != nil {
		fmt.Fprintf(w, "JSON ERROR FAM")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)

	/*
	   		fruits := make(map[string]int)
	       fruits["Apples"] = 25
	       fruits["Oranges"] = 10

	       vegetables := make(map[string]int)
	       vegetables["Carrats"] = 10
	       vegetables["Beets"] = 0

	       d := Data{fruits, vegetables}
	       p := Payload{d}

	       return json.MarshalIndent(p, "", "  ")
	*/
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CHECK HAS BEEN HANDLED")
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ALL OF YOUR CHECKS ARE BELONG TO US")
}

// Not working yet
func loggingHandler(w http.Handler) {
	fmt.Fprint(nil, "LOGGED AF")
}

// Old main
func mainOld() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/check", checkHandler)
	http.HandleFunc("/status", statusHandler)

	log.Print("Starting up on :8080, fam")
	log.Fatal(http.ListenAndServe(":8080", nil))
	// loggingHandler))
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/check/{checkId}", checkHandler).Methods("GET")
	router.HandleFunc("/status", statusHandler).Methods("GET")
	router.HandleFunc("/", defaultHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", router))
}
