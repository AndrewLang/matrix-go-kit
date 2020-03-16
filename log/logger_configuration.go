package log

import (
	"encoding/json"
)

// LoggerConfiguration configuration for logger
type LoggerConfiguration struct {
	LayoutNames []string      `json:"layoutNames`
	FileName    string        `json:"fileName"`
	FileSize    int64         `json:"fileSize"`
	MinLevel    int           `json:"minLevel"`
	UseColor    bool          `json:"useColor"`
	DebugStyle  LogLevelStyle `json:"debugStyle"`
	InfoStyle   LogLevelStyle `json:"infoStyle"`
	WarnStyle   LogLevelStyle `json:"warnStyle"`
	ErrorStyle  LogLevelStyle `json:"errorStyle"`
	FatalStyle  LogLevelStyle `json:"fatalStyle"`
}

// NewLoggerConfiguration create new configuration
func NewLoggerConfiguration(layouts []string) *LoggerConfiguration {
	return &LoggerConfiguration{
		LayoutNames: layouts,
		FileName:    Empty,
		FileSize:    DefaultLogFileSize,
		MinLevel:    LevelAll.Value,
		UseColor:    true,
		DebugStyle:  LogLevelStyle{ColorDefaultText.Name, "", ""},
		InfoStyle:   LogLevelStyle{ColorGreen.Name, "", ""},
		WarnStyle:   LogLevelStyle{ColorYellow.Name, "", ""},
		ErrorStyle:  LogLevelStyle{ColorMagenta.Name, "", ""},
		FatalStyle:  LogLevelStyle{ColorRed.Name, "", ""},
	}
}

// ToJSON Convert configuration to json string
func (config *LoggerConfiguration) ToJSON() string {
	content, _ := json.MarshalIndent(config, "", "\t")

	return string(content)
}

// FromJSON Load configuration from json string
func (config *LoggerConfiguration) FromJSON(jsonContent string) {
	json.Unmarshal([]byte(jsonContent), config)
}

/*=======================================================================*/

// LogTargetConfiguration configure log output
type LogTargetConfiguration struct {
	Name          string               `json:"name"`
	Type          string               `json:"type"`
	Configuration *LoggerConfiguration `json:"config"`
}

// NewLogTargetConfiguration create a new log target configuration
func NewLogTargetConfiguration(name string, loggerType string, layouts []string) *LogTargetConfiguration {
	return &LogTargetConfiguration{
		Name:          name,
		Type:          loggerType,
		Configuration: NewLoggerConfiguration(layouts),
	}
}

// ToJSON serialize to json string
func (config *LogTargetConfiguration) ToJSON() string {
	content, _ := json.MarshalIndent(config, "", "\t")

	return string(content)
}

// FromJSON deserialize from json string
func (config *LogTargetConfiguration) FromJSON(jsonContent string) {
	json.Unmarshal([]byte(jsonContent), config)
}

/*=======================================================================*/

// LogTargetConfigurations configuration for log targets
type LogTargetConfigurations struct {
	Targets []*LogTargetConfiguration `json:"targets"`
}

// NewLogTargetConfigurations new configuration
func NewLogTargetConfigurations() *LogTargetConfigurations {
	return &LogTargetConfigurations{
		Targets: make([]*LogTargetConfiguration, 0),
	}
}

// AddTarget add target configuration
func (config *LogTargetConfigurations) AddTarget(target *LogTargetConfiguration) *LogTargetConfigurations {
	config.Targets = append(config.Targets, target)
	return config
}

// GetTarget get target by name
func (config *LogTargetConfigurations) GetTarget(name string) *LogTargetConfiguration {
	for _, item := range config.Targets {
		if item.Name == name {
			return item
		}
	}
	return nil
}

// ToJSON serialize to json string
func (config *LogTargetConfigurations) ToJSON() string {
	content, _ := json.MarshalIndent(config, "", "\t")

	return string(content)
}

// FromJSON deserialize from json string
func (config *LogTargetConfigurations) FromJSON(jsonContent string) {
	json.Unmarshal([]byte(jsonContent), config)
}
