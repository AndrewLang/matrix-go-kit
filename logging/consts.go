package logging

// ColorInfo descrble console color
type ColorInfo struct {
	Name  string
	Value string
}

const (
	/*====================================== Common =============================*/

	// Empty empty string
	Empty string = ""
	// NewLine line terminator
	NewLine string = "\n"
	// OneMega 1M
	OneMega int64 = 1024 * 1024
	// DefaultLogFileSize defautl log file size
	DefaultLogFileSize int64 = OneMega * 2
	// DefaultLayoutPadding padding
	DefaultLayoutPadding int = 8

	// LeftBracket left bracket
	LeftBracket = "["
	// RightBracket right bracket
	RightBracket = "]"
	// Space space
	Space = " "

	// StyleSeparator sign uses to separator styles
	StyleSeparator = ","

	/*====================================== Logger names =============================*/

	// ConsoleLoggerName console logger
	ConsoleLoggerName = "Console"
	// FileLoggerName file logger
	FileLoggerName = "File"
	// JSONFileLoggerName json file logger
	JSONFileLoggerName = "JsonFile"
	// ComposeLoggerName name compose logger
	ComposeLoggerName = "Compose"

	/*====================================== Level names =============================*/

	// Name name
	Name = "Name"
	// Time time
	Time = "Time"
	// Level level
	Level = "Level"
	// Indent indent
	Indent = "Indent"
	// Message message
	Message = "Message"
	// Compose compose
	Compose = "Compose"

	// NameLayout name layout
	NameLayout = "NameLayout"
	// TimeLayout time layout
	TimeLayout = "TimeLayout"
	// LevelLayout level
	LevelLayout = "LevelLayout"
	// IndentLayout indent
	IndentLayout = "IndentLayout"
	// MessageLayout message
	MessageLayout = "MessageLayout"
	// ComposeLayout compose
	ComposeLayout = "ComposeLayout"

	/*====================================== Colors =============================*/

	// StyleEsc start color
	StyleEsc = "\033["
	// StyleEnd style end indicator, a style could be 37m, or combined style 1;2;32m
	StyleEnd = "m"

	// Color256Start 256 color start
	Color256Start = "38;5;"
	// Color256End 256 color end
	Color256End = ""
	// Color256BgStart 256 color start
	Color256BgStart = "48;5;"
	// Color256BgEnd 256 color end
	Color256BgEnd = ""

	/*====================================== Styles =============================*/

	// StyleBold bold style
	StyleBold = "1"
	// StyleDim dim
	StyleDim = "2"
	// StyleUnderline underline
	StyleUnderline = "4"
	// StyleBlink blink
	StyleBlink = "5"
	// StyleReverse reverse
	StyleReverse = "7"
	// StyleHidden hidden
	StyleHidden = "8"

	// ResetAllStyle clear color
	ResetAllStyle = StyleEsc + "0m"
	// ResetBold reset bold
	ResetBold = StyleEsc + "21m"
	// ResetDim reset dim
	ResetDim = StyleEsc + "22m"
	// ResetUnderline reset underline
	ResetUnderline = StyleEsc + "24m"
	// ResetBlink reset blink
	ResetBlink = StyleEsc + "25m"
	// ResetReverse reset reverse
	ResetReverse = StyleEsc + "27m"
	// ResetHidden reset hidden
	ResetHidden = StyleEsc + "28m"

	// ForegroundPrefix foreground fix
	ForegroundPrefix = "Fore"
	// BackgroundPrefix background prefix
	BackgroundPrefix = "Bg"
)

