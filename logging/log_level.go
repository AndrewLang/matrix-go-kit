package logging

var (
	// LevelAll all
	LevelAll LogLevel = LogLevel{"All", 0}

	// LevelDebug debug
	LevelDebug LogLevel = LogLevel{"Debug", 1}

	// LevelInfo inf
	LevelInfo = LogLevel{"Info", 2}

	// LevelWarn warn
	LevelWarn = LogLevel{"Warn", 3}

	// LevelError error
	LevelError = LogLevel{"Error", 4}

	// LevelFatal critical
	LevelFatal = LogLevel{"Fatal", 5}

	// LevelNone none
	LevelNone = LogLevel{"None", int(0 >> 1)}
)

// LogLevel represent level of log
type LogLevel struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
