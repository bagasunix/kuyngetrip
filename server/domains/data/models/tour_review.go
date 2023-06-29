package models

import "github.com/gofrs/uuid"

type TourReview struct {
	BaseModel
	TourScheduleID uuid.UUID    `gorm:"not null;"`
	TourSchedule   TourSchedule `gorm:"foreignKey:TourID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ParticipantID  uuid.UUID    `gorm:"not null;"`
	Participant    Participant  `gorm:"foreignKey:ParticipantID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Rating         float32      `gorm:"type:decimal(2,1)"`
	ReviewText     string
	Deleted
}

// Builder Object for TourReview
type TourReviewBuilder struct {
	BaseModelBuilder
	tourScheduleID uuid.UUID
	tourSchedule   TourSchedule
	participantID  uuid.UUID
	participant    Participant
	rating         float32
	reviewText     string
	DeletedBuilder
}

// Constructor for TourReviewBuilder
func NewTourReviewBuilder() *TourReviewBuilder {
	o := new(TourReviewBuilder)
	return o
}

// Build Method which creates TourReview
func (b *TourReviewBuilder) Build() *TourReview {
	o := new(TourReview)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.TourSchedule = b.tourSchedule
	o.TourScheduleID = b.tourScheduleID
	o.ParticipantID = b.participantID
	o.Participant = b.participant
	o.Rating = b.rating
	o.ReviewText = b.reviewText
	o.Deleted = *b.DeletedBuilder.Build()
	return o
}

// Setter method for the field tour of type *Tour in the object TourReviewBuilder
func (t *TourReviewBuilder) SetTour(tourScheduler TourSchedule) {
	t.tourSchedule = tourScheduler
	t.tourScheduleID = tourScheduler.ID
}

// Setter method for the field participant of type *Participant in the object TourReviewBuilder
func (t *TourReviewBuilder) SetParticipant(participant Participant) {
	t.participant = participant
	t.participantID = participant.ID
}

// Setter method for the field rating of type float32 in the object TourReviewBuilder
func (t *TourReviewBuilder) SetRating(rating float32) {
	t.rating = rating
}

// Setter method for the field reviewText of type string in the object TourReviewBuilder
func (t *TourReviewBuilder) SetReviewText(reviewText string) {
	t.reviewText = reviewText
}
