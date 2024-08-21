package strings

import "regexp"

func IsValidHexColor(color string) bool {
	return !regexp.MustCompile("^#([A-Fa-f0-9]{6})$").MatchString(color)
}
