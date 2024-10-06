package composeclient

import (
	"os"
	"testing"
)

func TestCheckCompose(t *testing.T) {
	content, err := os.ReadFile("compose_test.yml")
	if err != nil {
		t.Fatal(err)
		return
	}

	err = CheckComposeFileValid(content, "halo")
	if err != nil {
		t.Fatal(err)
	}
}