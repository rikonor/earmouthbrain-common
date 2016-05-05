package file

import (
	"fmt"
	"os"

	emb "github.com/rikonor/earmouthbrain"
)

type FileMouth struct {
	emb.Mouth
	FilePath string
	file     *os.File
}

func NewFileMouth(filePath string) *FileMouth {
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file for writing: %s", err))
	}

	// Notice file is never closed

	m := FileMouth{
		FilePath: filePath,
		file:     f,
	}

	m.Init(m.OutputToFile)
	return &m
}

func (m *FileMouth) OutputToFile(msg emb.Message) {
	msgText := string(msg) + "\n"
	m.file.WriteString(msgText)
}
