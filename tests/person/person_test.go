package person_test

import (
	"autotests-go-api/clients"
	"autotests-go-api/generators"
	"autotests-go-api/helpers"
	"testing"
)

// Отображение данных персонажей
func TestGetPerson_ExistingPersonId_RetunsCorrectData(t *testing.T) {
	// Arrange
	testCases := []struct {
		name     string
		personId int
		jsonFile string
	}{
		{
			name:     "Luke Skywalker",
			personId: 1,
			jsonFile: "lukeSkywalker.json",
		},
		{
			name:     "Darth Vader",
			personId: 4,
			jsonFile: "darthVader.json",
		},
		{
			name:     "Obi-Wan Kenobi",
			personId: 10,
			jsonFile: "obiWanKenobi.json",
		},
	}

	// Act
	// Assert
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ComparePersonWithExpected(t, tc.personId, tc.jsonFile)
		})
	}
}

// Отображение ошибки при получении персонажа с указанием несуществующего id в поле id
func TestGetPerson_UnexistingPersonId_Returns404Code(t *testing.T) {
	// Arrange
	var personId = 84

	// Act
	person, _ := clients.GetPerson(personId, 404)

	// Assert
	helpers.IsNil(person, t)
}

// Отображение ошибки при получении персонажа с указанием строкового значения в поле id
func TestGetPerson_StringValuePersonId_Returns404Code(t *testing.T) {
	// Arrange
	var personId = generators.RandomLatinString(1, 20)

	// Act
	person, _ := clients.GetPerson(personId, 404)

	// Assert
	helpers.IsNil(person, t)
}

// Отображение ошибки при получении персонажа с указанием bool значения в поле id
func TestGetPerson_BoolValuePersonId_Returns404Code(t *testing.T) {
	// Arrange
	var personId = generators.GetRandomBool()

	// Act
	person, _ := clients.GetPerson(personId, 404)

	// Assert
	helpers.IsNil(person, t)
}

// Получить первонажа из testData, получить персонажа из backend'a и сравнить их между собой
//
// Параметры:
//   - t: экземпляр testing.T для сообщения об ошибках
//   - personId: Фактическое полученное значение
//   - fileName: Название json-файла с данными целевого персонажа
func ComparePersonWithExpected(t *testing.T, personId int, fileName string) {
	expectedPerson := helpers.LoadPersonFromJSON(fileName)

	person, err := clients.GetPerson(personId)
	if err != nil {
		t.Fatalf("Ошибка при запросе GetPerson: %v", err)
	}

	helpers.AreDeepEqual(expectedPerson, person, t)
}
