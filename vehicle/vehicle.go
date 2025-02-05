package vehicle

type Vehicle struct {
	RegistrationNumber string
	Color              VehicleColor
}

func (v Vehicle) HasRegistrationNumber(registrationNumber string) bool {
	return v.RegistrationNumber == registrationNumber
}

func (v Vehicle) HasColor(color VehicleColor) bool {
	return v.Color == color
}

func NewVehicle(registrationNumber string, color VehicleColor) *Vehicle {
	return &Vehicle{
		RegistrationNumber: registrationNumber,
		Color:              color,
	}
}
