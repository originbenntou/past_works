package get_issue_comment

const (
	BACKLOGURL    = "https://oribenn.backlog.com"
	BEFOREDATE    = 18
	MAXCOMMENT    = 0
	ISSUESMAX     = 100
	GOROUTINMAX   = 10
	PROJECTIDWIKI = 18772
	WIKITITLE     = "hogehoge/hoge/"
)

// 配列は定数にできないため変数にする
var (
	// メイン課題のみ
	projectIds = []int{18772}
	// 開発メンバーのみ
	member = []string{
		"山本司",
	}
)
