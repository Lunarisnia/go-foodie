package models

import "gorm.io/gorm"

// TODO: Add CRUD for this

// Model for addon, relation is many-to-one with food
// Reason being in that the same name doesn't necessarily mean the same thing
// Like Pickle add on in a burger is not the same as the one in a fried rice
type AddOn struct {
	gorm.Model
	ID       uint
	Name     string
	Price    string
	PhotoUrl string
	InStock  bool
	FoodID   uint
}
