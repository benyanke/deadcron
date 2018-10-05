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
type Checks map[string]Check

type Payload struct {
	Stuff Data
}
type Data struct {
	Actions Actions
}

type Actions map[string]string
type Wrapper map[string]Actions


// Wrapper for the JSON Marshal function
// Returns a panic if the marshal was not
// a success, to ensure there's no unhandled
// errors
func JsonMustMarshal(data interface{}) []byte {
    out, err := json.Marshal(data)
    if err != nil {
        panic(err)
    }

    return out
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {

	Wrapper := make(map[string]Actions)
	Actions := make(map[string]string)

	Actions["/check"] = "Checkin from your host"
	Actions["/status"] = "View status of all checks"
	Actions["/"] = "dis."

	Wrapper["actions"] = Actions

	jData := JsonMustMarshal(Wrapper)

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

	checkId, found := mux.Vars(r)["checkId"]

	if !found {
		fmt.Fprintf(w, "No check ID specified. Can not continue")
		return
	}
	fmt.Fprintf(w, "CHECK "+checkId+" HAS BEEN HANDLED ")

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

	router.HandleFunc("/check", checkHandler).Methods("GET")
	router.HandleFunc("/check/", checkHandler).Methods("GET")
	router.HandleFunc("/check/{checkId}", checkHandler).Methods("GET")

	router.HandleFunc("/status", statusHandler).Methods("GET")
	router.HandleFunc("/", defaultHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", router))
}
