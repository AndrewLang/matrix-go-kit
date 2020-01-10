package logging

import (
	"strings"
)

// StringBuilder string builder use to build string with high performance
type StringBuilder struct {
	internalBuilder *strings.Builder
	Lines           int
}

// NewStringBuilder create a new string builder instance
func NewStringBuilder() *StringBuilder {
	var rawBuilder strings.Builder
	builder := new(StringBuilder)
	builder.Lines = 0
	builder.internalBuilder = &rawBuilder
	return builder
}

// Append append string value
func (b *StringBuilder) Append(value string) *StringBuilder {
	b.internalBuilder.WriteString(value)
	b.Lines = 1
	return b
}

// AppendLine append a line of string
func (b *StringBuilder) AppendLine(value string) *StringBuilder {
	b.internalBuilder.WriteString(value)
	b.internalBuilder.WriteString(NewLine)
	b.Lines++
	return b
}

// Reset reset to empty
func (b *StringBuilder) Reset() *StringBuilder {
	b.internalBuilder.Reset()
	b.Lines = 0
	return b
}

// String convert string builder to string
func (b *StringBuilder) String() string {
	return b.internalBuilder.String()
}
