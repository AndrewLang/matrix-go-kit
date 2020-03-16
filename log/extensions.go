package log

import (
	"strconv"
	"strings"
)

// Join concate given strings
func Join(values ...string) string {
	var builder strings.Builder

	for _, value := range values {
		builder.WriteString(value)
	}

	return builder.String()
}

// JoinStrings join
func JoinStrings(values []string, spearator string) string {
	var builder strings.Builder
	length := len(values)

	for index, value := range values {
		if IsNullOrEmpty(value) {
			continue
		}

		builder.WriteString(value)
		if length > 1 && index < length-1 {
			builder.WriteString(spearator)
		}
	}

	return builder.String()
}

// JoinWith join
func JoinWith(spearator string, values ...string) string {
	var builder strings.Builder
	length := len(values)

	for index, value := range values {
		if IsNullOrEmpty(value) {
			continue
		}

		builder.WriteString(value)
		if length > 1 && index < length-1 {
			builder.WriteString(spearator)
		}
	}

	return builder.String()
}

// PaddingRight padding right
func PaddingRight(original string, place string, maxLen int) string {
	builder := NewStringBuilder()
	formatter := Formatter{}
	length := len(original)

	builder.Append(original)
	if length < maxLen {
		builder.Append(formatter.Compose(place, maxLen-length))
	}

	return builder.String()
}

// NotNullOrEmpty return true if a string is NOT null or empty
func NotNullOrEmpty(value string) bool {
	return len(value) != 0
}

// IsNullOrEmpty return true if a string is nil or empty
func IsNullOrEmpty(value string) bool {
	return len(value) == 0
}

// IsNumber check whether a string can be converted to int number
func IsNumber(value string) bool {
	if _, err := strconv.Atoi(value); err == nil {
		return true
	}

	return false
}

// ToNumber Convert string to int
func ToNumber(value string) (int, error) {

	num, err := strconv.Atoi(value)

	return num, err
}
