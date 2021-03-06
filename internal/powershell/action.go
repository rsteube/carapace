package powershell

import (
	"encoding/json"
	"strings"

	"github.com/rsteube/carapace/internal/common"
)

var sanitizer = strings.NewReplacer( // TODO
	`$`, ``,
	"`", ``,
	"\n", ``,
	`\`, ``,
	`"`, ``,
	`'`, ``,
	`|`, ``,
	`>`, ``,
	`<`, ``,
	`&`, ``,
	`(`, ``,
	`)`, ``,
	`;`, ``,
	`#`, ``,
	`’`, ``,
	`,`, "`,",
)

func Sanitize(values ...string) []string {
	sanitized := make([]string, len(values))
	for index, value := range values {
		sanitized[index] = sanitizer.Replace(value)
	}
	return sanitized
}

func EscapeSpace(value string) string {
	return strings.Replace(value, " ", "` ", -1)
}

type completionResult struct {
	CompletionText string
	ListItemText   string
	ToolTip        string
}

// CompletionResult doesn't like empty parameters, so just replace with space if needed
func ensureNotEmpty(s string) string {
	if s == "" {
		return " "
	}
	return s
}

func ActionRawValues(callbackValue string, values ...common.RawValue) string {
	vals := make([]completionResult, 0, len(values))
	for _, val := range values {
		if val.Value != "" { // must not be empty - any empty `''` parameter in CompletionResult causes an error
			vals = append(vals, completionResult{
				CompletionText: EscapeSpace(sanitizer.Replace(val.Value)),
				ListItemText:   ensureNotEmpty(sanitizer.Replace(val.Display)),
				ToolTip:        ensureNotEmpty(sanitizer.Replace(val.Description)),
			})
		}
	}
	m, _ := json.Marshal(vals)
	return string(m)
}
