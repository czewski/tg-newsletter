package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/czewski/tg-newsletter/pkg/constants"
)

func SendMessage(article constants.MessageToSend) (err error) {
	url := `https://api.telegram.org/bot` + constants.ReadKey("botKey") + "/sendMessage"
	method := "GET"

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

func ProcessMessage(feed constants.Article) (message constants.MessageToSend, err error) {

	date, _ := time.Parse("2006-01-02t15:04:05z", feed.PublishedAt)

	message.ChatID = 237725036
	message.DisableWebPagePreview = false
	message.ParseMode = "HTML"
	message.Text = "<b> - " + strings.Title(feed.Title) + "</b>\n\n"
	message.Text += strings.ToUpper(feed.Description[:1]) + feed.Description[1:] + "\n\n"
	message.Text += "<i> - " + date.Format("2006-01-02") + "</i>\n\n"
	message.Text += "<a href=" + `"` + strings.ReplaceAll(strings.ToLower(feed.URL), " ", "") + `"` + ">Link.</a>"

	return message, nil
}
