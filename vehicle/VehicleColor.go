package vehicle

type VehicleColor int

// Declare related constants for each weekday starting with index 1
const (
	Red   VehicleColor = iota + 1 // EnumIndex = 1
	Green                         // EnumIndex = 2
	Blue                          // EnumIndex = 3
)
