package internal

import "regexp"

func FetchRegexValues(text string, regex string) []string {
	r, _ := regexp.Compile(regex)

	return r.FindAllString(text, -1)
}
