package console

import (
	"fmt"

	emb "github.com/rikonor/earmouthbrain"
)

type ConsoleMouth struct {
	emb.Mouth
}

func NewConsoleMouth() *ConsoleMouth {
	m := ConsoleMouth{}
	m.Init(m.OutputToConsole)
	return &m
}

func (m *ConsoleMouth) OutputToConsole(msg emb.Message) {
	fmt.Println("I have this to say:", msg)
}
