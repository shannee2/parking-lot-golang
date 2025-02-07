package vehicle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVehicleCreation(t *testing.T) {
	r := "KA-01-HH-1234"
	v := New(r, Red)
	assert.NotNil(t, v, "Expected vehicle to be created")
}

func TestVehicleHasRegistrationNumber(t *testing.T) {
	r := "KA-01-HH-1234"
	v := New(r, Red)
	assert.True(t, v.HasRegistrationNumber(r), "Expected vehicle to have registration number KA-01-HH-1234")
	assert.False(t, v.HasRegistrationNumber("KA-01-HH-5678"), "Expected vehicle to not have registration number KA-01-HH-5678")
}

func TestVehicleHasColor(t *testing.T) {
	r := "KA-01-HH-1234"
	v := New(r, Red)
	assert.True(t, v.HasColor(Red), "Expected vehicle to have color Red")
	assert.False(t, v.HasColor(Blue), "Expected vehicle to not have color Blue")
}
