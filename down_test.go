package composeclient

import (
	"os"
	"testing"
)

func TestDown(t *testing.T) {
	content, err := os.ReadFile("compose_test.yml")
	if err != nil {
		t.Fatal(err)
		return
	}

	err = Down(content, "halo")

	if err != nil {
		t.Fatal(err)
	}

}
