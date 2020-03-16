package logging

import (
	"fmt"
	"strings"
	"time"
)

// Formatter used to format text/string
type Formatter struct {
}

// NewFormatter create new formatter
func NewFormatter() *Formatter {
	return &Formatter{}
}

// Format give parameters
func (f Formatter) Format(format string, param ...interface{}) string {
	return fmt.Sprintf(format, param...)
}

// FormatConsoleStyle build a console style string with color and style settings
func (f Formatter) FormatConsoleStyle(styles ...string) string {
	// fmt.Println("Console style: ", styles, JoinWith(";", styles...))

	builder := NewStringBuilder()
	builder.Append(StyleEsc)
	builder.Append(JoinWith(";", styles...))
	builder.Append(StyleEnd)

	// fmt.Println("Console style: ", builder.String())

	return builder.String()
}

// FormatConsoleOutput Format console output with given styles
func (f Formatter) FormatConsoleOutput(value string, styles ...string) string {
	builder := NewStringBuilder()
	builder.Append(f.FormatConsoleStyle(styles...))
	builder.Append(value)
	builder.Append(ResetAllStyle)

	return builder.String()
}

// Format256Color format 256 color
func (f Formatter) Format256Color(num int) string {
	return fmt.Sprintf("%s%d", Color256Start, num)
}

// FormatBg256Color format 256 background color
func (f Formatter) FormatBg256Color(num int) string {
	return fmt.Sprintf("%s%d", Color256BgStart, num)
}

// FormatError format
func (f Formatter) FormatError(err error) string {
	return fmt.Sprintf("%s", err)
}

// FormatDateTime format given date
func (f Formatter) FormatDateTime(date time.Time) string {
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d.%02d", date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second(), date.Nanosecond()/1000000)
}

// FormatNow format now
func (f Formatter) FormatNow() string {
	return f.FormatDateTime(time.Now())
}

// Compose compose string with give char and count
func (f Formatter) Compose(char string, count int) string {
	builder := NewStringBuilder()
	for i := 0; i < count; i++ {
		builder.Append(char)
	}
	return builder.String()
}

// ToString conver to string
func (f Formatter) ToString(objects ...interface{}) string {
	var builder strings.Builder
	spearator := " "
	length := len(objects)

	for index, value := range objects {

		switch valueType := value.(type) {
		case fmt.Stringer:
			builder.WriteString(valueType.String())
		case string:
			builder.WriteString(valueType)
		default:
			builder.WriteString(fmt.Sprintf("%v", valueType))
		}

		if index < length-1 {
			builder.WriteString(spearator)
		}
	}
	return builder.String()
}
