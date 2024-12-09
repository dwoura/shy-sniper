package task

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"strconv"
	"time"
)

type TwitterMonitor struct {
	UseridToMonitor string
	LastTweet       string
}

func NewTwitterMonitor(userid string) (*TwitterMonitor, error) {
	return &TwitterMonitor{
		userid,
		"",
	}, nil
}

func (tm *TwitterMonitor) Start() {
	ticker := time.NewTicker(3 * 60 * time.Second)
	go func() {
		for range ticker.C {
			tweet, _ := getTweetsByUserid(tm.UseridToMonitor, 1)

			if tweet == tm.LastTweet {
				continue
			}
			fmt.Println("=====新推文开始=====")
			fmt.Println(tweet)
			fmt.Println("=====新推文结束=====")
			tm.LastTweet = tweet
		}
	}()
}

func getTweetsByUserid(userid string, count int) (string, error) {
	url := "https://twitter241.p.rapidapi.com/user-tweets?user=" + userid + "&count=" + strconv.Itoa(count)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", "64b751c587msh479136fb00af56ap1e4173jsnc5034d2fad1d")
	req.Header.Add("x-rapidapi-host", "twitter241.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	data := gjson.Get(string(body), "result.timeline.instructions.#(type==TimelineAddEntries).entries.0.content.itemContent.tweet_results.result.legacy.full_text")
	return data.String(), nil
}
