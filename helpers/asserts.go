package helpers

import (
	"reflect"
	"testing"
)

// Сравнение двух значений любого типа на глубокое равенство с помощью reflect.DeepEqual
//
// Параметры:
//   - expected: Ожидаемое значение
//   - actual: Фактическое полученное значение
//   - t: экземпляр testing.T для сообщения об ошибках
//
// Возвращает детализированное сообщение об ошибке, если значения не совпадают
func AreDeepEqual(expected, actual any, t *testing.T) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Отображается некорректное значение в:\nОжидалось: %+v\nПолучено : %+v", expected, actual)
	}
}

// Проверка, является ли значение nil
//
// Параметры:
//   - actual: Фактическое полученное значение
//   - t: экземпляр testing.T для сообщения об ошибках
//
// Возвращает детализированное сообщение об ошибке, если значения не совпадают
func IsNil(actual any, t *testing.T) {
	if actual == nil {
		t.Errorf("Ожидалось: nil, nil\nПолучено: %v, %v", nil, actual)
	}
}
