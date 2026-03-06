package modules

import (
	"regexp"
)

func GetLinks(body string) [][]string {
	re := regexp.MustCompile(`href="([^"]*)"`)
	matches := re.FindAllStringSubmatch(body, -1)
	return matches
}
