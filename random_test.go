package iridium

import (
	"fmt"
	"testing"
)

func TestRandomFloat(t *testing.T) {
	f, err := RandomFloat(16)
	if err != nil {
		t.Error(err)
	}
	t.Log(f)
	t.Log(fmt.Sprintf("%.16f", f))
}
