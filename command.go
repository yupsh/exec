package command

import (
	"context"
	"fmt"
	"io"
	"os/exec"

	yup "github.com/gloo-foo/framework"
)

type command yup.Inputs[string, flags]

func Exec(parameters ...any) yup.Command {
	return command(yup.Initialize[string, flags](parameters...))
}

func (p command) Executor() yup.CommandExecutor {
	return func(ctx context.Context, stdin io.Reader, stdout, stderr io.Writer) error {
		// Get command and arguments from positional parameters
		if len(p.Positional) == 0 {
			_, _ = fmt.Fprintf(stderr, "exec: no command specified\n")
			return fmt.Errorf("exec requires a command to execute")
		}

		cmdName := p.Positional[0]
		cmdArgs := p.Positional[1:]

		// Execute the command
		cmd := exec.CommandContext(ctx, cmdName, cmdArgs...)
		cmd.Stdin = stdin
		cmd.Stdout = stdout
		cmd.Stderr = stderr

		return cmd.Run()
	}
}
