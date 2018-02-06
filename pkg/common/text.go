package common

import "strings"

func RemoveLines(text string, start int, end int) (temp string) {
	splitted := strings.Split(text, "\n")
	for lineNo, line := range splitted {
		if !(lineNo >= start-1 && lineNo <= end-1) {
			temp += line
			if lineNo < len(splitted)-1 {
				temp += "\n"
			}
		}
	}
	return temp
}

func InsertLine(text string, index int, iLine string) (temp string) {
	splitted := strings.Split(text, "\n")
	for lineNo, line := range splitted {
		if lineNo == (index - 1) {
			temp += iLine + "\n" + line
		} else {
			temp += line
		}
		if lineNo < len(splitted)-1 {
			temp += "\n"
		}
	}
	return temp
}
