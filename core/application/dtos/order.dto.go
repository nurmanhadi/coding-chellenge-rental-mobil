package dtos

type AddRequestOrder struct {
	CarId           uint   `json:"car_id" validate:"required"`
	OrderDate       string `json:"order_date" validate:"required"`
	PickupDate      string `json:"pickup_date" validate:"required"`
	DropOffDate     string `json:"dropoff_date" validate:"required"`
	PickupLocation  string `json:"pickup_location" validate:"required,min=1,max=50"`
	DropOffLocation string `json:"dropoff_location" validate:"required,min=1,max50"`
}
type UpdateRequestOrder struct {
	OrderDate       *string `json:"order_date" validate:"omitempty"`
	PickupDate      *string `json:"pickup_date" validate:"omitempty"`
	DropOffDate     *string `json:"dropoff_date" validate:"omitempty"`
	PickupLocation  *string `json:"pickup_location" validate:"omitempty,min=1,max=50"`
	DropOffLocation *string `json:"dropoff_location" validate:"omitempty,min=1,max50"`
}
