package structs

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Name          string
	BarrierID     string
	BarrierSecret string
}
