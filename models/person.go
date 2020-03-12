
package models

var map _map[string]Person

type Person struct {  
    Id  string  `json:"id"`
    Name   string `json:"name"`
    LastName    string `json:"lastname"`
    Birthday    time.Time `json:"birthday"`
    Weight  float64 `json:"weight"`
    Height float64 `json:"height"`
} 

func init() {
_map = make(map[string]Person)
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


func GetPersonByHeight(height string) (Person[]){
  var persons []Person
  for k, v := range _map {
    if v.Height == height { 
     persons = append(persons, v)
    }
  }
  return persons
}

func GetPersonByWeight(weight string) (*Person[]){
	  var persons []Person
  for k, v := range _map {
    if v.Weight == weight { 
     persons = append(persons, v)
    }
  }
  return persons
}

func DeletePerson(Id string) Person {
	var person Person
	person := _map[Id]
  delete(_map,Id )
	return person
}