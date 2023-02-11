package race

import "testing"

func TestGetCounter(t *testing.T) {
	counter := GetCounter()
	if counter != 5000 {
		t.Error("unexpected counter:", counter)
	}
}
