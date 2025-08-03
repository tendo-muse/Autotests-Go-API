package person_test

import (
	"autotests-go-api/clients"
	asserts "autotests-go-api/helpers"
	"autotests-go-api/models"
	"testing"
)

// Отображение данных персонажа Luke Skywalker (ID=1)
func TestGetPerson(t *testing.T) {
	// Arrange
	var personId = 1
	expectedPerson := models.Person{
		Name:      "Luke Skywalker",
		Height:    "172",
		Mass:      "77",
		HairColor: "blond",
		SkinColor: "fair",
		EyeColor:  "blue",
		BirthYear: "19BBY",
		Gender:    "male",
		Homeworld: "https://swapi.info/api/planets/1",
		Films: []string{
			"https://swapi.info/api/films/1",
			"https://swapi.info/api/films/2",
			"https://swapi.info/api/films/3",
			"https://swapi.info/api/films/6",
		},
		Species: []string{},
		Vehicles: []string{
			"https://swapi.info/api/vehicles/14",
			"https://swapi.info/api/vehicles/30",
		},
		Starships: []string{
			"https://swapi.info/api/starships/12",
			"https://swapi.info/api/starships/22",
		},
		Created: "2014-12-09T13:50:51.644000Z",
		Edited:  "2014-12-20T21:17:56.891000Z",
		URL:     "https://swapi.info/api/people/1",
	}

	// Act
	person, err := clients.GetPerson(personId)
	if err != nil {
		t.Fatalf("GetPerson failed: %v", err)
	}

	// Assert
	asserts.AreDeepEqual("Persone Luke Skywalker", expectedPerson, *person, t)
}
