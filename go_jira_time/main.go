package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	url := "https://originbenntou.atlassian.net/rest/agile/1.0/board/1/issue?worklog"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// ヘッダー情報
	req.SetBasicAuth("hoge", "hogehoge")
	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	byteArray, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// JSONデコード
	var decorded interface{}

	// なぜポインタを渡すのか何7日
	if err := json.Unmarshal(byteArray, &decorded); err != nil {
		log.Fatalln(err)
	}

	result := make(map[string]float64)

	// TODO: JIRA_APIを勉強して出直す
	for k, data := range decorded.(map[string]interface{}) {
		if k == "issues" {
			for _, data2 := range data.([]interface{}) {
				for k3, data3 := range data2.(map[string]interface{}) {
					if k3 == "fields" {
						for k4, data4 := range data3.(map[string]interface{}) {
							if k4 == "worklog" {
								for k5, data5 := range data4.(map[string]interface{}) {
									if k5 == "worklogs" {
										for _, data6 := range data5.([]interface{}) {
											worklog := data6.(map[string]interface{})
											//if time.Now().AddDate(0, 0, -1).Unix() < worklog["updated"]. && worklog["updated"] < time.Now() {
											//	result[worklog["updateAuthor"].(map[string]interface{})["displayName"].(string)] += worklog["timeSpentSeconds"].(float64)
											//}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(result)
}
