package news

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/czewski/tg-newsletter/pkg/constants"
)

func ProcessNews(field string) (resp constants.NewsResult, err error) {
	url := `https://newsapi.org/v2/everything`
	method := "GET"

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	q := req.URL.Query()
	q.Add("q", field)
	q.Add("from", time.Now().AddDate(0, 0, -10).Format("2006-01-02"))
	q.Add("to", time.Now().AddDate(0, 0, 0).Format("2006-01-02"))
	q.Add("sortBy", "popularity")
	q.Add("pageSize", "5")
	req.URL.RawQuery = q.Encode()

	req.Header.Add("x-api-key", constants.ReadKey("newsKey"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body, &resp)
	return resp, err
}
