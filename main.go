package main

import (
	"./sites"
	"./db"
	"./model"
)

func main() {
	abehiroshiData := sites.Abehiroshi()
	apperances := model.ConvertApperanceFromSiteApperance(abehiroshiData)
	db := db.Init()
	defer db.Close()

	for _, apperance := range apperances {
		db.AutoMigrate(&model.Appearance{})
		db.Where(
			model.Appearance{
				Name: apperance.Name,
				AppearanceType: apperance.AppearanceType,
				Date: apperance.Date,
				Description: apperance.Description }).FirstOrCreate(&model.Appearance{})
	}
}
