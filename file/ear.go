package file

import (
	"fmt"
	"os"

	"github.com/hpcloud/tail"
	emb "github.com/rikonor/earmouthbrain"
)

type FileEar struct {
	emb.Ear
	FilePaths []string
}

// NewFileEar - Create a new FileEar
func NewFileEar(filePaths ...string) *FileEar {
	for _, filePath := range filePaths {
		if _, err := os.Stat(filePath); err != nil {
			if _, err := os.Create(filePath); err != nil {
				panic(fmt.Sprintf("Failed to create file: %s", filePath))
			}
		}
	}

	fe := FileEar{
		FilePaths: filePaths,
	}

	go fe.Listen()

	return &fe
}

func (fe *FileEar) Listen() {
	for _, filePath := range fe.FilePaths {
		go func(filePath string) {
			t, err := tail.TailFile(filePath, tail.Config{Follow: true})
			if err != nil {
				panic(fmt.Sprintf("Failed to tail file: %s", filePath))
			}

			for msgText := range t.Lines {
				msg := emb.StringToMessage(msgText.Text)

				fe.RelayMessage(msg)
			}
		}(filePath)
	}
}
