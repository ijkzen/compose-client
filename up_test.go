package composeclient

import (
	"os"
	"testing"
)

func TestUp(t *testing.T) {
	content, err := os.ReadFile("compose_test.yml")
	if err != nil {
		t.Fatal(err)
		return
	}

	err = Up(content, "halo")

	if err != nil {
		t.Fatal(err)
		return
	}
}
