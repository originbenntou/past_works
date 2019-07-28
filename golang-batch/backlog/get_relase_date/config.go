package main

const (
	BACKLOGURL    = "https://oribenn.backlog.com"
	BEFOREDATE    = 15
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
	keywords = []string{
		"リリース",
		"商用",
		"反映",
		"デプロイ",
	}
)
