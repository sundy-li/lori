package utils

import "regexp"

func NamedRegexpGroup(str string, reg *regexp.Regexp) (ng map[string]string, matched bool) {
	rst := reg.FindStringSubmatch(str)
	if len(rst) < 1 {
		return
	}
	ng = make(map[string]string)
	lenRst := len(rst)
	sn := reg.SubexpNames()
	for k, v := range sn {
		if k == 0 || v == "" {
			continue
		}
		if k+1 > lenRst {
			break
		}
		ng[v] = rst[k]
	}
	matched = true
	return
}
