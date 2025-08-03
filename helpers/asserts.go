package asserts

import (
	"reflect"
	"testing"
)

func AreDeepEqual(fieldName string, expected, actual any, t *testing.T) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Отображается некорректное значение в %s:\nОжидалось: %+v\nПолучено: %+v", fieldName, expected, actual)
	}
}
