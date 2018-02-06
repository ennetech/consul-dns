package common

import (
	"io/ioutil"
	"os"
)

func ReadFile(filename string) []byte {
	b, err := ioutil.ReadFile(filename) // just pass the file name
	if err != nil {
		panic(err)
	}
	return b
}

func WriteFile(filename string, content []byte) bool {
	err := ioutil.WriteFile(filename, content, os.ModePerm) // just pass the file name
	if err != nil {
		return false
	}
	return true
}
