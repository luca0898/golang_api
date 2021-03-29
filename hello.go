package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type event struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type allEvents []event

var events = allEvents{}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

func findOneEvent(res http.ResponseWriter, req *http.Request) {
	eventId := mux.Vars(req)["id"]

	for _, s := range events {
		if s.ID == eventId {
			res.WriteHeader(http.StatusOK)
			json.NewEncoder(res).Encode(s)
			break
		}
	}
}

func findAllEvents(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(events)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/event", findAllEvents).Methods("GET")
	router.HandleFunc("/event/{id}", findOneEvent).Methods("GET")
	router.HandleFunc("/event", createEvent).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
