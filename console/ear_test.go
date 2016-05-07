package console

import (
	"fmt"

	emb "github.com/rikonor/earmouthbrain"
)

// func TestConsoleEarCanListenToStdin(t *testing.T) {
// 	ce := NewConsoleEar()
// 	go ce.Listen()
//
// 	var capturedMessage emb.Message
//
// 	ce.RegisterMessageHandler(func(msg emb.Message) {
// 		capturedMessage = msg
// 	})
//
// 	// TODO: Getting a bad descriptor error here
// 	// Write to Stdin
// 	_, err := os.Stdin.WriteString("Test\n")
// 	if err != nil {
// 		t.Error("Failed to write to Stdin", err)
// 	}
//
// 	// Wait so handler can run
// 	time.Sleep(1 * time.Millisecond)
//
// 	if capturedMessage != "Test" {
// 		t.Error("Failed to listen to Stdin")
// 	}
// }

func ExampleConsoleEar() {
	ce := NewConsoleEar()

	ce.RegisterMessageHandler(func(msg emb.Message) {
		fmt.Println("Test")
	})

	// Start listening on Stdin
	ce.Listen()
}
