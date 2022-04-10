package bot

import (
	"fmt"
	"strings"

	mongoCon "github.com/czewski/tg-newsletter/pkg/mongo"
	"github.com/czewski/tg-newsletter/pkg/news"
	"github.com/czewski/tg-newsletter/pkg/telegram"
	"go.mongodb.org/mongo-driver/mongo"
)

//AddToFeed - Adicionar um termo a fila personalizada do USER
func AddToFeed(userID string, toAdd string, client *mongo.Client) {
	if len(toAdd) <= 0 {
		return
	}

	//Remove espacos
	toAdd = strings.TrimSpace(toAdd)

	//Insere na fila
	mongoCon.InsertInCollection(userID, toAdd, client)
}

//Now - Retorna 5 news para determinada query e determinado usuario
func Now(userID string, query string, client *mongo.Client) {
	if len(query) <= 0 {
		return
	}

	//Remove espacos
	query = strings.TrimSpace(query)

	//Seleciona news baseada na fila
	resp, _ := news.ProcessNews(query)

	if len(resp.Articles) <= 0 {
		//Sem artigos disponiveis para essa query/

	}

	//Envia news pro usuario
	for _, article := range resp.Articles {
		//Formata mensagem
		message, err := telegram.ProcessNewsToMessage(article, query)
		if err != nil {
			fmt.Println(err)
		}

		//Verifica se ja foi enviado
		sent, err := mongoCon.CheckSentNews(userID, client)
		if err != nil {
			fmt.Println(err)
		}

		var exist bool
		for _, v := range sent {
			if article.URL == v {
				exist = true
				break
			}
		}
		if exist {
			continue
		}

		//Envia para o telegram
		err = telegram.SendUniqueMessage(message)
		if err != nil {
			fmt.Println(err)
		}

		//Salva link na fila de enviados para o usuario
		err = mongoCon.InsertSentNewsToUser(article.URL, userID, client)
		if err != nil {
			fmt.Println(err)
		}
	}
}

//Feed - Retorna a fila personalizada com news atualizadas para determinado usuario
func Feed(userID string, client *mongo.Client) {
	//Busca fila especifica do usuario (feed) - "237725036"
	fila := mongoCon.ObtainResult(userID, client)

	for _, item := range fila {
		//Seleciona news baseada na fila
		resp, _ := news.ProcessNews(item)

		//Envia news pro usuario
		for _, article := range resp.Articles {
			//Formata mensagem
			message, err := telegram.ProcessNewsToMessage(article, item)
			if err != nil {
				fmt.Println(err)
			}

			//Verifica se ja foi enviado
			sent, err := mongoCon.CheckSentNews(userID, client)
			if err != nil {
				fmt.Println(err)
			}

			var exist bool
			for _, v := range sent {
				if article.URL == v {
					exist = true
					break
				}
			}
			if exist {
				continue
			}

			//Envia para o telegram
			err = telegram.SendUniqueMessage(message)
			if err != nil {
				fmt.Println(err)
			}

			//Salva link na fila de enviados para o usuario
			err = mongoCon.InsertSentNewsToUser(article.URL, userID, client)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
