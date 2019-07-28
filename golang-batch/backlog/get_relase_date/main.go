package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	backlog "github.com/griffin-stewie/go-backlog"
)

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
	log.Println("Get_RELEASE_DATE Start")

	now := time.Now().Format("2006-01-02")
	since := time.Now().Add(time.Duration(-BEFOREDATE*24) * time.Hour).Format("2006-01-02")

	// プロジェクトIDごとの課題を取得
	var issueListByPids []backlog.IssueSlice

	log.Println("Get Issues By ProjectId")

	var wg1 sync.WaitGroup
	for _, v := range projectIds {
		wg1.Add(1)

		// GETパラメータセット
		var option customIssueOption
		option.ProjectIDs = []int{v}
		option.Count = ISSUESMAX
		option.CreatedSince = since
		option.CreatedUntil = now

		// FIXME: 以下のようにするとエラーになって辛い
		//option := customIssueOption{
		//	projectIds: []int{v},
		//	...
		//}

		go func() {
			defer wg1.Done()

			issues, err := client.getIssuesWithCustomOption(option)
			if err != nil {
				log.Fatalf("ERROR: %s", err.Error())
			}

			if len(issues) > 0 {
				issueListByPids = append(issueListByPids, issues)
			}
		}()
	}
	wg1.Wait()

	var count int
	for _, issueList := range issueListByPids {
		count += len(issueList)
	}

	log.Println("Issues Count is", count)

	var result string
	var parentSummary string

	for _, issueList := range issueListByPids {
		for _, issue := range issueList {
			for _, word := range keywords {
				if strings.Contains(*issue.Summary, word) {

					//if issue.ParentIssueID != nil {
					//	parentSummary = client.getTargetIssueSummary(*issue.ParentIssueID)
					//}

					parentSummary = client.getTargetIssueSummary(1023785)

					result += "### " + *issue.Summary + " " + BACKLOGURL + "/view/" + *issue.IssueKey + "\n" +
						"親課題: " + parentSummary + "\n" +
						"期限: " + issue.DueDate.Format("2006-01-02")
				}
			}
		}
	}

	if result == "" {
		result += "直近のリリースはありません"
	}

	fmt.Println(result)
}
