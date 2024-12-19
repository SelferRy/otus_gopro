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
	t.Run("read directory", func(t *testing.T) {
		target := EnvValue{"bar", false}
		result := defineEnvVal("./testdata/env/BAR")
		require.Equal(t, target, result)
	})
}
