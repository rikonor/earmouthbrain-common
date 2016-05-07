package http

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	emb "github.com/rikonor/earmouthbrain"
)

func TestHTTPEarCanListen(t *testing.T) {
	he := NewHTTPEar("8081")

	var capturedMessage emb.Message

	he.RegisterMessageHandler(func(msg emb.Message) {
		capturedMessage = msg
	})

	// Send the ear a message
	_, err := http.Post("http://localhost:8081", "text/plain", strings.NewReader("Test"))
	if err != nil {
		t.Error("Failed to connect to http ear")
	}

	// Give ear time to handle the request
	time.Sleep(100 * time.Millisecond)

	if capturedMessage != emb.StringToMessage("Test") {
		t.Error("Failed to send ear message")
	}
}

func ExampleHTTPEar() {
	he := NewHTTPEar("8081")

	he.RegisterMessageHandler(func(msg emb.Message) {
		fmt.Println(msg)
	})
}
