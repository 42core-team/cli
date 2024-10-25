package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func TrimStringWithIndicator(input string, maxLength int, indicator string) string {
	if len(input) <= maxLength {
		return input
	}

	indicatorLength := len(indicator)
	if indicatorLength >= maxLength {
		return indicator[:maxLength]
	}

	halfLength := (maxLength - indicatorLength) / 2
	start := input[:halfLength]
	end := input[len(input)-halfLength:]

	return fmt.Sprintf("%s%s%s", start, indicator, end)
}

func ExtractWinner(logs string) (int, error) {
	re := regexp.MustCompile(`Game over, winner is team (\d*)`)
	matches := re.FindStringSubmatch(logs)
	if len(matches) < 2 {
		return 0, errors.New("winner not found in logs")
	}
	if matches[1] == "" {
		return 0, errors.New("winner number is missing in logs")
	}
	winner, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, err
	}
	return winner, nil
}
