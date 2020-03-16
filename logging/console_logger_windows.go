package log

import (
	"os"

	"golang.org/x/sys/windows"
)

func (logger *ConsoleLogger) initializeConsole() {
	/*
	  set console mode to enable virtual terminal processing,
	  otherwise it may not work on some windows
	*/
	stdout := windows.Handle(os.Stdout.Fd())
	var originalMode uint32

	windows.GetConsoleMode(stdout, &originalMode)
	windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}
