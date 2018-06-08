package go_common

import "strings"

type MatchFunc func(s string) bool

func lit(s string) MatchFunc {
	return func(b string) bool {
		return s == b
	}
}

func sequence(a, b MatchFunc) MatchFunc {

	return func(s string) bool {
		for i := 0; i < len(s)-1; i++ {
			if a(s[0:i]) && b(s[i:]) {
				return true
			}
		}
		return false
	}
}

func either(a, b MatchFunc) MatchFunc {

	return func(s string) bool {
		return a(s) || b(s)
	}
}

func any() MatchFunc {

	return func(s string) bool {
		return true
	}
}

func oneof(s string) MatchFunc {

	return func(b string) bool {
		return strings.Contains(s, b)
	}
}
