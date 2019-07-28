package get_issue_comment

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/url"
	"strconv"
)

type customPutOption struct {
	ProjectId int `url:"projectId,omitempty"`
	customUpdateOption
}

type customUpdateOption struct {
	Name    string `url:"name,omitempty"`
	Content string `url:"content,omitempty"`
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

func (c *customClient) updateWiki(wikiId float64, params customUpdateOption) (bool, error) {
	q, err := query.Values(params)
	if err != nil {
		return false, err
	}

	_, err = c.Execute("PATCH", "/api/v2/wikis/"+"42611", q)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *customClient) putResult(params customPutOption) (bool, error) {
	q, err := query.Values(params)
	if err != nil {
		return false, err
	}

	_, err = c.Post("/api/v2/wikis", q)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Wiki本文生成
func createContent(resultIssues []result) (string, string) {
	var contentPart1, contentPart2 string

	contentPart1 = `# %s

## 集計概要

- 作成日時: %s ~ %s
    - 対象件数
        - %d件
- コメントが6件を超えている
    - 対象件数
        - %d件
- 割合
    - %.1f％

## リスト

%s
`

	for _, issue := range resultIssues {
		contentPart2 += fmt.Sprintf("### %s %s %d件\n%s\n\n---\n\n", issue.issueKey, issue.summary, issue.commentNum, issue.commentMember)
	}

	return contentPart1, contentPart2
}
