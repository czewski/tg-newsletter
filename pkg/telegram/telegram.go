package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/czewski/tg-newsletter/pkg/constants"
)

func SendUniqueMessage(article constants.MessageToSend) (err error) {
	url := `https://api.telegram.org/bot` + constants.ReadKey("botKey") + "/sendMessage"
	method := "POST"

	client := &http.Client{Timeout: 10 * time.Second}

	payload, err := json.Marshal(article)
	if err != nil {
		fmt.Println("Erro ao montar payload")
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println(err)
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return errors.New("erro ao enviar mensagem")
	}
	return nil
}

func ProcessMessage(feed constants.Article, search string) (message constants.MessageToSend, err error) {

	date, _ := time.Parse("2006-01-02t15:04:05z", feed.PublishedAt)

	message.ChatID = 237725036
	message.DisableWebPagePreview = false
	message.ParseMode = "HTML"
	message.Text = "<b> - [" + search + "] " + strings.Title(feed.Title) + "</b>\n\n"
	message.Text += strings.ToUpper(feed.Description[:1]) + feed.Description[1:] + "\n\n"
	message.Text += "<i> - Publicado em: " + date.Format("2006-01-02") + "</i>\n\n"
	message.Text += "<a href=" + `"` + strings.ReplaceAll(strings.ToLower(feed.URL), " ", "") + `"` + ">Link.</a>"

	return message, nil
}

func GetMessages(lastUpdateID string) (resp constants.MessageReceived, err error) {
	url := `https://api.telegram.org/bot` + constants.ReadKey("botKey") + "/getUpdates"
	method := "GET"

	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	q := req.URL.Query()
	q.Add("limit", "1")
	q.Add("offset", lastUpdateID)
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return resp, errors.New("erro ao buscar mensagem")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body, &resp)

	return resp, nil
}
