package get_issue_comment

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"sync"
	"time"

	backlog "github.com/griffin-stewie/go-backlog"
)

type result struct {
	issueKey      string
	summary       string
	commentNum    int
	commentMember []string
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
	log.Println("Get_Issue_Comment Start")

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
		option.UpdatedSince = since

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

	// 条件に一致する課題を取得
	// メンバーがコメントしている かつ コメントのレス数が6を超えている
	var resultIssues []result

	log.Println("Get Issues By MemberComment and OverComment")

	// 並列処理は最大10
	limit := make(chan struct{}, GOROUTINMAX)

	var wg2 sync.WaitGroup
	for _, issueList := range issueListByPids {
		for _, issue := range issueList {
			wg2.Add(1)

			go func(issue *backlog.Issue) {
				defer wg2.Done()
				limit <- struct{}{}

				existMemberComment, memberNames := client.getMemberCommentIssue(issue)
				if existMemberComment {
					commentNum, isOver := client.getOverCount(issue)
					if isOver {
						resultIssues = append(resultIssues, result{
							issueKey:      *issue.IssueKey,
							summary:       *issue.Summary,
							commentNum:    commentNum,
							commentMember: memberNames,
						})
					}
				}

				<-limit
			}(issue)
		}
		wg2.Wait()
	}

	log.Println("Target Issues Count is", len(resultIssues))

	// 結果をWikiへ出力
	contentPart1, contentPart2 := createContent(resultIssues)

	var (
		updateOk bool
		err      error
	)
	for k, v := range client.getWikiNameList(PROJECTIDWIKI) {
		// タイトルが一致するWikiが既にあれば更新
		if k == WIKITITLE+now {
			// update用リクエストパラメータセット
			var updateOption customUpdateOption

			updateOption.Name = WIKITITLE + now
			updateOption.Content = fmt.Sprintf(contentPart1, updateOption.Name, since, now, count, len(resultIssues), float64(len(resultIssues)/count*100), contentPart2)

			updateOk, err = client.updateWiki(v, updateOption)
			if err != nil {
				log.Fatalf("ERROR: %s", err.Error())
			}

			log.Println("Update Result Success")
			break
		}
	}

	var putOk bool

	// 更新がなければ新規作成
	if !updateOk {
		// put用リクエストパラメータセット
		var putOption customPutOption

		putOption.ProjectId = PROJECTIDWIKI
		putOption.Name = WIKITITLE + now
		putOption.Content = fmt.Sprintf(contentPart1, putOption.Name, since, now, count, len(resultIssues), float64(len(resultIssues)/count*100), contentPart2)

		putOk, err = client.putResult(putOption)
		if err != nil {
			log.Fatalf("ERROR: %s", err.Error())
		}
	}

	if putOk {
		log.Println("Put Result Success")
	}

	log.Println("Get_Issue_Comment Finish")
}
