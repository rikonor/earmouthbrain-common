package http

import (
	"fmt"
	"io/ioutil"
	"net/http"

	emb "github.com/rikonor/earmouthbrain"
)

type HTTPEar struct {
	emb.Ear
	Port string
}

func NewHTTPEar(port string) *HTTPEar {
	he := HTTPEar{Port: port}
	go he.Listen()
	return &he
}

func getHTTPHandler(msgHandler emb.MessageHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		msg := emb.ByteSliceToMessage(data)
		msgHandler(msg)
		fmt.Fprint(w, "OK")
	}
}

func (he *HTTPEar) Listen() {
	sMux := http.NewServeMux()
	sMux.HandleFunc("/", getHTTPHandler(he.RelayMessage))
	http.ListenAndServe(":"+he.Port, sMux)
}
