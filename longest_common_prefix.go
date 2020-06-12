package main

import (
	"strings"
)

func LongestUniqueSubstring(strs []string) string {
	switch len(strs) {
	case 0:
		return ""
	case 1:
		return strs[0]
	}

	prefix := strs[0]

	for _, str := range strs[1:] {
		for strings.Index(str, prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}
