package common

import "testing"

func TestRemoveLines(t *testing.T) {
	original := `1
2
3
4
5
6
7
8
9
10`

	expected := `1
2
3
4
5
6
7
9
10`

	computed := RemoveLines(original, 8, 8)

	if (computed != expected) {
		t.Error("Computed is different from expected",computed)
	}

}

func TestInsertLine(t *testing.T) {
	original := `1
2
3
4
5
6
7
8
9
10`

	expected := `1
2
3
4
5
HELLO
6
7
8
9
10`

	computed := InsertLine(original, 6, "HELLO")

	if (computed != expected) {
		t.Error("Computed is different from expected",computed)
	}

}
