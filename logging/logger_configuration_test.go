package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func defaultLogginConfig() *LoggerConfiguration {
	configuration := NewLoggerConfiguration([]string{Time, Level, Name, Indent, Message})
	configuration.DebugStyle = LogLevelStyle{"245", "24", ""}
	configuration.InfoStyle = LogLevelStyle{"85", "", "1"}
	configuration.WarnStyle = LogLevelStyle{"226", "124", "4"}
	configuration.ErrorStyle = LogLevelStyle{"166", "232", "1,4"}
	configuration.FatalStyle = LogLevelStyle{"196", "11", "7"}

	return configuration
}

func defaultLogTargetConfig() *LogTargetConfiguration {
	configuration := NewLogTargetConfiguration("Test", ConsoleLoggerName, []string{Time, Level, Name, Indent, Message})
	return configuration
}

func defaultLogTargetConfigs() *LogTargetConfigurations {
	config := NewLogTargetConfigurations()
	config.AddTarget(defaultLogTargetConfig())

	return config
}

func composeLogTargetConfigs() *LogTargetConfigurations {
	config := NewLogTargetConfigurations()
	config.AddTarget(defaultLogTargetConfig())

	fileConfig := NewLogTargetConfiguration("File", FileLoggerName, []string{Time, Level, Name, Indent, Message})
	fileConfig.Configuration.FileName = "./test/compose_logger.txt"

	config.AddTarget(fileConfig)

	jsonConfig := NewLogTargetConfiguration("JsonFile", JSONFileLoggerName, []string{Message})
	jsonConfig.Configuration.FileName = "./test/compose_logger.json"
	jsonConfig.Configuration.MinLevel = LevelFatal.Value
	config.AddTarget(jsonConfig)

	return config
}
func verifyLoggerConfiguration(t *testing.T, configuration *LoggerConfiguration) {
	assert.Equal(t, "", configuration.FileName, "")
	assert.Equal(t, 0, configuration.MinLevel, "")
	assert.Equal(t, int64(2097152), configuration.FileSize, "")
	assert.Equal(t, true, configuration.UseColor, "")
	assert.Equal(t, "Time", configuration.LayoutNames[0], "")
	assert.Equal(t, "Level", configuration.LayoutNames[1], "")
	assert.Equal(t, "Name", configuration.LayoutNames[2], "")
	assert.Equal(t, "Indent", configuration.LayoutNames[3], "")
	assert.Equal(t, "Message", configuration.LayoutNames[4], "")

	assert.Equal(t, "245", configuration.DebugStyle.Foreground, "")
	assert.Equal(t, "24", configuration.DebugStyle.Background, "")
	assert.Equal(t, "", configuration.DebugStyle.Styles, "")

	assert.Equal(t, "56", configuration.InfoStyle.Foreground, "")
	assert.Equal(t, "234", configuration.InfoStyle.Background, "")
	assert.Equal(t, "1", configuration.InfoStyle.Styles, "")

	assert.Equal(t, "226", configuration.WarnStyle.Foreground, "")
	assert.Equal(t, "124", configuration.WarnStyle.Background, "")
	assert.Equal(t, "4", configuration.WarnStyle.Styles, "")

	assert.Equal(t, "166", configuration.ErrorStyle.Foreground, "")
	assert.Equal(t, "232", configuration.ErrorStyle.Background, "")
	assert.Equal(t, "1,4", configuration.ErrorStyle.Styles, "")

	assert.Equal(t, "196", configuration.FatalStyle.Foreground, "")
	assert.Equal(t, "11", configuration.FatalStyle.Background, "")
	assert.Equal(t, "7", configuration.FatalStyle.Styles, "")
}

func TestSerializeJsonConfiguration(t *testing.T) {
	configuration := defaultLogginConfig()

	json := configuration.ToJSON()

	logger := NewConsoleLogger("Test_Env").Configure(configuration)
	logger.Info("Json logger configuration").
		Info(json)
}

func TestDeserializeJsonConfiguration(t *testing.T) {
	content := `
	{
			"LayoutNames": [
					"Time",
					"Level",
					"Name",
					"Indent",
					"Message"
			],
			"fileName": "",
			"fileSize": 2097152,
			"minLevel": 0,
			"useColor": true,
			"debugStyle": {
					"foreground": "245",
					"background": "24",
					"styles": ""
			},
			"infoStyle": {
					"foreground": "56",
					"background": "234",
					"styles": "1"
			},
			"warnStyle": {
					"foreground": "226",
					"background": "124",
					"styles": "4"
			},
			"errorStyle": {
					"foreground": "166",
					"background": "232",
					"styles": "1,4"
			},
			"fatalStyle": {
					"foreground": "196",
					"background": "11",
					"styles": "7"
			}
	}`

	configuration := NewLoggerConfiguration([]string{Time, Level, Name, Indent, Message})
	configuration.FromJSON(content)

	verifyLoggerConfiguration(t, configuration)
}

