package utils

import (
	"fmt"
	"strings"
)

func DecorateMessage(message string) string {
	return fmt.Sprintf("%s%s%s",
		strings.Repeat("=", 10),
		message,
		strings.Repeat("=", 10),
	)
}
