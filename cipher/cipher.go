package cipher

import (
	"fmt"
	"strings"
)

func Caesar(s string, l int) string {
	var result []string
	for _, c := range s {
		result = append(result, rotateAscii(c, l))
	}

	return strings.Join(result, "")
}

func rotateAscii(c rune, l int) string {
	ascii := int(c)

	// LARGE
	if 65 <= ascii && ascii <= 90 {
		r := ascii - 64 + l
		m := r % 26
		return fmt.Sprintf("%c", m+64)
	}

	// small
	if 97 <= ascii && ascii <= 122 {
		r := ascii - 96 + l
		m := r % 26
		return fmt.Sprintf("%c", m+96)
	}

	return " "

}
