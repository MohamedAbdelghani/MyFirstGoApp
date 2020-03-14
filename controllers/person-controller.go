package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mohamedabdelghani/myfirstgoapp/models"
	"github.com/mohamedabdelghani/myfirstgoapp/utils"
)

var person models.Person

func Add(w http.ResponseWriter, r *http.Request) {
	person := &models.Person{}
	utils.ParseBody(r, person)
	b := person.AddPerson()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personId := vars["id"]
	Id, err := strconv.ParseInt(personId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	personDetails := models.GetPersonById(Id)
	res, _ := json.Marshal(personDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetByWeight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	weight := vars["weight"]
	Weight, err := strconv.ParseFloat(weight, 64)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	personDetails := models.GetPersonByWeight(Weight)
	res, _ := json.Marshal(personDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetByHeight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	height := vars["height"]
	Height, err := strconv.ParseFloat(height, 64)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	personDetails := models.GetPersonByHeight(Height)
	res, _ := json.Marshal(personDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var updatePerson = &models.Person{}
	utils.ParseBody(r, updatePerson)
	vars := mux.Vars(r)
	personId := vars["id"]
	Id, err := strconv.ParseInt(personId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	personDetails := models.GetPersonById(Id)

	if updatePerson.Name != "" {
		personDetails.Name = updatePerson.Name
	}
	if updatePerson.LastName != "" {
		personDetails.LastName = updatePerson.LastName
	}
	if !updatePerson.Birthdate.IsZero() {
		personDetails.Birthdate = updatePerson.Birthdate
	}
	if updatePerson.Weight > 0 {
		personDetails.Weight = updatePerson.Weight
	}
	if updatePerson.Height > 0 {
		personDetails.Height = updatePerson.Height
	}
	// save data to file (should abstract this logic away from controller)
	file, _ := json.MarshalIndent(updatePerson, "", " ")
	_ = ioutil.WriteFile("logs/persons.json", file, 0644)

	personDetails.UpdatePerson()
	res, _ := json.Marshal(personDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personId := vars["id"]
	Id, err := strconv.ParseInt(personId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	person := models.DeletePerson(Id)
	res, _ := json.Marshal(person)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
