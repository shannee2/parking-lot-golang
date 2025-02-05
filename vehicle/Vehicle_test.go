package vehicle

import (
	"testing"
)

func TestVehicle(t *testing.T) {
	v := Vehicle{"KA-01-HH-1234", Red}
	if !v.HasRegistrationNumber("KA-01-HH-1234") {
		t.Errorf("Expected registration number to be KA-01-HH-1234, but got %s", v.RegistrationNumber)
	}
	if !v.HasColor(Red) {
		t.Errorf("Expected color to be Red, but got %v", v.Color)
	}
}
