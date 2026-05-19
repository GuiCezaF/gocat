package reader

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type File struct {
	filePath string
	data     []byte
	lines    int
	loaded   bool
}

func NewFile(path string) *File {
	return &File{
		filePath: path,
	}
}

func (f *File) load() error {
	if f.loaded {
		return nil
	}

	data, err := os.ReadFile(f.filePath)
	if err != nil {
		return fmt.Errorf("erro ao ler arquivo: %w", err)
	}

	f.data = data

	scanner := bufio.NewScanner(bytes.NewReader(data))

	lines := 0

	for scanner.Scan() {
		lines++
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("erro ao contar linhas: %w", err)
	}

	f.lines = lines
	f.loaded = true

	return nil
}

func (f *File) Read() ([]byte, error) {
	if err := f.load(); err != nil {
		return nil, err
	}

	return f.data, nil
}

func (f *File) CountLines() (int, error) {
	if err := f.load(); err != nil {
		return 0, err
	}

	return f.lines, nil
}

