package vehicle

import (
	"testing"
)

func TestVehicleCreation(t *testing.T) {
	registrationNumber := "KA-01-HH-1234"
	v := Vehicle{registrationNumber, Red}
	if !v.HasRegistrationNumber(registrationNumber) {
		t.Errorf("Expected registration number to be KA-01-HH-1234, but got %s", v.registrationNumber)
	}
	if !v.HasColor(Red) {
		t.Errorf("Expected color to be Red, but got %v", v.color)
	}
}

// More tests
