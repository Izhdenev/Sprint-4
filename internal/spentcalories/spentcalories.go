package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	parts := strings.Split(data, ",")

	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("")
	}
	if steps <= 0 {
		return 0, "", 0, errors.New("")
	}
	activityType := parts[1]

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("")
	}
	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("")
	}
	return steps, activityType, duration, nil
}

func distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	totalDistanceMeters := float64(steps) * stepLength
	distanceKm := totalDistanceMeters / float64(mInKm)
	return distanceKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	distKm := distance(steps, height)
	durationHours := duration.Hours()

	if durationHours <= 0 {
		return 0
	}

	speed := distKm / float64(durationHours)
	return speed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activityType, duration, err := parseTraining(data)
	if err != nil {
		return "", fmt.Errorf("")
	}

	distKm := distance(steps, height)
	speed := meanSpeed(steps, height, duration)

	var calories float64
	var calErr error

	switch activityType {
	case "Ходьба":
		calories, calErr = WalkingSpentCalories(steps, weight, height, duration)
	case "Бег":
		calories, calErr = RunningSpentCalories(steps, weight, height, duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	if calErr != nil {
		return "", fmt.Errorf("")
	}
	result := fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		activityType, duration.Hours(), distKm, speed, calories)

	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("неправильные параметры(бег)")
	}

	speed := meanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()

	runCalories := (weight * speed * durationMinutes) / minInH
	return runCalories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("неправильные параметры (шаг)")
	}

	speed := meanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()

	baseCalories := (weight * speed * durationMinutes) / minInH
	walkCalories := baseCalories * walkingCaloriesCoefficient

	return walkCalories, nil
}
