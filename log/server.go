package log

import (
	"io/ioutil"
	slog "log"
	"net/http"
	"os"
)

var log *slog.Logger

type filelog string

func (fl filelog) Write(data []byte) (int, error) {

	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)

	if err != nil {
		return 0, err
	}
	defer f.Close()
	return fl.Write(data)
}

func Run(destination string) {
	log = slog.New(filelog(destination), "go", slog.LstdFlags)
}

func RegisterHandler() {

	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			msg, err := ioutil.ReadAll(r.Body)
			if err != nil || len(msg) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			write(string(msg))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
}

func write(msg string) {
	log.Printf("%v\n", msg)
}
