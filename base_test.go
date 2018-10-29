package elevator

import (
	"testing"
)

func TestRequest(t *testing.T) {
	// I don't know how to test this?
	err := CheckAndRequestElevation()
	if err != nil {
		t.Errorf("Error returned : " + err.Error())
	}
}
