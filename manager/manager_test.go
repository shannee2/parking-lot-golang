package manager

import (
	"github.com/stretchr/testify/assert"
	//"parkinglot/errors"
	//"parkinglot/parkinglot"
	//"parkinglot/vehicle"
	"testing"
)

func TestManagerCreation(t *testing.T) {
	m := NewManager()
	assert.NotNil(t, m)
}
