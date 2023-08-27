package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Tour struct {
	BaseModel
	TourName        string
	TourTypeID      uuid.UUID `gorm:"not null;type:uuid;"`
	TourType        TourType  `gorm:"foreignKey:TourTypeID;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	TourDescription string
	TourPrice       float64
	TourStartDate   time.Time
	TourEndDate     time.Time
	TourCapacity    int8
}

// Builder Object for Tour
type TourBuilder struct {
	BaseModelBuilder
	tourName        string
	tourTypeID      uuid.UUID
	tourType        TourType
	tourDescription string
	tourPrice       float64
	tourStartDate   time.Time
	tourEndDate     time.Time
	tourCapacity    int8
}

// Constructor for TourBuilder
func NewTourBuilder() *TourBuilder {
	o := new(TourBuilder)
	return o
}

// Build Method which creates Tour
func (b *TourBuilder) Build() *Tour {
	o := new(Tour)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.TourName = b.tourName
	o.TourTypeID = b.tourTypeID
	o.TourType = b.tourType
	o.TourDescription = b.tourDescription
	o.TourPrice = b.tourPrice
	o.TourStartDate = b.tourStartDate
	o.TourEndDate = b.tourEndDate
	o.TourCapacity = b.tourCapacity
	return o
}

// Setter method for the field tourName of type string in the object TourBuilder
func (t *TourBuilder) SetTourName(tourName string) {
	t.tourName = tourName
}

// Setter method for the field tourType of type TourType in the object TourBuilder
func (t *TourBuilder) SetTourType(tourType TourType) {
	t.tourType = tourType
	t.tourTypeID = tourType.ID
}

// Setter method for the field tourDescription of type string in the object TourBuilder
func (t *TourBuilder) SetTourDescription(tourDescription string) {
	t.tourDescription = tourDescription
}

// Setter method for the field tourPrice of type float64 in the object TourBuilder
func (t *TourBuilder) SetTourPrice(tourPrice float64) {
	t.tourPrice = tourPrice
}

// Setter method for the field tourStartDate of type time.Time in the object TourBuilder
func (t *TourBuilder) SetTourStartDate(tourStartDate time.Time) {
	t.tourStartDate = tourStartDate
}

// Setter method for the field tourEndDate of type time.Time in the object TourBuilder
func (t *TourBuilder) SetTourEndDate(tourEndDate time.Time) {
	t.tourEndDate = tourEndDate
}

// Setter method for the field tourCapacity of type int8 in the object TourBuilder
func (t *TourBuilder) SetTourCapacity(tourCapacity int8) {
	t.tourCapacity = tourCapacity
}
