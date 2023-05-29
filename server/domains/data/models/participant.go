package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Participant struct {
	BaseModel
	TourScheduleID   uuid.UUID     `gorm:"not null;type:uuid;"`
	TourSchedule     *TourSchedule `gorm:"foreignKey:TourScheduleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserID           uuid.UUID     `gorm:"not null;type:uuid;"`
	User             *User         `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	RegistrationDate time.Time
	Payment          []Payment    `gorm:"foreignKey:ParticipantID"`
	TourReview       []TourReview `gorm:"foreignKey:ParticipantID"`
}

// Builder Object for Participant
type ParticipantBuilder struct {
	BaseModelBuilder
	tourScheduleID   uuid.UUID
	tourSchedule     *TourSchedule
	userID           uuid.UUID
	user             *User
	registrationDate time.Time
	payment          []Payment
	tourReview       []TourReview
}

// Constructor for ParticipantBuilder
func NewParticipantBuilder() *ParticipantBuilder {
	o := new(ParticipantBuilder)
	return o
}

// Build Method which creates Participant
func (b *ParticipantBuilder) Build() *Participant {
	o := new(Participant)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.TourScheduleID = b.tourScheduleID
	o.TourSchedule = b.tourSchedule
	o.UserID = b.userID
	o.User = b.user
	o.RegistrationDate = b.registrationDate
	o.Payment = b.payment
	o.TourReview = b.tourReview
	return o
}

// Setter method for the field tour of type *Tour in the object ParticipantBuilder
func (p *ParticipantBuilder) SetTour(tourSchedule *TourSchedule) {
	p.tourSchedule = tourSchedule
	p.tourScheduleID = tourSchedule.ID
}

// Setter method for the field user of type *User in the object ParticipantBuilder
func (p *ParticipantBuilder) SetUser(user *User) {
	p.user = user
	p.userID = user.ID
}

// Setter method for the field registrationDate of type time.Time in the object ParticipantBuilder
func (p *ParticipantBuilder) SetRegistrationDate(registrationDate time.Time) {
	p.registrationDate = registrationDate
}

// Setter method for the field payment of type []Payment in the object ParticipantBuilder
func (p *ParticipantBuilder) SetPayment(payment []Payment) {
	p.payment = payment
}

// Setter method for the field tourReview of type []TourReview in the object ParticipantBuilder
func (p *ParticipantBuilder) SetTourReview(tourReview []TourReview) {
	p.tourReview = tourReview
}
