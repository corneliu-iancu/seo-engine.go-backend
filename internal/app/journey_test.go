package app

import (
	"fmt"
	"regexp"
	"testing"
)

func TestApplication_GetMatch(t *testing.T) {
	str := "{country}"
	re := regexp.MustCompile(`{[a-zA-Z^0-9]*?\}`)

	matches := re.FindAllString(str, -1)

	fmt.Println(len(matches) == 1)
}
