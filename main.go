package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/czewski/tg-newsletter/pkg/bot"
	mongoCon "github.com/czewski/tg-newsletter/pkg/mongo"
	"github.com/czewski/tg-newsletter/pkg/telegram"
)

func main() {
	fmt.Println("Iniciando processo do servidor, as: " + time.Now().Format("2006-01-02 15:04:05"))

	//Client mongoDB
	client := mongoCon.ConnectDB()

	//Handle disconnect
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	//Busca no mongo o LastSent
	lastread, err := mongoCon.CheckLastSentTelegram("237725036", client)
	if err != nil {
		fmt.Println(err)
	}

	//Check last sent to bot
	resp, err := telegram.GetMessages(lastread)
	if err != nil {
		fmt.Println(err)
	}

	//Check message from user, see if any command was added
	if len(resp.Result) <= 0 {
		//Return?????? - Adicionar funcao pra retornar ao inicio do loop (com timer entre chamadas)
		return
	}

	for _, message := range resp.Result {
		userID := strconv.Itoa(message.Message.From.ID)

		//Faz algo com a mensagem
		if message.Message.Entities[0].Type == "bot_command" {
			uniqueMessage := message.Message.Text

			//Add - Adiciona a fila
			if strings.Contains(uniqueMessage, "/add") {
				bot.AddToFeed(userID, uniqueMessage[4:], client)
			}

			//Now - Noticias referentes a x mensagem
			if strings.Contains(uniqueMessage, "/now") {
				bot.Now(userID, uniqueMessage[4:], client)
			}

			//Feed - Chama sua fila
			if strings.Contains(uniqueMessage, "/feed") {
				bot.Feed(userID, client)
			}
		}

		//Atualizar no mongo o lastread
		err = mongoCon.UpdateLastRead(userID, strconv.Itoa(message.UpdateID+1), client)
		if err != nil {
			fmt.Println(err)
		}
	}

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
