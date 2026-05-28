package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

// Основные константы, необходимые для расчетов.

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")

	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("неверный формат строки")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка преобразования количества шагов")
	}

	if steps <= 0 {
		return 0, 0, errors.New("количество шагов должно быть больше 0")
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("")
	}
	if duration <= 0 {
		return 0, 0, fmt.Errorf("")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {

	if weight <= 0 {
		return fmt.Sprintf("Ошибка: вес должен быть больше 0")
	}
	if height <= 0 {
		return fmt.Sprintf("Ошибка: рост должен быть больше 0")
	}

	// Проверка на пустую строку
	if data == "" {
		return fmt.Sprintf("Ошибка: входные данные не могут быть пустыми")
	}

	steps, duration, err := parsePackage(data)
	if err != nil {

		return fmt.Sprintf("Ошибка обработки данных: %v", err)
	}

	distanceMeters := float64(steps) * stepLength
	distanceKm := distanceMeters / mInKm

	calories, errCalories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if errCalories != nil {
		return fmt.Sprintf("Ошибка расчёта калорий: %v", errCalories)
	}

	result := fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		steps, distanceKm, calories)
	return result
}
