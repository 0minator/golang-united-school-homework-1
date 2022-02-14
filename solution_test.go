package solution

import "testing"

func TestGetMessage(t *testing.T) {
	var expected = "Hello 🗺️ !"
	if got := GetMessage(); got != expected {
		t.Errorf("GetMessage() = %v, want %v", got, expected)
	}
}
