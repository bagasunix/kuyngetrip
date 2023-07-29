package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Province struct {
	ID   string `gorm:"type:varchar(100);primaryKey;unique" json:"code"`
	Name string `gorm:"type:varchar(255)" json:"province"`
}

type City struct {
	ID         string `gorm:"type:varchar(100);primaryKey;unique" json:"code"`
	ProvinceID string `gorm:"type:varchar(100);index"`
	Name       string `json:"city" gorm:"type:varchar(255)"`
}

type District struct {
	ID     string `gorm:"type:varchar(100);unique;primaryKey" json:"code"`
	CityID string `gorm:"type:varchar(100);index"`
	Name   string `json:"district" gorm:"type:varchar(255)"`
}

type Village struct {
	ID         string `gorm:"type:varchar(100);unique;primaryKey" json:"code"`
	DistrictID string `gorm:"type:varchar(100);index"`
	Name       string `json:"village" gorm:"type:varchar(255)"`
	Postal     int64  `json:"postal"`
	Latitude   string `gorm:"type:varchar(255)"`
	Longitude  string `gorm:"type:varchar(255)"`
	Elevation  string `gorm:"type:varchar(255)"`
}

type JSProvince struct {
	ID     string `gorm:"primaryKey;unique" json:"code"`
	Name   string `json:"province"`
	Cities []JSCity
}

type JSCity struct {
	ID         string `gorm:"primaryKey;unique" json:"code"`
	ProvinceID string `gorm:"index"`
	Name       string `json:"city" `
	Districts  []JSDistrict
}

type JSDistrict struct {
	ID       string `gorm:"unique;primaryKey" json:"code"`
	CityID   string `gorm:"index"`
	Name     string `json:"district" `
	Villages []JSVillage
}

type JSVillage struct {
	ID         string `gorm:"unique;primaryKey" json:"code"`
	DistrictID string `gorm:"index"`
	Name       string `json:"village" `
	Postal     int64  `json:"postal"`
	Latitude   float64
	Longitude  float64
	Elevation  float64
}

func connect() (*gorm.DB, error) {
	host := "localhost"
	port := 5432
	user := "bagasunix"
	password := "Kambing04"
	dbname := "kuyngetrip"

	// membuat string koneksi (DSN)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// membuka koneksi ke database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db, nil
}

func DataMigration() {
	db, err := connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.AutoMigrate(&Province{}, &City{}, &District{}, &Village{})
	if err != nil {
		panic("failed to migrate database")
	}

	file, err := ioutil.ReadFile("./collection.json")
	if err != nil {
		log.Fatal(err)
	}

	// unmarshal json data
	var provinces []JSProvince
	err = json.Unmarshal(file, &provinces)
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Begin()

	// loop through provinces
	for _, province := range provinces {
		// insert province data into province table
		// ...

		mProv := Province{ID: province.ID, Name: province.Name}
		if err = db.Create(&mProv).Error; err != nil {
			db.Rollback()
			fmt.Println(err)
			return
		}
		// loop through cities
		for _, city := range province.Cities {
			// insert city data into city table
			// ...
			mCity := City{ID: strings.Replace(city.ID, ".", "", -1), ProvinceID: province.ID, Name: city.Name}

			if err = db.Create(&mCity).Error; err != nil {
				db.Rollback()
				fmt.Println(err)
				return
			}

			// loop through districts
			for _, district := range city.Districts {
				// insert district data into district table
				// ...
				mDist := District{ID: strings.Replace(district.ID, ".", "", -1), CityID: city.ID, Name: district.Name}

				if err = db.Create(&mDist).Error; err != nil {
					db.Rollback()
					fmt.Println(err)
					return
				}
				// loop through villages
				for _, village := range district.Villages {
					// insert village data into village table
					// ...
					mVill := Village{ID: strings.Replace(village.ID, ".", "", -1), DistrictID: strings.Replace(district.ID, ".", "", -1), Name: village.Name, Postal: village.Postal, Latitude: strconv.FormatFloat(village.Latitude, 'f', -1, 64), Longitude: strconv.FormatFloat(village.Longitude, 'f', -1, 64)}

					if err = db.Create(&mVill).Error; err != nil {
						db.Rollback()
						fmt.Println(err)
						return
					}
				}
			}
		}
	}
	db.Commit()
}
