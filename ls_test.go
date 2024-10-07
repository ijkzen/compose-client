package composeclient

import "testing"

func TestList(t *testing.T) {
	containers, err := List()

	if err != nil {
		t.Fatal(err)
	} else {
		for _, container := range containers {
			t.Log(container.Names)
		}
	}
}
