package vehicle

type Vehicle struct {
	registrationNumber string
	color              VehicleColor
}

func (v Vehicle) HasRegistrationNumber(registrationNumber string) bool {
	return v.registrationNumber == registrationNumber
}

func (v Vehicle) HasColor(color VehicleColor) bool {
	return v.color == color
}

func NewVehicle(registrationNumber string, color VehicleColor) *Vehicle {
	return &Vehicle{
		registrationNumber: registrationNumber,
		color:              color,
	}
}
