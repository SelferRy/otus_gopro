package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

func (e Environment) StringSlice() []string {
	res := make([]string, 0, len(e))
	for key, envVal := range e {
		res = append(res, fmt.Sprintf("%v=%v", key, envVal.Value))
	}
	return res
}

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		slog.Info("error os.ReadDir\n%w", slog.Any("error", err))
		return nil, err
	}

	env := make(Environment, len(files))
	for _, file := range files {
		fileName := filepath.Join(dir, file.Name())
		env[file.Name()], err = defineEnvVal(fileName)
		if err != nil {
			return nil, err
		}
	}
	return env, nil
}

func defineEnvVal(fileName string) (EnvValue, error) {
	file, err := os.Open(fileName)
	if err != nil {
		slog.Error("error with os.Open", slog.Any("error", err))
		return EnvValue{}, err
	}
	defer func() {
		err = file.Close()
		if err != nil {
			slog.Error("error with os.File().Close", slog.Any("error", err))
		}
	}()

	// handle fully empty case
	var fileInfo os.FileInfo
	if fileInfo, err = file.Stat(); err == nil && fileInfo.Size() == 0 {
		return EnvValue{"", true}, nil
	} else if err != nil {
		slog.Error("error file.Stat()\n%w", slog.Any("error", err))
		return EnvValue{}, err
	}

	// handle other cases
	val, errVal := func() (string, error) {
		reader := bufio.NewReader(file)
		var line []byte
		line, _, err = reader.ReadLine()
		if err != nil {
			slog.Error("error with bufio.NewReader(file).ReadLine\n%w", slog.Any("error", err))
			return "", err
		}
		str := strings.TrimRight(string(line), " \n\t") // look at EMPTY and TRIM cases
		str = strings.ReplaceAll(str, "\x00", "\n")     // look at FOO case
		return str, nil
	}()
	if errVal != nil {
		return EnvValue{}, errVal
	}

	return EnvValue{val, false}, nil
}
