package file

import (
	"fmt"
	"os"
	"testing"
	"time"

	emb "github.com/rikonor/earmouthbrain"
)

func cleanUp(filePaths ...string) {
	for _, filePath := range filePaths {
		err := os.Remove(filePath)
		if err != nil {
			panic(fmt.Sprintln("Failed to remove file:", filePath))
		}
	}
}

func TestFileEarCanReceiveFromFile(t *testing.T) {
	fe := NewFileEar("./input")
	defer cleanUp("./input")

	var capturedMessage emb.Message

	// Register handler so it can capture messages
	fe.RegisterMessageHandler(func(msg emb.Message) {
		capturedMessage = msg
	})

	// Write to file
	f, err := os.OpenFile("./input", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		t.Error("Failed to open file for writing:", err)
	}

	// Notice FileEar requires newline to distinguish between messages
	_, err = f.WriteString("Test\n")
	if err != nil {
		t.Error("Failed to write to file")
	}

	// Wait so handler can run
	time.Sleep(1 * time.Millisecond)

	if capturedMessage != "Test" {
		t.Error("Failed to listen to file")
	}
}

func ExampleFileEar() {
	fe := NewFileEar("./input")
	defer cleanUp("./input")

	var capturedMessage emb.Message

	// Register handler so it can capture messages
	fe.RegisterMessageHandler(func(msg emb.Message) {
		capturedMessage = msg
	})

	// Write to file
	f, _ := os.OpenFile("./input", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	// Notice FileEar requires newline to distinguish between messages
	f.WriteString("Test\n")

	// Wait so handler can run
	time.Sleep(1 * time.Millisecond)

	fmt.Println(capturedMessage)

	// Output: Test
}

func ExampleFileEar_multiple_files() {
	fe := NewFileEar("./input1", "./input2")
	defer cleanUp("./input1", "./input2")

	var capturedMessage emb.Message

	// Register handler so it can capture messages
	fe.RegisterMessageHandler(func(msg emb.Message) {
		capturedMessage = msg
	})

	// Write to first file
	f1, _ := os.OpenFile("./input1", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	// Notice FileEar requires newline to distinguish between messages
	f1.WriteString("Test1\n")

	// Wait so handler can run
	time.Sleep(1 * time.Millisecond)

	fmt.Println(capturedMessage)

	// Write to second file
	f2, _ := os.OpenFile("./input2", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	// Notice FileEar requires newline to distinguish between messages
	f2.WriteString("Test2\n")

	// Wait so handler can run
	time.Sleep(1 * time.Millisecond)

	fmt.Println(capturedMessage)

	// Output:
	// Test1
	// Test2
}
