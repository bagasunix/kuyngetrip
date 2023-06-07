package models

import "github.com/gofrs/uuid"

type Destination struct {
	BaseModel
	TourID                 uuid.UUID `gorm:"not null;type:uuid;"`
	Tour                   *Tour     `gorm:"foreignKey:TourID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	DestinationName        string    `gorm:"not null;"`
	DestinationDescription string
	Country
	Province
	City
	SubDistrict
	Village
	Coordinate
	ImageDestination []TourImage `gorm:"foreignKey:id"`
	Deleted
}

// Builder Object for Destination
type DestinationBuilder struct {
	BaseModelBuilder
	tourID uuid.UUID
	tour   *Tour
	CountryBuilder
	ProvinceBuilder
	CityBuilder
	SubDistrictBuilder
	VillageBuilder
	CoordinateBuilder
	destinationName        string
	destinationDescription string
	imageDestination       []TourImage
	DeletedBuilder
}

// Constructor for DestinationBuilder
func NewDestinationBuilder() *DestinationBuilder {
	o := new(DestinationBuilder)
	return o
}

// Build Method which creates Destination
func (b *DestinationBuilder) Build() *Destination {
	o := new(Destination)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.TourID = b.tourID
	o.Tour = b.tour
	o.Country = *b.CountryBuilder.Build()
	o.Province = *b.ProvinceBuilder.Build()
	o.City = *b.CityBuilder.Build()
	o.SubDistrict = *b.SubDistrictBuilder.Build()
	o.Village = *b.VillageBuilder.Build()
	o.DestinationName = b.destinationName
	o.DestinationDescription = b.destinationDescription
	o.ImageDestination = b.imageDestination
	o.Deleted = *b.DeletedBuilder.Build()
	return o
}

// Setter method for the field tour of type *Tour in the object DestinationBuilder
func (d *DestinationBuilder) SetTour(tour *Tour) {
	d.tour = tour
	d.tourID = tour.ID
}

// Setter method for the field destinationName of type string in the object DestinationBuilder
func (d *DestinationBuilder) SetDestinationName(destinationName string) {
	d.destinationName = destinationName
}

// Setter method for the field destinationDescription of type string in the object DestinationBuilder
func (d *DestinationBuilder) SetDestinationDescription(destinationDescription string) {
	d.destinationDescription = destinationDescription
}

// Setter method for the field imageDestination of type []TourImage in the object DestinationBuilder
func (d *DestinationBuilder) SetImageDestination(imageDestination []TourImage) {
	d.imageDestination = imageDestination
}