var (
	/*====================================== Color names =============================*/

	// ColorDefaultText default text color
	ColorDefaultText = ColorInfo{"DefaultText", "39"}
	// ColorBlack black
	ColorBlack = ColorInfo{"Black", "30"}
	// ColorRed red
	ColorRed = ColorInfo{"Red", "31"}
	// ColorGreen green
	ColorGreen = ColorInfo{"Green", "32"}
	// ColorYellow yellow
	ColorYellow = ColorInfo{"Yellow", "33"}
	// ColorBlue blue
	ColorBlue = ColorInfo{"Blue", "34"}
	// ColorMagenta magenta
	ColorMagenta = ColorInfo{"Magenta", "35"}
	// ColorCyan cyan
	ColorCyan = ColorInfo{"Cyan", "36"}
	// ColorLightGray light gray
	ColorLightGray = ColorInfo{"LightGray", "37"}
	// ColorBgDefault default bg
	ColorBgDefault = ColorInfo{"Default", "49"}
	// ColorBgBlack black
	ColorBgBlack = ColorInfo{"Black", "40"}
	// ColorBgRed red
	ColorBgRed = ColorInfo{"Red", "41"}
	// ColorBgGreen green
	ColorBgGreen = ColorInfo{"Green", "42"}
	// ColorBgYellow yellow
	ColorBgYellow = ColorInfo{"Yellow", "43"}
	// ColorBgBlue blue
	ColorBgBlue = ColorInfo{"Blue", "44"}
	// ColorBgMagenta magenta
	ColorBgMagenta = ColorInfo{"Magenta", "45"}
	// ColorBgCyan cyan
	ColorBgCyan = ColorInfo{"Cyan", "46"}
	// ColorBgLightGray light gray
	ColorBgLightGray = ColorInfo{"LightGray", "47"}
	// ColorBgDarkGray dark gray
	ColorBgDarkGray = ColorInfo{"DarkGray", "100"}
	// ColorBgLightRed light red
	ColorBgLightRed = ColorInfo{"LightRed", "101"}
	// ColorBgLightGreen light green
	ColorBgLightGreen = ColorInfo{"LightGreen", "102"}
	// ColorBgLightYellow light yellow
	ColorBgLightYellow = ColorInfo{"LightYellow", "103"}
	// ColorBgLightBlue light blue
	ColorBgLightBlue = ColorInfo{"LightBlue", "104"}
	// ColorBgLightMagenta light magenta
	ColorBgLightMagenta = ColorInfo{"LightMagenta", "105"}
	// ColorBgLightCyan light cyan
	ColorBgLightCyan = ColorInfo{"LightCyan", "106"}
	// ColorBgWhite bg white
	ColorBgWhite = ColorInfo{"White", "107"}
	// ColorDarkGray dark gray
	ColorDarkGray = ColorInfo{"DarkGray", "90"}
	// ColorLightRed light red
	ColorLightRed = ColorInfo{"LightRed", "91"}
	// ColorLightGreen light green
	ColorLightGreen = ColorInfo{"LightGreen", "92"}
	// ColorLightYellow light yellow
	ColorLightYellow = ColorInfo{"LightYellow", "93"}
	// ColorLightBlue lightblue
	ColorLightBlue = ColorInfo{"LightBlue", "94"}
	// ColorLightMagenta light magenta
	ColorLightMagenta = ColorInfo{"LightMagenta", "95"}
	// ColorLightCyan light cyan
	ColorLightCyan = ColorInfo{"LightCyan", "96"}
	// ColorWhite white
	ColorWhite = ColorInfo{"White", "97"}
	// ColorBrightBlack bright black
	ColorBrightBlack = ColorInfo{"BrightBlack", "30;1"}
	// ColorBrightRed bright red
	ColorBrightRed = ColorInfo{"BrightRed", "31;1"}
	// ColorBrightGreen bright green
	ColorBrightGreen = ColorInfo{"BrightGreen", "32;1"}
	// ColorBrightYellow brightyellow
	ColorBrightYellow = ColorInfo{"BrightYellow", "33;1"}
	// ColorBrightBlue bright blue
	ColorBrightBlue = ColorInfo{"BrightBlue", "34;1"}
	// ColorBrightMagenta bright magenta
	ColorBrightMagenta = ColorInfo{"BrightMagenta", "35;1"}
	// ColorBrightCyan bright cyan
	ColorBrightCyan = ColorInfo{"BrightCyan", "36;1"}
	// ColorBrightWhite bright white
	ColorBrightWhite = ColorInfo{"BrightWhite", "37;1"}

	// KnownStyles know styles
	KnownStyles = map[string]string{
		"Bold":      "1",
		"Dim":       "2",
		"Underline": "4",
		"Blink":     "5",
		"Reverse":   "6",
		"Hidden":    "8",
	}
	// KnownColors known colors
	KnownColors = map[string]string{
		"ForeDefaultText":  "39",
		"ForeBlack":        "30",
		"ForeRed":          "31",
		"ForeGreen":        "32",
		"ForeYellow":       "33",
		"ForeBlue":         "34",
		"ForeMagenta":      "35",
		"ForeCyan":         "36",
		"ForeLightGray":    "37",
		"ForeDarkGray":     "90",
		"ForeLightRed":     "91",
		"ForeLightGreen":   "92",
		"ForeLightYellow":  "93",
		"ForeLightBlue":    "94",
		"ForeLightMagenta": "95",
		"ForeLightCyan":    "96",
		"ForeLightWhite":   "97",
		"BgDefault":        "49",
		"BgBlack":          "40",
		"BgRed":            "41",
		"BgGreen":          "42",
		"BgYellow":         "43",
		"BgBlue":           "44",
		"BgMagenta":        "45",
		"BgCyan":           "46",
		"BgLightGray":      "47",
		"BgDarkGray":       "100",
		"BgLightRed":       "101",
		"BgLightGreen":     "102",
		"BgLightYellow":    "103",
		"BgLightBlue":      "104",
		"BgLightMagenta":   "105",
		"BgLightCyan":      "106",
		"BgLightWhite":     "107",
		"BrightBlack":      "30;1",
		"BrightRed":        "31;1",
		"BrightGreen":      "32;1",
		"BrightYellow":     "33;1",
		"BrightBlue":       "34;1",
		"BrightMagenta":    "35;1",
		"BrightCyan":       "36;1",
		"BrightWhite":      "37;1",
	}
)
