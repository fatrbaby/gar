package shortcut

import "strings"

func CleanKeywords(keywords []string) []string {
	wd := make([]string, 0, len(keywords))

	for _, word := range keywords {
		w := strings.TrimSpace(strings.TrimSpace(word))

		if len(w) > 0 {
			wd = append(wd, w)
		}
	}

	return wd
}
