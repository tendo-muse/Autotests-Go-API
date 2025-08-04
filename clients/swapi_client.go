package clients

import (
	"autotests-go-api/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "https://swapi.info/api/people/"

// Получить персонажа по id
//
// Параметры:
//   - personId: id персонажа
//   - expectedStatusCodes: Ожидаемый статус-код
//
// Возвращает данные персонажа и детализированное сообщение об ошибке, если http-запрос упал с ошибкой
func GetPerson(personId any, expectedStatusCodes ...int) (*models.Person, error) {
	// Установка значения по умолчанию, если оно отсутствует
	expectedStatusCode := 200

	switch len(expectedStatusCodes) {
	case 0:
		break
	case 1:
		expectedStatusCode = expectedStatusCodes[0]
	default:
		return nil, fmt.Errorf("нет поддержки на количество статус-кодов: %v (ожидается 0 или 1)", expectedStatusCodes)
	}

	response, err := http.Get(fmt.Sprintf("%s%d", baseURL, personId))

	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе: %v", err)
	}

	if response.StatusCode != expectedStatusCode {
		return nil, fmt.Errorf("не совпал ожидаемый статус-код: %d (ожидался %d)", response.StatusCode, expectedStatusCode)
	}

	// Закрытие данных после их использования
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать тело ответа: %w", err)
	}

	var person models.Person
	if err := json.Unmarshal(body, &person); err != nil {
		return nil, fmt.Errorf("не удалось распарсить тело ответа: %w", err)
	}

	return &person, nil
}