func TestNewLogTargetConfiguration(t *testing.T) {
	config := defaultLogTargetConfig()

	assert.Equal(t, "Test", config.Name)
	assert.NotNil(t, config.Configuration)
}

func TestToJsonLogTargetConfiguration(t *testing.T) {
	config := defaultLogTargetConfig()

	json := config.ToJSON()

	logger := NewConsoleLogger("Test_Env").Configure(config.Configuration)
	logger.Info("Log target configuration").
		Info(json)
}

func TestFromJsonLogTargetConfiguration(t *testing.T) {
	content := `{
	"name": "Test",
	"config": {
			"LayoutNames": [
					"Time",
					"Level",
					"Name",
					"Indent",
					"Message"
			],
			"fileName": "",
			"fileSize": 2097152,
			"minLevel": 0,
			"useColor": true,
			"debugStyle": {
				"foreground": "245",
				"background": "24",
				"styles": ""
		},
		"infoStyle": {
				"foreground": "56",
				"background": "234",
				"styles": "1"
		},
		"warnStyle": {
				"foreground": "226",
				"background": "124",
				"styles": "4"
		},
		"errorStyle": {
				"foreground": "166",
				"background": "232",
				"styles": "1,4"
		},
		"fatalStyle": {
				"foreground": "196",
				"background": "11",
				"styles": "7"
		}
	}
}`

	configuration := &LogTargetConfiguration{}
	configuration.FromJSON(content)

	verifyLoggerConfiguration(t, configuration.Configuration)

	logger := NewConsoleLogger("Test_Env").Configure(configuration.Configuration)
	logger.Info("Log target configuration").
		Info("", configuration.Name, configuration.Configuration)
}

func TestNewLogTargetConfigurations(t *testing.T) {
	config := NewLogTargetConfigurations()

	assert.NotNil(t, config)
	assert.Equal(t, 0, len(config.Targets))
}

func TestLogTargetConfigurationsAddTarget(t *testing.T) {
	config := defaultLogTargetConfigs()
	config.AddTarget(NewLogTargetConfiguration("Test_Json", JSONFileLoggerName, []string{Time, Level, Name, Message}))

	json := config.ToJSON()

	logger := NewConsoleLogger("Test_Env").Configure(config.Targets[0].Configuration)
	logger.Info("Log target configuration").
		Info(json)
}

func TestLogTargetConfigurationsToJSON(t *testing.T) {
	config := defaultLogTargetConfigs()
	json := config.ToJSON()

	logger := NewConsoleLogger("Test_Env").Configure(config.Targets[0].Configuration)
	logger.Info("Log target configuration").
		Info(json)
}

func TestLogTargetConfigurationsFromJSON(t *testing.T) {
	content := ` {
	"targets": [
			{
					"name": "Test",
					"type": "Console",
					"config": {
							"LayoutNames": [
									"Time",
									"Level",
									"Name",
									"Indent",
									"Message"
							],
							"fileName": "",
							"fileSize": 2097152,
							"minLevel": 0,
							"useColor": true,
							"debugStyle": {
								"foreground": "245",
								"background": "24",
								"styles": ""
						},
						"infoStyle": {
								"foreground": "56",
								"background": "234",
								"styles": "1"
						},
						"warnStyle": {
								"foreground": "226",
								"background": "124",
								"styles": "4"
						},
						"errorStyle": {
								"foreground": "166",
								"background": "232",
								"styles": "1,4"
						},
						"fatalStyle": {
								"foreground": "196",
								"background": "11",
								"styles": "7"
						}
					}
			},
			{
					"name": "Test_Json",
					"type": "JsonFile",
					"config": {
							"LayoutNames": [
									"Time",
									"Level",
									"Name",
									"Message"
							],
							"fileName": "",
							"fileSize": 2097152,
							"minLevel": 0,
							"useColor": true,
							"debugStyle": {
								"foreground": "245",
								"background": "24",
								"styles": ""
						},
						"infoStyle": {
								"foreground": "56",
								"background": "234",
								"styles": "1"
						},
						"warnStyle": {
								"foreground": "226",
								"background": "124",
								"styles": "4"
						},
						"errorStyle": {
								"foreground": "166",
								"background": "232",
								"styles": "1,4"
						},
						"fatalStyle": {
								"foreground": "196",
								"background": "11",
								"styles": "7"
						}
					}
			}
	]
}`

	config := NewLogTargetConfigurations()
	config.FromJSON(content)
	assert.Equal(t, 2, len(config.Targets))

	target1 := config.GetTarget("Test")
	target2 := config.GetTarget("Test_Json")

	assert.NotNil(t, target1)
	assert.NotNil(t, target2)

	assert.Equal(t, "Console", target1.Type)
	assert.Equal(t, "JsonFile", target2.Type)

	verifyLoggerConfiguration(t, target1.Configuration)
}
