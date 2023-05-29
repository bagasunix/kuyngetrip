package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type TourSchedule struct {
	BaseModel
	TourID      uuid.UUID
	TourDate    time.Time
	Capacity    int
	Tour        Tour          `gorm:"foreignKey:TourID"`
	Participant []Participant `gorm:"foreignKey:ParticipantID"`
}

// Builder Object for TourSchedule
type TourScheduleBuilder struct {
	BaseModelBuilder
	tourID      uuid.UUID
	tour        Tour
	capacity    int
	participant []Participant
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
	o.Participant = b.participant
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
func (t *TourScheduleBuilder) SetParticipant(participant []Participant) {
	t.participant = participant
}
