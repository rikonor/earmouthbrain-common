package console

import (
	"fmt"

	emb "github.com/rikonor/earmouthbrain"
)

// ConsoleEar - listens to input from the console
type ConsoleEar struct {
	emb.Ear
}

// NewConsoleEar - Create a new ConsoleEar
func NewConsoleEar() *ConsoleEar {
	ce := ConsoleEar{}
	return &ce
}

func (ce *ConsoleEar) Listen() {
	for {
		var msgText string
		fmt.Scanln(&msgText)

		msg := emb.StringToMessage(msgText)
		ce.RelayMessage(msg)
	}
}
