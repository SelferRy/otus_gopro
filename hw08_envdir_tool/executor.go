package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
)

const (
	ExitCodeCommandInvokedCannotExecute = 126
	ExitCodeCommandNotFound             = 127
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	if len(cmd) < 1 {
		return ExitCodeCommandNotFound
	}
	command, args := cmd[0], cmd[1:]
	execCmd := exec.Command(command, args...)
	execCmd.Stdin, execCmd.Stdout, execCmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	err := setupEnv(env)
	if err != nil {
		slog.Error("problem with setupEnv", slog.Any("error", err))
	}
	if err = execCmd.Run(); err != nil {
		slog.Error("execution error", slog.Any("error:", err))
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return exitErr.ExitCode()
		}
		return ExitCodeCommandInvokedCannotExecute
	}
	return execCmd.ProcessState.ExitCode()
}

// change environment variable.
func setupEnv(env Environment) error {
	for key, val := range env {
		if val.NeedRemove {
			err := os.Unsetenv(key)
			if err != nil {
				return fmt.Errorf("os.Unsetenv problem.\n%w", err)
			}
		}
		err := os.Setenv(key, val.Value)
		if err != nil {
			return fmt.Errorf("os.Setenv problem.\n%w", err)
		}
	}
	return nil
}
