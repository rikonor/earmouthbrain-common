package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	emb "github.com/rikonor/earmouthbrain"
)

type messageCaptureServer struct {
	capturedMessage string
}

func (mcs *messageCaptureServer) captureMessage(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(fmt.Sprintf("Failed to read message"))
	}

	mcs.capturedMessage = string(b)

	fmt.Fprintf(w, "OK")
}

func (mcs *messageCaptureServer) Listen(port string) {
	sMux := http.NewServeMux()
	sMux.HandleFunc("/", mcs.captureMessage)
	http.ListenAndServe(":"+port, sMux)
}

func TestHTTPMouthCanMakeRequests(t *testing.T) {
	// Start a message capturing server
	mcs := &messageCaptureServer{}
	go mcs.Listen("8080")

	hm := NewHTTPMouth("localhost:8080")

	msg := emb.StringToMessage("Test")
	hm.Say(msg)

	// Give mouth time to process message and the server to handle the request
	time.Sleep(100 * time.Millisecond)

	// Check that message was received
	if mcs.capturedMessage != "Test" {
		t.Error("Failed to send message through HTTP mouth")
	}
}

func ExampleHTTPMouth() {
	hm := NewHTTPMouth("localhost:8080")

	msg := emb.StringToMessage("Test")
	hm.Say(msg)

	// Give mouth time to process message
	time.Sleep(1 * time.Millisecond)
}

func ExampleHTTPMouth_multiple_targets() {
	hm := NewHTTPMouth("localhost:8080", "localhost:8081")

	msg := emb.StringToMessage("Test")
	hm.Say(msg)

	// Give mouth time to process message
	time.Sleep(1 * time.Millisecond)
}
