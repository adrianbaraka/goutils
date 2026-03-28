// Provides functions to print colored logs.
package echo

import (
	"fmt"
	"io"
	"os"
	"sync"

	"golang.org/x/term"
)

// Global vars
const escape string = "\x1B["
const reset string = escape + "0m"

type Colour string
type LogLevel int

type Logger struct {
	level    LogLevel
	useColor bool
	out      io.Writer
	mu       sync.Mutex // ensure order is maintained if logging through go routines
}

// Returns a new logger with log level functionality.
func NewLogger(l LogLevel, out io.Writer) *Logger {
	return &Logger{
		level:    l,
		useColor: shouldColor(out),
		out:      out,
	}
}

// colors
const (
	Red          Colour = Colour(escape + "38;5;196m")
	Green        Colour = Colour(escape + "92m")
	Yellow       Colour = Colour(escape + "93m")
	Blue         Colour = Colour(escape + "38;5;27m")
	Magenta      Colour = Colour(escape + "95m")
	Cyan         Colour = Colour(escape + "96m")
	DefaultColor Colour = Colour("")
)

// Log levels
const (
	Fatal LogLevel = iota
	Error
	Warn
	Info
	Debug
	Trace
)

func shouldColor(w io.Writer) bool {
	if os.Getenv("FORCE_COLOR") != "" {
		//force-color.org
		// also for something like --color always
		return true
	}
	if os.Getenv("NO_COLOR") != "" {
		// no-color.org
		return false
	}
	// Check if the writer is a terminal
	if f, ok := w.(*os.File); ok {
		return term.IsTerminal(int(f.Fd()))
	}
	return false
}

// these adhere to the log level
// --------------------------------------------------------------------------------------------------

// Main function that handles the logic
func (l *Logger) log(color Colour, level LogLevel, w io.Writer, format string, a ...any) (int, error) {
	if level > l.level {
		return 0, nil
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	var n int
	var err error
	// TODO add time
	if l.useColor {
		fmt.Fprint(w, string(color)) // color
		n, err = fmt.Fprintf(w, format, a...)
		fmt.Fprint(w, string(reset)) // reset

		return n, err
	}
	return fmt.Fprintf(w, format, a...)
}

// Fechof formats and writes colored output to the provided writer with log level filtering. Overrides the initialized writer
func (l *Logger) Fechof(color Colour, level LogLevel, w io.Writer, format string, a ...any) (int, error) {
	return l.log(color, level, w, format, a...)
}

// Fecholn writes colored output with a newline to the provided writer with log level filtering. Overrides the initialized writer
func (l *Logger) Fecholn(color Colour, level LogLevel, w io.Writer, a ...any) (int, error) {
	msg := fmt.Sprintln(a...)
	return l.log(color, level, w, "%v", msg)
}

// Echoln writes colored output with a newline to configured writer with log level filtering.
func (l *Logger) Echoln(color Colour, level LogLevel, a ...any) (int, error) {
	msg := fmt.Sprintln(a...)
	return l.log(color, level, l.out, "%v", msg)
}

// Echof formats and writes colored output to configured writer with log level filtering.
func (l *Logger) Echof(color Colour, level LogLevel, format string, a ...any) (int, error) {
	return l.log(color, level, l.out, format, a...)
}

// some convenience methods

// Success message [Info]
func (l *Logger) Success(a ...any) {
	msg := fmt.Sprintln(a...)
	l.Echof(Green, Info, "%v", msg)
}

// Debug message [Debug]
func (l *Logger) Debug(a ...any) {
	msg := fmt.Sprintln(a...)
	l.Echof(Cyan, Debug, "%v", msg)
}

// error message [Error]
func (l *Logger) Error(a ...any) {
	msg := fmt.Sprintln(a...)
	l.Echof(Red, Error, "%v", msg)
}

// --------------------------------------------------------------------------------------------------

// Just generic no logger is needed
// main function for the generic ones
func log(color Colour, w io.Writer, format string, a ...any) (int, error) {
	l := NewLogger(Info, w)
	return l.log(color, Info, w, format, a...)
}

// Fechof formats and writes colored output to the provided writer.
func Fechof(color Colour, w io.Writer, format string, a ...any) (int, error) {
	return log(color, w, format, a...)
}

// Fecholn writes colored output with a newline to the provided writer.
func Fecholn(color Colour, w io.Writer, a ...any) (int, error) {
	msg := fmt.Sprintln(a...)
	return log(color, w, "%v", msg)
}

// Echoln writes colored output with a newline to stdout.
func Echoln(color Colour, a ...any) (int, error) {
	msg := fmt.Sprintln(a...)
	return log(color, os.Stdout, "%v", msg)
}

// Echof formats and writes colored output to stdout.
func Echof(color Colour, format string, a ...any) (int, error) {
	return log(color, os.Stdout, format, a...)
}
