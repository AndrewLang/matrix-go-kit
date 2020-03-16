package app

import (
	"os"
	"path"
	"path/filepath"
)

// ApplicationContext context for application
type ApplicationContext struct {
	Configuration *Configuration
}

// NewApplicationContext create a new context instance
func NewApplicationContext() *ApplicationContext {
	workDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	configFile := path.Join(workDir, DefaultConfigFile)
	configuration := NewConfiguration().FromFile(configFile)

	context := &ApplicationContext{
		Configuration: configuration,
	}

	return context
}
