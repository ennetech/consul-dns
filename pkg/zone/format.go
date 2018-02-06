package zone

import (
	"fmt"
	"regexp"
	"strings"
)

// Leaving this here as a base for zone parsing/formatting
func parse(s string) {

	rx := regexp.MustCompile("(?P<name>.*)[ \t]+(?P<ttl>[0-9]+)[ \t]+IN[ \t]+(?P<type>[A-Z]+)[ \t]+(?P<value>.*)")

	for _, line := range strings.Split(s, "\n") {
		fmt.Println(line)
		match := rx.FindStringSubmatch(line)
		result := make(map[string]string)
		for i, name := range rx.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}
		fmt.Println(match)
	}

}
