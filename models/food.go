package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	ID          uint
	Name        string
	PhotoUrl    string
	Description string
	Rating      int
	BasePrice   float64
	AddOns      []AddOn
}

var FoodModel Food

func (f *Food) Create(name string, price float64) int {
	newFood := Food{
		Name:      name,
		BasePrice: price,
	}
	DB.Create(&newFood)
	return int(newFood.ID)
}

func (f *Food) FindOneById(id string) Food {
	food := Food{}
	DB.Model(&Food{}).Preload("AddOns").First(&food, id)
	return food
}

func (f *Food) FindByName(name string) []Food {
	var foods []Food
	DB.Model(&Food{}).Preload("AddOns").Where("name LIKE ?", "%"+name+"%").Find(&foods)
	return foods
}

func (f *Food) Update(updatedFood *Food) {
	DB.Save(*updatedFood)
}

func (f *Food) Updates(update Food, query interface{}, args ...interface{}) {
	DB.Table(f.getTable()).Where(query, args...).Updates(update)
}

func (f *Food) Delete(food Food) {
	DB.Delete(&food)
}

func (f *Food) String() string {
	return fmt.Sprintf("Foodname: %v", f.Name)
}

func (f *Food) getTable() string {
	return "foods"
}
