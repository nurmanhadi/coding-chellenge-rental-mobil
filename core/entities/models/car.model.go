package models

type Car struct {
	Id    uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string  `gorm:"char(50);not null" json:"car_name"`
	Rate  float32 `gorm:"not null" json:"month_rate"`
	Image string  `gorm:"type:char(250);not null" json:"image"`
}