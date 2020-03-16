package app

import (
	"testing"

	io "github.com/andrewlang/matrix-go-kit/io"
	"github.com/stretchr/testify/assert"
)

func TestCreateLogger(t *testing.T) {
	logger := CreateLogger("test")

	assert.NotNil(t, logger)
}

func TestCreateLoggerWithFile(t *testing.T) {

	file := io.NewFile(LoggingConfigFile)
	file.Write(`
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
	}`)
	defer file.Delete()

	logger := CreateLogger("test")

	assert.NotNil(t, logger)
}
