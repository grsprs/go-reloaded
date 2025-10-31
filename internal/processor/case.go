package processor

import (
	"regexp"
	"strings"
)

// applyCaseTransformations applies (low), (up), (cap) modifiers including quoted text
func applyCaseTransformations(text string) string {
	re := regexp.MustCompile(`'?([\p{L}\s]+?)'?\s*\((low|up|cap)(?:,\s*(\d+))?\)`)

	return re.ReplaceAllStringFunc(text, func(match string) string {
		submatches := re.FindStringSubmatch(match)
		if len(submatches) < 3 {
			return match
		}

		content := strings.TrimSpace(submatches[1])
		modifier := strings.ToLower(submatches[2])

		switch modifier {
		case "low":
			content = strings.ToLower(content)
		case "up":
			content = strings.ToUpper(content)
		case "cap":
			words := strings.Fields(strings.ToLower(content))
			for i := range words {
				if len(words[i]) > 0 {
					words[i] = strings.ToUpper(string(words[i][0])) + words[i][1:]
				}
			}
			content = strings.Join(words, " ")
		}

		if strings.Contains(match, "'") {
			return "'" + content + "'"
		}
		return content
	})
}
