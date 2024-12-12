package models

import "time"

type Order struct {
	Id              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CarId           uint      `json:"car_id"`
	OrderDate       time.Time `gorm:"not null" json:"order_date"`
	PickupDate      time.Time `gorm:"not null" json:"pickup_date"`
	DropOffDate     time.Time `gorm:"not null" json:"dropoff_date"`
	PickupLocation  string    `gorm:"type:char(50);not null" json:"pickup_location"`
	DropOffLocation string    `gorm:"type:char(50);not null" json:"dropoff_location"`
	Car             Car       `gorm:"foreignKey:CarId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
