package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/adrianbaraka/goutils/echo"
)

type RunCmdConfig struct {
	LogLevel      echo.LogLevel
	ShouldColor   bool
	CaptureStdout bool
	StreamOutput  bool
}

// NewRunner loglevel is the level above which it won't be printed
func NewRunner(logLevel echo.LogLevel, shouldColor bool, captureStdout bool, streamoutput bool) *RunCmdConfig {
	return &RunCmdConfig{
		LogLevel:      logLevel,
		ShouldColor:   shouldColor,
		CaptureStdout: captureStdout,
		StreamOutput:  streamoutput,
	}
}

func (runner RunCmdConfig) RunCmd(loglevel echo.LogLevel, name string, args ...string) (stdoutBuffer []string, err error, exitCode int) {
	command := exec.Command(name, args...)

	// get stdout
	stdout, err := command.StdoutPipe()
	if err != nil {
		return nil, err, -1
	}

	// get stderr
	stderr, err := command.StderrPipe()
	if err != nil {
		return nil, err, -1
	}

	// run the command
	if err := command.Start(); err != nil {
		return nil, err, -1
	}

	// configure a logger as stdout/stderr maybe alot
	l := echo.NewLogger(runner.LogLevel, os.Stdout)

	l.Echof(echo.DefaultColor, echo.Trace, "Running Command: '%v %v'\n", name, strings.Join(args, " "))

	// capture stdout
	stdoutBuf := []string{}

	// capture stderr
	stderrBuf := []string{}

	// Helper to scan and log
	copyFunc := func(r io.Reader, w io.Writer, color echo.Colour, loglev echo.LogLevel) {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			if runner.StreamOutput {
				l.Fecholn(color, loglev, w, scanner.Text())
			}
			if w == os.Stdout && runner.CaptureStdout {
				stdoutBuf = append(stdoutBuf, scanner.Text())
			}
			if w == os.Stderr {
				stderrBuf = append(stderrBuf, scanner.Text())
			}
		}
	}

	// if should color
	green := echo.Green
	red := echo.Red

	if !runner.ShouldColor {
		green = echo.DefaultColor
		red = echo.DefaultColor
	}

	// stdout's log level is the passed loglevel while stderr's is echo.Error
	go copyFunc(stdout, os.Stdout, green, loglevel)
	go copyFunc(stderr, os.Stderr, red, echo.Error)

	// wait for commnd to finish
	if err := command.Wait(); err != nil {
		return stdoutBuf, fmt.Errorf("%v", strings.Join(stderrBuf, " ")), command.ProcessState.ExitCode()
	}
	return stdoutBuf, nil, command.ProcessState.ExitCode()

}

// Run a command from $PATH. If it is not found an error is returned.
// An exit code of -1 shows some other error occurred. Error should be checked.
//
//	If shouldColor is set to true the stdout is printed in green and stderr in red else the default color is used.
//
// If streamoutput is set to true stdout/stderr from the process will be streamed back.
// The exit code is returned also since in some processes an exit code of say 1 is a warning and is acceptable.
//
//	For arguements pass them as separate strings eg "ls", "-l"
func RunCmd(shouldColor bool, captureStdout bool, streamOutput bool, name string, args ...string) (stdoutBuffer []string, err error, exitCode int) {
	level := echo.Info

	if !streamOutput {
		level = echo.Error
	}
	r := NewRunner(level, shouldColor, captureStdout, streamOutput)

	return r.RunCmd(echo.Info, name, args...)
}

// Takes a list of executables that are required to be in the system path
//
//	returns an error if any of them are not present.
func RequireTools(names ...string) error {
	var notInPath []string
	for _, name := range names {
		_, err := exec.LookPath(name)
		if err != nil {
			notInPath = append(notInPath, name)
		}
	}

	if len(notInPath) != 0 {
		return fmt.Errorf("The following executables have not been found in your system path: %v", strings.Join(notInPath, ", "))
	}
	return nil
}
