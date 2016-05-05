package http

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	emb "github.com/rikonor/earmouthbrain"
)

type HTTPMouth struct {
	emb.Mouth
	// Host - format <host>:<port>
	// Example - localhost:8080
	Hosts []string
}

func NewHTTPMouth(hosts ...string) *HTTPMouth {
	m := HTTPMouth{Hosts: hosts}
	m.Init(m.OutputToHTTP)
	return &m
}

func (m *HTTPMouth) OutputToHTTP(msg emb.Message) {
	for _, host := range m.Hosts {
		_, err := http.Post("http://"+host, "text/plain", strings.NewReader(string(msg)))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to send message to HTTP endpoint: http://%s: %s\n", host, err)
		}
	}
}
