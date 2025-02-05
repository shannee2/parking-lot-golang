package vehicle

type Vehicle struct {
	RegistrationNumber string
	Color              VehicleColor
}

func (v Vehicle) HasRegistrationNumber(registrationNumber string) bool {
	return v.RegistrationNumber == registrationNumber
}

func (v *Vehicle) HasColor(color VehicleColor) bool {
	return v.Color == color
}
