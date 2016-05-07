package console

import (
	"time"

	emb "github.com/rikonor/earmouthbrain"
)

func ExampleConsoleMouth() {
	cm := NewConsoleMouth()

	msg := emb.StringToMessage("Test")
	cm.Say(msg)

	// Give mouth time to process message
	time.Sleep(1 * time.Millisecond)

	// Output: Test
}
