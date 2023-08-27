package models

import "github.com/gofrs/uuid"

type TourGuide struct {
	BaseModel
	TourID               uuid.UUID `gorm:"not null;type:uuid;"`
	Tour                 Tour      `gorm:"foreignKey:TourID;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	UserID               uuid.UUID `gorm:"not null;type:uuid;"`
	User                 User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	TourGuideDescription string
}

// Builder Object for TourGuide
type TourGuideBuilder struct {
	BaseModelBuilder
	tourID               uuid.UUID
	tour                 Tour
	userID               uuid.UUID
	user                 User
	tourGuideDescription string
}

// Constructor for TourGuideBuilder
func NewTourGuideBuilder() *TourGuideBuilder {
	o := new(TourGuideBuilder)
	return o
}

// Build Method which creates TourGuide
func (b *TourGuideBuilder) Build() *TourGuide {
	o := new(TourGuide)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.TourID = b.tourID
	o.Tour = b.tour
	o.UserID = b.userID
	o.User = b.user
	o.TourGuideDescription = b.tourGuideDescription
	return o
}

// Setter method for the field tour of type Tour in the object TourGuideBuilder
func (t *TourGuideBuilder) SetTour(tour Tour) {
	t.tour = tour
	t.tourID = tour.ID
}

// Setter method for the field user of type User in the object TourGuideBuilder
func (t *TourGuideBuilder) SetUser(user User) {
	t.user = user
	t.userID = user.ID
}

// Setter method for the field tourGuideDescription of type string in the object TourGuideBuilder
func (t *TourGuideBuilder) SetTourGuideDescription(tourGuideDescription string) {
	t.tourGuideDescription = tourGuideDescription
}
