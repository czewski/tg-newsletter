package main

import (
	"fmt"
	"time"

	"github.com/czewski/tg-newsletter/pkg/server"
)

func main() {
	fmt.Println("Iniciando processo do servidor, as: " + time.Now().String())
	server.CreateServer()

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

		//server.CreateServer()*/
}
