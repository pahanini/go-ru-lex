package rulex

import (
	"strings"
)

// HasAnyPrefix returns true and prefix value if string has any prefix
func HasAnyPrefix(s string, prefix []string) (bool, string) {
	for _, p := range prefix {
		if strings.HasPrefix(s, p) {
			return true, p
		}
	}
	return false, ""
}

// Is returns true if rune is in s string
func Is(r rune, s string) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}
	return false;
}

// Has returns true if runes has at least one rune of string
func Has(runes []rune, s string) bool {
	for _, r := range runes {
		if Is(r, s) {
			return true
		}
	}
	return false;
}





