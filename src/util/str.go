package util

import "regexp"

var doubleBracketRegexp = regexp.MustCompile(`\{\{([^}]+)}}`)

func GetValueFromBracket(data string) []string {
	subMatch := doubleBracketRegexp.FindAllStringSubmatch(data, -1)
	var res []string
	for _, re := range subMatch {
		if len(re) < 2 {
			continue
		}
		res = append(res, re[1])
	}
	return res
}
