package logging

// LogMessage message for logging
type LogMessage struct {
	Name     string
	Level    string
	Message  string
	Datetime string
	Indent   int
	Data     []interface{}
}

// NewMessage create a message
func NewMessage(name string, level string, indent int, message string, data ...interface{}) LogMessage {
	formatter := Formatter{}
	contextData := make([]interface{}, 0)
	for _, item := range data {
		contextData = append(contextData, item)
	}

	return LogMessage{
		Name:     name,
		Level:    level,
		Message:  message,
		Datetime: formatter.FormatNow(),
		Indent:   indent,
		Data:     contextData,
	}
}
