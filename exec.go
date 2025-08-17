package exec

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	yup "github.com/yupsh/framework"
	"github.com/yupsh/framework/opt"
	localopt "github.com/yupsh/exec/opt"
)

// Flags represents the configuration options for the exec command
type Flags = localopt.Flags

// Command implementation
type command opt.Inputs[string, Flags]

// Exec creates a new exec command with the given parameters
// First parameter is the command, rest are arguments
func Exec(parameters ...any) yup.Command {
	cmd := command(opt.Args[string, Flags](parameters...))
	// Set defaults
	if cmd.Flags.InheritEnv == localopt.InheritEnvFlag(false) {
		cmd.Flags.InheritEnv = localopt.InheritEnv // Default to inheriting environment
	}
	if cmd.Flags.Shell == "" && cmd.Flags.UseShell {
		cmd.Flags.Shell = "/bin/sh" // Default shell
	}
	return cmd
}

func (c command) Execute(ctx context.Context, stdin io.Reader, stdout, stderr io.Writer) error {
	if len(c.Positional) == 0 {
		return fmt.Errorf("exec: no command specified")
	}

	var cmd *exec.Cmd

	if bool(c.Flags.UseShell) {
		// Run through shell
		shell := string(c.Flags.Shell)
		if shell == "" {
			shell = "/bin/sh"
		}

		// Join all arguments into a single command string
		cmdStr := strings.Join(c.Positional, " ")
		cmd = exec.CommandContext(ctx, shell, "-c", cmdStr)
	} else {
		// Run directly
		cmdName := c.Positional[0]
		args := c.Positional[1:]
		cmd = exec.CommandContext(ctx, cmdName, args...)
	}

	// Set working directory
	if c.Flags.WorkingDir != "" {
		cmd.Dir = string(c.Flags.WorkingDir)
	}

	// Set environment
	if bool(c.Flags.InheritEnv) {
		cmd.Env = os.Environ()
	}

	// Add custom environment variables
	for _, envVar := range c.Flags.EnvVars {
		if cmd.Env == nil {
			cmd.Env = []string{}
		}
		cmd.Env = append(cmd.Env, string(envVar))
	}

	// Set up IO
	cmd.Stdin = stdin
	cmd.Stdout = stdout

	if bool(c.Flags.Quiet) {
		cmd.Stderr = nil // Discard stderr
	} else {
		cmd.Stderr = stderr
	}

	// Handle interactive mode
	if bool(c.Flags.Interactive) {
		// For interactive commands, we might need to connect to the terminal
		// For interactive commands, connect to terminal if available
		cmd.Stdin = stdin
		cmd.Stdout = stdout
		cmd.Stderr = stderr
	}

	// Execute the command
	err := cmd.Run()

	// Handle errors based on flags
	if err != nil && bool(c.Flags.IgnoreErrors) {
		// Ignore the error and continue
		return nil
	}

	return err
}

func (c command) String() string {
	return fmt.Sprintf("exec %v", c.Positional)
}

// Wrapper creates a predefined exec command with preset arguments
// This allows creating reusable command patterns
func Wrapper(name string, presetArgs []string, parameters ...any) yup.Command {
	// Combine preset arguments with provided parameters
	allParams := make([]any, 0, len(presetArgs)+len(parameters))

	// Add preset arguments
	for _, arg := range presetArgs {
		allParams = append(allParams, arg)
	}

	// Add provided parameters
	allParams = append(allParams, parameters...)

	return Exec(allParams...)
}

// Helper function for creating command wrappers
// This is mainly useful for other packages that want to build on exec
