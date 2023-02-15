package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	Name  string
	Price float64
}

var FoodModel Food

func (f *Food) Create(name string, price float64) int {
	newFood := Food{
		Name:  name,
		Price: price,
	}
	DB.Create(&newFood)
	return int(newFood.ID)
}

func (f *Food) FindOneById(id int) Food {
	food := Food{}
	DB.First(&food, id)
	return food
}

func (f *Food) FindByName(name string) []Food {
	var foods []Food
	DB.Where("name LIKE ?", "%"+name+"%").Find(&foods)
	return foods
}

func (f *Food) Update(updatedFood *Food) {
	DB.Save(*updatedFood)
}

func (f *Food) Updates(update Food, query interface{}, args ...interface{}) {
	DB.Table(f.GetTable()).Where(query, args...).Updates(update)
}

func (f *Food) String() string {
	return fmt.Sprintf("Foodname: %v", f.Name)
}

func (f *Food) GetTable() string {
	return "foods"
}
