package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReadDir(t *testing.T) {
	// Place your code here
	t.Run("read directory", func(t *testing.T) {
		target := Environment{
			"BAR":   EnvValue{"bar", false},
			"EMPTY": EnvValue{"", false},
			"FOO":   EnvValue{"   foo", false},
			"HELLO": EnvValue{"foo", false},
			"UNSET": EnvValue{"", true},
		}
		result, err := ReadDir("./testdata/env/")
		require.NoError(t, err)
		require.Equal(t, target, result)
	})
}

func TestDefineEnvVal(t *testing.T) {
	// Place your code here
	t.Run("BAR", func(t *testing.T) {
		target := EnvValue{"bar", false}
		result, err := defineEnvVal("./testdata/env/BAR")
		require.NoError(t, err)
		require.Equal(t, target, result)
	})
	t.Run("EMPTY", func(t *testing.T) {
		target := EnvValue{"", false}
		result, err := defineEnvVal("./testdata/env/EMPTY")
		require.NoError(t, err)
		require.Equal(t, target, result)
	})
	t.Run("FOO", func(t *testing.T) {
		target := EnvValue{"   foo\nwith new line", false}
		result, err := defineEnvVal("./testdata/env/FOO")
		require.NoError(t, err)
		require.Equal(t, target, result)
	})
	t.Run("HELLO", func(t *testing.T) {
		target := EnvValue{`"hello"`, false}
		result, err := defineEnvVal("./testdata/env/HELLO")
		require.NoError(t, err)
		require.Equal(t, target, result)
	})
	t.Run("TRIM", func(t *testing.T) {
		target := EnvValue{"   some a", false}
		result, err := defineEnvVal("./testdata/env/TRIM")
		require.NoError(t, err)
		require.Equal(t, target, result)
	})
	t.Run("UNSET", func(t *testing.T) {
		target := EnvValue{"", true} // the file is fully empty
		result, err := defineEnvVal("./testdata/env/UNSET")
		require.NoError(t, err)
		require.Equal(t, target, result)
	})
}
