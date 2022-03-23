package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/czewski/tg-newsletter/pkg/telegram"
)

const (
	htmlIndex = `<html><body>Welcome!</body></html>`
	httpPort  = "134.122.13.88:443"
)

var (
	flgProduction          = false
	flgRedirectHTTPToHTTPS = false
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, htmlIndex)
}

func makeServerFromMux(mux *http.ServeMux) *http.Server {
	// set timeouts so that a slow or malicious client doesn't
	// hold resources forever
	return &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
}

func makeHTTPServer() *http.Server {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", handleIndex)
	return makeServerFromMux(mux)

}

func makeHTTPToHTTPSRedirectServer() *http.Server {
	handleRedirect := func(w http.ResponseWriter, r *http.Request) {
		newURI := "https://" + r.Host + r.URL.String()
		http.Redirect(w, r, newURI, http.StatusFound)
	}
	mux := &http.ServeMux{}
	mux.HandleFunc("/", handleRedirect)
	return makeServerFromMux(mux)
}

func parseFlags() {
	flag.BoolVar(&flgProduction, "production", false, "if true, we start HTTPS server")
	flag.BoolVar(&flgRedirectHTTPToHTTPS, "redirect-to-https", false, "if true, we redirect HTTP to HTTPS")
	flag.Parse()
}

func main() {
	fmt.Println("Iniciando processo do servidor, as: " + time.Now().String())
	//server.CreateServer()
	//just make the server get every x seconds, fuck https
	//telegram.GetMessages("165466380")
	telegram.Sender()
}

/*
		var search = "BBAS3"
		//resp, _ := news.ProcessNews("BBAS3")
		resp := constants.NewsResult{}
		err := json.Unmarshal([]byte(string(constants.Newsss)), &resp)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(resp)
		//fmt.Println(len(resp.Articles))
		for _, article := range resp.Articles {

			message, err := telegram.ProcessMessage(article, search)
			if err != nil {
				fmt.Println(err)
			}

			asd, _ := json.Marshal(message)
			fmt.Println(string(asd))

			err = telegram.SendMessage(message)
			if err != nil {
				fmt.Println(err)
			}
		}

		//server.CreateServer()
}
*/
