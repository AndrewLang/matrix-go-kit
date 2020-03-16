package log

import (
	"strings"
)

// LogLevelStyle level style configuration with foreground, background color and styles
type LogLevelStyle struct {
	Foreground string `json:"foreground"`
	Background string `json:"background"`
	Styles     string `json:"styles"`
}

func (style LogLevelStyle) parseForeground() string {
	color := ""

	if IsNullOrEmpty(style.Foreground) {
		return color
	}

	if IsNumber(style.Foreground) {
		formatter := NewFormatter()
		number, _ := ToNumber(style.Foreground)
		color = formatter.Format256Color(number)
	} else {
		if value, ok := KnownColors[ForegroundPrefix+style.Foreground]; ok {
			color = value
		}
	}
	return color
}

func (style LogLevelStyle) parseBackground() string {
	color := ""

	if IsNullOrEmpty(style.Background) {
		return color
	}

	if IsNumber(style.Background) {
		formatter := NewFormatter()
		number, _ := ToNumber(style.Background)
		color = formatter.FormatBg256Color(number)
	} else {
		if value, ok := KnownColors[BackgroundPrefix+style.Background]; ok {
			color = value
		}
	}

	return color
}

func (style LogLevelStyle) parseStyle() []string {
	styles := make([]string, 0)

	if IsNullOrEmpty(style.Background) {
		return styles
	}

	for _, part := range strings.Split(style.Styles, StyleSeparator) {
		value := strings.Trim(part, Space)
		if IsNullOrEmpty(value) {
			continue
		}

		if s, ok := KnownStyles[value]; ok {
			styles = append(styles, s)
		} else {
			styles = append(styles, value)
		}
	}

	return styles
}

func (style LogLevelStyle) parseLevelStyles() []string {
	styles := make([]string, 0)
	part := style.parseForeground()

	if NotNullOrEmpty(part) {
		styles = append(styles, part)
	}

	part = style.parseBackground()
	if NotNullOrEmpty(part) {
		styles = append(styles, part)
	}

	parts := style.parseStyle()
	if len(parts) > 0 {
		styles = append(styles, parts...)
	}

	return styles
}
