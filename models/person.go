
package models

import ("time")

var  _map =map[string]*Person{}

type Person struct {  
    Id  string  `json:"id"`
    Name   string `json:"name"`
    LastName    string `json:"lastname"`
    Birthday    time.Time `json:"birthday"`
    Weight  float64 `json:"weight"`
    Height float64 `json:"height"`
} 

func init() {
}

func (b *Person) AddPerson() *Person {
  _map[b.Id] = b
	return b
}

func (b *Person) UpdatePerson() *Person {
  _map[b.Id] = b
	return b
}

func GetPersonById(Id string) (*Person){
	person := _map[Id]
	return person
}


func GetPersonByHeight(height float64) ([]Person){
  var persons []Person
  for _, v := range _map {
    if v.Height == height { 
     persons = append(persons, *v)
    }
  }
  return persons
}

func GetPersonByWeight(weight float64) ([]Person){
	var persons []Person
  for _, v := range _map {
    if v.Weight == weight { 
     persons = append(persons, *v)
    }
  }
  return persons
}

func DeletePerson(Id string) (*Person) {
	person := _map[Id]
  delete(_map,Id )
	return person
}