package main

import (
	"context"
	"fmt"
	"time"

	mongoCon "github.com/czewski/tg-newsletter/pkg/mongo"
	"github.com/czewski/tg-newsletter/pkg/news"
	"github.com/czewski/tg-newsletter/pkg/telegram"
)

func main() {
	fmt.Println("Iniciando processo do servidor, as: " + time.Now().Format("2006-01-02 15:04:05"))
	//server.CreateServer()
	//just make the server get every x seconds, fuck https
	//telegram.GetMessages("165466380")
	//telegram.Sender()
	client := mongoCon.ConnectDB()

	//Busca fila
	fila := mongoCon.ObtainResult("237725036", client)

	//Handle disconnect
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	fmt.Println(fila)

	for _, v := range fila {
		//Seleciona news baseada na fila
		resp, _ := news.ProcessNews(v)

		//Envia news pro usuario
		for _, article := range resp.Articles {
			//Formata mensagem
			message, err := telegram.ProcessNewsToMessage(article, v)
			if err != nil {
				fmt.Println(err)
			}

			//Verifica se ja foi enviado
			sent, err := mongoCon.CheckSentNews("237725036", client)
			var exist bool
			for _, v := range sent {
				if article.URL == v {
					exist = true
					break
				}
			}
			if exist == true {
				continue
			}

			//Envia para o telegram
			err = telegram.SendUniqueMessage(message)
			if err != nil {
				fmt.Println(err)
			}

			//Salva link na fila de enviados para o usuario
			err = mongoCon.InsertSentNewsToUser(article.URL, "237725036", client)
			if err != nil {
				fmt.Println(err)
			}
		}

	}

	//Insere na fila
	//mongoCon.InsertInCollection("237725036", "CPLE6", client)
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
