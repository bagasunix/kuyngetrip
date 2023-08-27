package models

import (
	"github.com/gofrs/uuid"
)

type TourDestination struct {
	BaseModel
	TourID                 uuid.UUID `gorm:"not null;type:uuid;"`
	Tour                   Tour      `gorm:"foreignKey:TourID;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	DestinationName        string
	DestinationDescription string
	VillageId              int64
	Village                Village `gorm:"foreignKey:VillageId"`
	SubDistrictId          int32
	SubDistrict            SubDistrict `gorm:"foreignKey:SubDistrictId"`
	CityId                 int16
	City                   City `gorm:"foreignKey:CityId"`
	ProvinceId             int16
	Province               Province `gorm:"foreignKey:ProvinceId"`
}

// Builder Object for TourDestination
type TourDestinationBuilder struct {
	BaseModelBuilder
	tourID                 uuid.UUID
	tour                   Tour
	destinationName        string
	destinationDescription string
	villageId              int64
	village                Village
	subDistrictId          int32
	subDistrict            SubDistrict
	cityId                 int16
	city                   City
	provinceId             int16
	province               Province
}

// Constructor for TourDestinationBuilder
func NewTourDestinationBuilder() *TourDestinationBuilder {
	o := new(TourDestinationBuilder)
	return o
}

// Build Method which creates TourDestination
func (b *TourDestinationBuilder) Build() *TourDestination {
	o := new(TourDestination)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.TourID = b.tourID
	o.Tour = b.tour
	o.DestinationName = b.destinationName
	o.DestinationDescription = b.destinationDescription
	o.VillageId = b.villageId
	o.Village = b.village
	o.SubDistrict = b.subDistrict
	o.SubDistrictId = b.subDistrictId
	o.City = b.city
	o.CityId = b.cityId
	o.Province = b.province
	o.ProvinceId = b.provinceId
	return o
}

// Setter method for the field tour of type Tour in the object TourDestinationBuilder
func (t *TourDestinationBuilder) SetTour(tour Tour) {
	t.tour = tour
	t.tourID = tour.ID
}

// Setter method for the field village of type Village in the object TourDestinationBuilder
func (u *TourDestinationBuilder) SetVillage(village Village) {
	u.village = village
	u.villageId = village.Id
}

// Setter method for the field subDistrict of type SubDistrict in the object TourDestinationBuilder
func (u *TourDestinationBuilder) SetSubDistrict(subDistrict SubDistrict) {
	u.subDistrict = subDistrict
	u.subDistrictId = subDistrict.Id
}

// Setter method for the field city of type City in the object TourDestinationBuilder
func (u *TourDestinationBuilder) SetCity(city City) {
	u.city = city
	u.cityId = city.Id
}

// Setter method for the field province of type Province in the object TourDestinationBuilder
func (u *TourDestinationBuilder) SetProvince(province Province) {
	u.province = province
	u.provinceId = province.Id
}

// Setter method for the field destinationName of type string in the object TourDestinationBuilder
func (t *TourDestinationBuilder) SetDestinationName(destinationName string) {
	t.destinationName = destinationName
}

// Setter method for the field destinationDescription of type string in the object TourDestinationBuilder
func (t *TourDestinationBuilder) SetDestinationDescription(destinationDescription string) {
	t.destinationDescription = destinationDescription
}
