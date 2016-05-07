package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	emb "github.com/rikonor/earmouthbrain"
)

func TestFileMouthCanOutputToFile(t *testing.T) {
	fm := NewFileMouth("./output")
	defer cleanUp("./output")

	msg := emb.StringToMessage("Test")
	fm.Say(msg)

	// Give mouth time to process message
	time.Sleep(1 * time.Millisecond)

	// Open file and check content
	f, err := os.Open("./output")
	if err != nil {
		t.Error("Failed to open file")
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error("Failed to read file")
	}

	// Notice FileMouth appends newline
	if string(b) != "Test\n" {
		t.Error("Failed to write desired message to file")
	}
}

func ExampleFileMouth() {
	fm := NewFileMouth("./output")
	defer cleanUp("./output")

	msg := emb.StringToMessage("Test")
	fm.Say(msg)

	// Give mouth time to process message
	time.Sleep(1 * time.Millisecond)

	// Open file and check content
	f, _ := os.Open("./output")
	b, _ := ioutil.ReadAll(f)

	fmt.Println(string(b))

	// Output: Test
}
