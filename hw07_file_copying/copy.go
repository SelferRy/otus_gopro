package main

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrEmptyFile             = errors.New("file is empty")
	ErrSameFiles             = errors.New("files are same")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fromFile, fromStat, errFromFile := makeReadAttrs(fromPath)
	if errFromFile != nil {
		return fmt.Errorf("error with makeReadAttrs.\n%w", errFromFile)
	}
	defer func() {
		if err := fromFile.Close(); err != nil {
			slog.Any("error with fromFile.Close()", err)
		}
	}()

	toFile, errToFile := os.Create(toPath)
	if errToFile != nil {
		return fmt.Errorf("problems with file for recording.\n%w", errToFile)
	}
	defer func() {
		if err := toFile.Close(); err != nil {
			slog.Any("error with toFile.Close()", err)
		}
	}()

	sameFiles := func() (bool, error) {
		toStat, errToStat := toFile.Stat()
		if errToStat != nil {
			return false, fmt.Errorf("error with toFile.Stat().\n%w", errToStat)
		}
		return os.SameFile(fromStat, toStat), nil
	}
	if areSame, errSame := sameFiles(); errSame != nil {
		return errSame
	} else if areSame {
		return ErrSameFiles
	}

	if offset > 0 {
		_, err := fromFile.Seek(offset, io.SeekStart)
		if err != nil {
			return fmt.Errorf("error with fromFIle.Seek.\n%w", err)
		}
	}

	limit = defineLimit(fromStat.Size(), limit, offset)

	if _, err := io.CopyN(toFile, fromFile, limit); err != nil {
		return fmt.Errorf("copy error: %w", err)
	}
	return nil
}

func makeReadAttrs(fromPath string) (*os.File, os.FileInfo, error) {
	fromFile, errFromFile := os.Open(fromPath)
	if errFromFile != nil {
		return nil, nil, fmt.Errorf("error with os.Open(fromPath).\n%w", errFromFile)
	}
	fromStat, errFromStat := fromFile.Stat()
	if errFromStat != nil {
		return nil, nil, fmt.Errorf("error with fromFile.Stat().\n%w", errFromStat)
	}
	if errFromCheck := checkFromFile(fromStat, offset); errFromCheck != nil {
		return nil, nil, errFromCheck
	}
	return fromFile, fromStat, nil
}

func checkFromFile(fromStat os.FileInfo, offset int64) error {
	if !fromStat.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	size := fromStat.Size()
	switch {
	case size == 0:
		return fmt.Errorf("size = 0.\n%w", ErrEmptyFile)
	case size < offset:
		slog.Int64("size = ", size)
		slog.Int64("offset = ", offset)
		return fmt.Errorf("offset is too much.\n%w", ErrOffsetExceedsFileSize)
	default:
		return nil
	}
}

func defineLimit(size, limit, offset int64) int64 {
	if limit < 1 || size < offset+limit {
		return size - offset
	}
	return limit
}
