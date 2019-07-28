package get_issue_comment

import (
	"encoding/json"
	backlog "github.com/griffin-stewie/go-backlog"
	"net/url"
	"strconv"
)

type customIssueOption struct {
	backlog.IssuesOption
	UpdatedSince string `url:"updatedSince,omitempty"`
	UpdatedUntil string `url:"updatedUntil,omitempty"`
}

func (c *customClient) getIssuesWithCustomOption(option customIssueOption) (backlog.IssueSlice, error) {
	params, err := option.Values()
	if err != nil {
		return nil, err
	}

	bytes, err := c.Get("/api/v2/issues/", params)
	if err != nil {
		return nil, err
	}

	var issues backlog.IssueSlice
	json.Unmarshal(bytes, &issues)

	return issues, nil
}

func (c *customClient) getMemberCommentIssue(issue *backlog.Issue) (bool, []string) {
	id := strconv.Itoa(*issue.ID)
	bytes, _ := c.Get("/api/v2/issues/"+id+"/comments/", url.Values{})

	// 手っ取り早くインターフェースで対応
	var res []map[string]interface{}
	json.Unmarshal(bytes, &res)

	var isContain bool
	var commentNames []string
	for _, v := range res {
		commentName := v["createdUser"].(map[string]interface{})["name"].(string)
		commentNames = append(commentNames, commentName)

		// コメントした人がメンバーにいるかどうか
		for _, memberName := range member {
			if memberName == commentName {
				isContain = true
			}
		}
	}

	return isContain, getCommenterPointList(commentNames)
}

func (c *customClient) getOverCount(issue *backlog.Issue) (int, bool) {
	id := strconv.Itoa(*issue.ID)
	bytes, _ := c.Get("/api/v2/issues/"+id+"/comments/count", url.Values{})

	res := make(map[string]int)
	json.Unmarshal(bytes, &res)

	return res["count"], res["count"] > MAXCOMMENT
}

// コメントした人の名前をリスト化して返す
func getCommenterPointList(commenterList []string) (display []string) {
	list := make(map[string]int)
	for _, name := range commenterList {
		list[name]++
	}

	// 表示用に整形
	for name, point := range list {
		display = append(display, name+":"+strconv.Itoa(point))
	}

	return
}
