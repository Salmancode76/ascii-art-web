package PKG

import (
	"fmt"
	"strings"
)

func PrintA(arr *[8]string) string {
	var textBuilder strings.Builder

	for _, line := range arr {
		fmt.Println(line)
		textBuilder.WriteString("\n")

		textBuilder.WriteString(line)
	}

	return textBuilder.String()
}
