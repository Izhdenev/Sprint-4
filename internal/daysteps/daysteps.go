package daysteps

import (
	"errors"
	"fmt"
	"log"
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
		return 0, 0, fmt.Errorf("")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("")
	}

	if steps <= 0 {
		return 0, 0, errors.New("")
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
		log.Print("Неверный вес")
		return ""
	}
	if height <= 0 {
		log.Print("Неверный рост")
		return ""
	}

	if data == "" {
		log.Print("Неверная длительность")
		return ""
	}

	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Print("Неверное время ")
		return fmt.Sprintf("")
	}

	distanceMeters := float64(steps) * stepLength
	distanceKm := distanceMeters / mInKm

	calories, errCalories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if errCalories != nil {
		return ""
	}

	result := fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		steps, distanceKm, calories)
	return result
}
