package main

import (
	"bytes"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	tests := []struct {
		cmd  []string
		out  string
		code int
	}{
		{
			cmd:  []string{"sh", "-c", "echo $HELLO $BAR"},
			out:  "\"hello\" bar\n",
			code: 0,
		},
		{
			cmd:  []string{"wrong-cmd"},
			out:  "",
			code: ExitCodeCommandInvokedCannotExecute,
		},
		{
			cmd:  []string{},
			out:  "",
			code: ExitCodeCommandNotFound,
		},
	}
	for _, tc := range tests {
		// setup environment
		env, err := ReadDir("testdata/env")
		require.NoError(t, err)

		// setup stdout reader
		r, exitCode := func() (*os.File, int) {
			// setup stdout catching
			r, w, err := os.Pipe()
			defer func() {
				err := w.Close()
				if err != nil {
					slog.Error("error os.Pipe().w.Close()\n%w", slog.Any("error", err))
				}
			}()
			origStdout := os.Stdout
			defer func() { os.Stdout = origStdout }() // revert stdout
			os.Stdout = w                             // after that all stdout info will be caught and then read by r
			require.NoError(t, err)

			// run application and catch stdout
			exitCode := RunCmd(tc.cmd, env)
			return r, exitCode
		}()

		// read whole file
		var buf bytes.Buffer
		_, err = buf.ReadFrom(r)

		require.NoError(t, err)
		require.Equal(t, tc.code, exitCode)
		require.Equal(t, tc.out, buf.String())
	}
}
