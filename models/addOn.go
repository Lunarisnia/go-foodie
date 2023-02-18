package models

import (
	"fmt"

	"gorm.io/gorm"
)

// TODO: Add CRUD for this

// Model for addon, relation is many-to-one with food
// Reason being in that the same name doesn't necessarily mean the same thing
// Like Pickle add on in a burger is not the same as the one in a fried rice
type AddOn struct {
	gorm.Model
	ID       uint
	Name     string
	Price    float64
	PhotoUrl string
	InStock  bool `gorm:"default:true"`
	FoodID   uint
}

var AddOnModel AddOn

func (f *AddOn) Create(name, photoUrl string, price float64, inStock bool, foodId uint) int {
	newData := AddOn{
		Name:     name,
		Price:    price,
		PhotoUrl: photoUrl,
		InStock:  inStock,
		FoodID:   foodId,
	}
	DB.Create(&newData)
	return int(newData.ID)
}

func (f *AddOn) FindOneById(id string) AddOn {
	food := AddOn{}
	DB.Model(&AddOn{}).First(&food, id)
	return food
}

func (f *AddOn) FindByName(name string) []AddOn {
	var foods []AddOn
	DB.Model(&AddOn{}).Where("name LIKE ?", "%"+name+"%").Find(&foods)
	return foods
}

func (f *AddOn) Update(updatedFood *AddOn) {
	DB.Save(*updatedFood)
}

func (f *AddOn) Updates(update AddOn, query interface{}, args ...interface{}) {
	DB.Table(f.getTable()).Where(query, args...).Updates(update)
}

func (f *AddOn) Delete(addOn AddOn) {
	DB.Delete(&addOn)
}

func (f *AddOn) String() string {
	return fmt.Sprintf("Foodname: %v", f.Name)
}

func (f *AddOn) getTable() string {
	return "add_ons"
}
