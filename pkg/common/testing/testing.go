package testing

import (
	"reflect"
	"strconv"
	"testing"
)

func AssertLen(t *testing.T, l interface{}, count int) {
	ln := reflect.ValueOf(l).Len()
	if ln != count {
		t.Error("The count is " + strconv.Itoa(ln) + " not " + strconv.Itoa(count))
	}
}
