package zhong

import "testing"

func TestMottos(t *testing.T) {
	mottos := Mottos()
	if len(mottos) == 0 {
		t.Error("no mottos")
	}
}
