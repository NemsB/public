package correct

import "strings"

func ConcatParams(args []string) string {
	return strings.Join(args, "\n")
}
