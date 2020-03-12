package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
  "models"
	"github.com/gorilla/mux"
	"net/http"
  "io/ioutil"
)

var person models.Person

func Add(w http.ResponseWriter, r *http.Request) {
	person := &models.Person{}
	utils.ParseBody(r, person)
	b:= person.AddPerson()
	res,_ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personId := vars["id"]
	personDetails, _:= models.GetPersonById(personId)
	res, _ := json.Marshal(personDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func GetByWeight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	weight := vars["weight"]
  Weight, err:= strconv.ParseFloat64(weight, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	personDetails, _:= models.GetPersonByWeight(Weight)
	res, _ := json.Marshal(personDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetByHeight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	height := vars["height"]
  Height, err:= strconv.ParseFloat64(weight, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	personDetails, _:= models.GetPersonByHeight(Height)
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
	personDetails, db:= models.GetPersonById(personId)

	if updatePerson.Name != "" {
		personDetails.Name = updatePerson.Name
	}
	if updatePerson.LastName != "" {
		personDetails.Height = updatePerson.Height
	}
	if updatePerson.Birthday!= "" {
		personDetails.Height = updatePerson.Height
	}
  if updatePerson.Weight!= "" {
		personDetails.Height = updatePerson.Height
	}
  if updatePerson.Height!= "" {
		personDetails.Height = updatePerson.Height
	}
// save data to file (should abstract this logic away from controller)
 file, _ := json.MarshalIndent(updatePerson, "", " ")
	_ = ioutil.WriteFile("logs/persons.json", file, 0644)
  
  b:= personDetails.UpdatePerson()
	res, _ := json.Marshal(personDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personId := vars["id"]
	person:= models.DeletePerson(personId)
	res, _ := json.Marshal(person)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}