package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"

	backlog "github.com/griffin-stewie/go-backlog"
)

const (
	BACKLOGURL = "https://oribenn.backlog.com"
	PROJECTID  = 18772
)

type customClient struct {
	backlog.Client
}

type branch struct {
	title string
	next  []branch
}

func newClient(baseURL *url.URL, APIKey string) (c customClient) {
	c.BaseURL = baseURL
	c.APIKey = APIKey

	return
}

var client customClient

func init() {
	token := os.Getenv("BACKLOG_TOKEN")
	if token == "" {
		log.Fatalln("You need Backlog access token.")
	}

	URL, err := url.Parse(BACKLOGURL)
	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}

	client = newClient(URL, token)
}

func main() {
	var tree []branch

	wikiList := client.getWikiNameList(PROJECTID)
	fmt.Println(wikiList)
	for title, _ := range wikiList {
		splitTitles := strings.Split(title, "/")
		if len(tree) == 0 {
			// ループの初回は無条件ですべてツリーに追加
			tree = append(tree, createBranch(splitTitles))
		} else {
			recursiveOrder(&tree, splitTitles)
		}
	}
	fmt.Println(tree)
}

func (c *customClient) getWikiNameList(projectId int) map[string]float64 {
	q := url.Values{}
	q.Add("projectIdOrKey", strconv.Itoa(projectId))
	bytes, _ := c.Get("/api/v2/wikis", q)

	var wikis []map[string]interface{}
	json.Unmarshal(bytes, &wikis)

	wikiList := make(map[string]float64)
	for _, wiki := range wikis {
		if wikiName := wiki["name"].(string); wikiName != "" {
			wikiList[wikiName] = wiki["id"].(float64)
		}
	}

	return wikiList
}

func recursiveOrder(bs *[]branch, ss []string) {
	fmt.Println(bs)
	if len(*bs) == 0 {
		*bs = append(*bs, createBranch(ss))
		return
	}

	for _, b := range *bs {
		fmt.Println(b.title, ss[0])
		if ss[0] == b.title {
			// タイトルが存在すればその下の階層を再帰的にチェックする
			recursiveOrder(&b.next, ss[1:])
		}
	}

	// 今いる階層にタイトルが存在しなければ新しいブランチを追加
	*bs = append(*bs, createBranch(ss))
	return
}

func createBranch(ss []string) (b branch) {
	var bs []branch
	for _, s := range ss {
		bs = append(bs, branch{s, []branch{}})
	}

	var tmp branch
	for i := 0; i < len(bs); i++ {
		b = bs[i]
		tmp = b.next[0]
	}
	return
}
