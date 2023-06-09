package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type TourSchedule struct {
	BaseModel
	TourID       uuid.UUID `gorm:"not null;"`
	Tour         Tour      `gorm:"foreignKey:TourID"`
	TourDate     time.Time
	Capacity     int           `gorm:"not null;"`
	Participants []Participant `gorm:"foreignKey:TourScheduleID"`
	Deleted
}

// Builder Object for TourSchedule
type TourScheduleBuilder struct {
	BaseModelBuilder
	tourID       uuid.UUID
	tour         Tour
	capacity     int
	participants []Participant
	DeletedBuilder
}

// Constructor for TourScheduleBuilder
func NewTourScheduleBuilder() *TourScheduleBuilder {
	o := new(TourScheduleBuilder)
	return o
}

// Build Method which creates TourSchedule
func (b *TourScheduleBuilder) Build() *TourSchedule {
	o := new(TourSchedule)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.TourID = b.tourID
	o.Tour = b.tour
	o.Capacity = b.capacity
	o.Participants = b.participants
	o.Deleted = *b.DeletedBuilder.Build()
	return o
}

// Setter method for the field tour of type Tour in the object TourScheduleBuilder
func (t *TourScheduleBuilder) SetTour(tour Tour) {
	t.tour = tour
	t.tourID = tour.ID
}

// Setter method for the field capacity of type int in the object TourScheduleBuilder
func (t *TourScheduleBuilder) SetCapacity(capacity int) {
	t.capacity = capacity
}

// Setter method for the field participant of type []Participant in the object TourScheduleBuilder
func (t *TourScheduleBuilder) SetParticipant(participants []Participant) {
	t.participants = participants
}
