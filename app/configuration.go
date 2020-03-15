package app

import (
	"encoding/json"

	"github.com/andrewlang/matrix-go-kit/io"
)

// Configuration configuration model for test station agent
type Configuration struct {
}

// NewConfiguration create new configuration
func NewConfiguration() *Configuration {
	config := &Configuration{}
	return config
}

// ToJSON serialize to json string
func (config *Configuration) ToJSON() string {
	content, _ := json.MarshalIndent(config, "", "\t")
	return string(content)
}

// String convert to string
func (config *Configuration) String() string {
	return config.ToJSON()
}

// FromJSON load configuration from json string
func (config *Configuration) FromJSON(jsonContent string) *Configuration {
	json.Unmarshal([]byte(jsonContent), config)
	return config
}

// ToFile save configuration to file
func (config *Configuration) ToFile(filePath string) *Configuration {
	content := config.ToJSON()
	io.NewFile(filePath).Write(content)
	return config
}

// FromFile load configuration from file
func (config *Configuration) FromFile(filePath string) *Configuration {
	file := io.NewFile(filePath)

	if file.Exists() {
		content := file.ReadAll()
		config.FromJSON(content)
	}
	return config
}
