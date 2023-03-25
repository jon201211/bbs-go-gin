package html

import (
	"bbs-go/util/simple"
)

func GetSummary(htmlStr string, summaryLen int) string {
	if summaryLen <= 0 || simple.IsEmpty(htmlStr) {
		return ""
	}
	text := simple.GetHtmlText(htmlStr)
	return simple.GetSummary(text, summaryLen)
}
