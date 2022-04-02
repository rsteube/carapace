package zsh

import (
	"fmt"
	"strings"

	"github.com/rsteube/carapace/internal/common"
	"github.com/rsteube/carapace/pkg/style"
)

var sanitizer = strings.NewReplacer(
	"\n", ``,
	"\r", ``,
	"\t", ``,
	`'`, `'\''`,
)

// ActionRawValues formats values for zsh
func ActionRawValues(currentWord string, nospace bool, values common.RawValues) string {
	filtered := make([]common.RawValue, 0)

	maxLength := 0
	hasDescriptions := false
	for _, r := range values {
		if strings.HasPrefix(r.Value, currentWord) {
			filtered = append(filtered, r)
			if length := len(r.Display); length > maxLength {
				maxLength = length
			}
			hasDescriptions = hasDescriptions || r.Description != ""
		}
	}

	vals := make([]string, len(filtered))
	for index, val := range filtered {
		val.Value = sanitizer.Replace(val.Value)
		if nospace {
			val.Value = val.Value + "\001"
		}
		val.Display = sanitizer.Replace(val.Display)
		val.Description = sanitizer.Replace(val.Description)

		if strings.TrimSpace(val.Description) == "" {
			if hasDescriptions {
				vals[index] = fmt.Sprintf("%v\t%v", val.Value, style.FormatAnsi(val.Display, val.Style))
			} else {
				// TODO compadd strips display values of ansi escape codes during tabular completion (so leave that out that for now)
				vals[index] = fmt.Sprintf("%v\t%v", val.Value, val.Display)
			}
		} else {
			vals[index] = fmt.Sprintf("%v\t%v\002 %v-- %v", val.Value, style.FormatAnsi(val.Display, val.Style), strings.Repeat(" ", maxLength-len(val.Display)), val.TrimmedDescription())
		}
	}
	return strings.Join(vals, "\n")
}