package dtos

type UpdateRequestCar struct {
	Name      *string  `json:"car_name" validate:"omitempty,min=1,max=50"`
	DayRate   *float32 `json:"day_rate" validate:"omitempty"`
	MonthRate *float32 `json:"month_rate" validate:"omitempty"`
}
