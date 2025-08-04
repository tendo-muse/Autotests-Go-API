package helpers

import (
	"autotests-go-api/models"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// Вернуть абсолютный путь к файлу в папке testdata
//
// Параметры:
//   - filename: Название файла
//
// Возвращает абсолютный путь к файлу в папке testdata
func GetTestDataPath(filename string) string {
	_, currentFile, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filepath.Dir(currentFile))
	fullPath := filepath.Join(baseDir, "testData", filename)

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		log.Fatal("Не удалось получить абсолютный путь к файлу в папке: %w", err)
		return ""
	}

	return fullPath
}

// Загрузить и серилизовать персонажа из testData
//
// Параметры:
//   - filename: Название файла
//
// Возвращает данные пользователя из testData
func LoadPersonFromJSON(filename string) *models.Person {
	// Получаем абсолютный путь к файлу
	path := GetTestDataPath(filepath.Join("persons", filename))

	// Читаем файл
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Не удалось прочитать тело ответа: %w", err)
		return nil
	}

	// Парсим JSON
	var person models.Person
	if err := json.Unmarshal(data, &person); err != nil {
		log.Fatal("Не удалось распарсить тело ответа: %w", err)
		return nil
	}

	return &person
}
