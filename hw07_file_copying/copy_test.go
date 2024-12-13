package main

import (
	"errors"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("copy /dev/urandom", func(t *testing.T) {
		err := Copy("/dev/urandom", "/tmp", 0, 0)
		require.Equal(t, ErrUnsupportedFile, errors.Unwrap(err))
	})
	t.Run("file doesn't exist", func(t *testing.T) {
		err := Copy("test.txt", "/tmp", 0, 0)
		require.NotNil(t, err)
	})
	t.Run("copy directory", func(t *testing.T) {
		err := Copy("/tmp", "test.txt", 0, 0)
		require.Equal(t, ErrUnsupportedFile, errors.Unwrap(err))
	})
	t.Run("copy to directory", func(t *testing.T) {
		_, _ = os.Create("test.txt")
		defer func() {
			if err := os.Remove("test.txt"); err != nil {
				slog.Any("error", err)
			}
		}()
		err := Copy("test.txt", "/tmp", 0, 0)
		require.Equal(t, ErrEmptyFile, errors.Unwrap(errors.Unwrap(err)))
	})
	t.Run("permissions", func(t *testing.T) {
		_, _ = os.Create("test1.txt")
		defer func() {
			if err := os.Remove("test1.txt"); err != nil {
				slog.Any("error", err)
			}
		}()

		_, _ = os.Create("test2.txt")
		defer func() {
			if err := os.Remove("test2.txt"); err != nil {
				slog.Any("error", err)
			}
		}()

		_ = os.Chmod("test2.txt", 0444)

		err := Copy("test1.txt", "test2.txt", 0, 0)
		require.NotNil(t, err)
	})
	t.Run("copy file to itself", func(t *testing.T) {
		f, _ := os.Create("test.txt")
		defer func() {
			if err := os.Remove("test.txt"); err != nil {
				slog.Any("error", err)
			}
		}()
		_, _ = f.WriteString("test")

		err := Copy("test.txt", "test.txt", 0, 0)
		require.NotNil(t, err)
	})
}
