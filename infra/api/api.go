package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go-api-test/modules/person"

	"github.com/gorilla/mux"
)

var People []person.Person

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(People)
}
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range People {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&person.Person{})
}
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person person.Person
	json.NewDecoder(r.Body).Decode(&person)
	person.ID = fmt.Sprint(len(People) + 1)
	People = append(People, person)
}
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range People {
		if item.ID == params["id"] {
			People = append(People[:index], People[index+1:]...)
			return
		}
	}
	json.NewEncoder(w).Encode(People)
}

func Init() *mux.Router {
	var router = mux.NewRouter()
	router.HandleFunc("/get", GetPeople).Methods("GET")
	router.HandleFunc("/find/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/create", CreatePerson).Methods("POST")
	router.HandleFunc("/delete/{id}", DeletePerson).Methods("DELETE")
	return router
}
