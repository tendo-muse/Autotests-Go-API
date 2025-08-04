package generators

import (
	"math/rand"
	"time"
)

// Получить случайное bool-значение
//
// Возвращает случайное bool-значение на основе времени
func GetRandomBool() bool {
	rand.Seed(time.Now().UnixNano()) // Оставить! Иначе будет всегда false
	randomBool := rand.Intn(2) == 1
	return randomBool
}

// Получить случайное string-значение
//
// Параметры:
//   - minLength: Минимальная длина строки (включительно)
//   - maxLength: Максимальная длина строки (включительно)
//
// Возвращает случайное bool-значение на основе времени
func RandomLatinString(minLength, maxLength int) string {
	rand.Seed(time.Now().UnixNano())

	if minLength < 1 {
		minLength = 1
	}
	if maxLength < minLength {
		maxLength = minLength
	}

	length := minLength + rand.Intn(maxLength-minLength+1)
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		charType := rand.Intn(2) // 0 or 1

		if charType == 0 {
			result[i] = byte('A' + rand.Intn(26)) // Заглавные A-Z
		} else {
			result[i] = byte('a' + rand.Intn(26)) // Строчные a-z
		}
	}

	return string(result)
}
