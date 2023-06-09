package models

import (
	"github.com/gofrs/uuid"
)

type Tour struct {
	BaseModel
	TourGuideID     uuid.UUID `gorm:"not null;"`
	User            User      `gorm:"foreignKey:TourGuideID;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	TourName        string    `gorm:"not null;"`
	TourDescription string
	TourPrice       float64        `gorm:"not null;"`
	Reviews         []TourReview   `gorm:"foreignKey:id"`
	TourSchedule    []TourSchedule `gorm:"foreignKey:id"`
	Destinations    []Destination  `gorm:"foreignKey:id"`
	Deleted
}

// Builder Object for Tour
type TourBuilder struct {
	BaseModelBuilder
	tourGuideID     uuid.UUID
	user            User
	tourName        string
	tourDescription string
	tourPrice       float64
	review          []TourReview
	tourschedule    []TourSchedule
	DeletedBuilder
	destinations []Destination
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
	o.TourGuideID = b.tourGuideID
	o.User = b.user
	o.TourName = b.tourName
	o.TourDescription = b.tourDescription
	o.TourPrice = b.tourPrice
	o.Reviews = b.review
	o.TourSchedule = b.tourschedule
	o.Deleted = *b.DeletedBuilder.Build()
	o.Destinations = b.destinations
	return o
}

// Setter method for the field tourGuideID of type uuid.UUID in the object TourBuilder
func (t *TourBuilder) SetTourGuide(tourGuideID uuid.UUID) {
	t.tourGuideID = tourGuideID
	t.user.ID = tourGuideID
}

// Setter method for the field tourName of type string in the object TourBuilder
func (t *TourBuilder) SetTourName(tourName string) {
	t.tourName = tourName
}

// Setter method for the field tourDescription of type string in the object TourBuilder
func (t *TourBuilder) SetTourDescription(tourDescription string) {
	t.tourDescription = tourDescription
}

// Setter method for the field tourPrice of type float64 in the object TourBuilder
func (t *TourBuilder) SetTourPrice(tourPrice float64) {
	t.tourPrice = tourPrice
}

// Setter method for the field review of type []TourReview in the object TourBuilder
func (t *TourBuilder) SetReview(review []TourReview) {
	t.review = review
}

// Setter method for the field tourschedule of type []TourSchedule in the object TourBuilder
func (t *TourBuilder) SetTourschedule(tourschedule []TourSchedule) {
	t.tourschedule = tourschedule
}

// Setter method for the field destinations of type []Destination in the object TourBuilder
func (t *TourBuilder) SetDestinations(destinations []Destination) {
	t.destinations = destinations
}
