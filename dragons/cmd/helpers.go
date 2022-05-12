package main

import "strings"

func Prettier(s string) string {
	stripSlash := strings.Replace(s, "\"", "", -1)
	return stripSlash
}
