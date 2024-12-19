package main

import (
	"bufio"
	"fmt"
	"log"
	"log/slog"
	"os"
)

type Environment map[string]EnvValue

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
		log.Fatal(err)
	}

	env := make(Environment, len(files))
	for _, file := range files {
		fmt.Println(file.Name())
		env[file.Name()] = defineEnvVal(file.Name())

	}
	return nil, nil
}

func defineEnvVal(fileName string) EnvValue {
	file, err := os.Open(fileName)
	if err != nil {
		slog.Error("error with os.Open", err)
	}
	defer func() {
		err := file.Close()
		slog.Error("error with os.File().Close", err)
	}()
	val := func() string {
		reader := bufio.NewReader(file)
		str, _, err := reader.ReadLine()
		if err != nil {
			slog.Error("error with bufio.NewReader(file).ReadLine%w", err)
		}
		return string(str)
	}()

	return EnvValue{val, false}
}
