package services

import (
	"fmt"
	"regexp"
	"strconv"
)

// currrentRam, maxRam, error
func parseRam(ramStr string) (string, string, error) {
	// Регулярное выражение для поиска двух чисел, разделенных слэшем
	re := regexp.MustCompile(`(\d+)/(\d+) MB`)

	// Применяем регулярное выражение
	matches := re.FindStringSubmatch(ramStr)
	if len(matches) != 3 {
		return "", "", fmt.Errorf("invalid format")
	}

	// Преобразуем найденные строки в числа
	currentRam, err1 := strconv.Atoi(matches[1])
	maxRam, err2 := strconv.Atoi(matches[2])
	if err1 != nil || err2 != nil {
		return "", "", fmt.Errorf("error converting values to integers")
	}

	// Формируем строку вида "847/961"
	return fmt.Sprintf("%d", currentRam), fmt.Sprintf("%d", maxRam), nil
}

func parseMemory(memoryStr string) (string, string, error) {
	// Регулярное выражение для поиска двух чисел с плавающей точкой, разделенных пробелами и словом "of"
	re := regexp.MustCompile(`Used (\d+\.\d+) GB of (\d+\.\d+) GB`)

	// Применяем регулярное выражение
	matches := re.FindStringSubmatch(memoryStr)
	if len(matches) != 3 {
		return "", "", fmt.Errorf("invalid format")
	}

	// Возвращаем строковые значения для текущего и максимального объема памяти
	return matches[1], matches[2], nil
}
