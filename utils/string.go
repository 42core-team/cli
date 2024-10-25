package utils

import "fmt"

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
