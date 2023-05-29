package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Payment struct {
	BaseModel
	TourScheduleID uuid.UUID     `gorm:"not null;type:uuid;"`
	TourSchedule   *TourSchedule `gorm:"foreignKey:TourScheduleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ParticipantID  uuid.UUID     `gorm:"not null;type:uuid;"`
	Participant    *Participant  `gorm:"foreignKey:ParticipantID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PaymentDate    time.Time
	PaymentStatus  string  `gorm:"size:10"`
	PaymentAmount  float64 `gorm:"type:decimal(10,2)"`
	PaymentMethod  string  `gorm:"size:50"`
}

// Builder Object for Payment
type PaymentBuilder struct {
	BaseModelBuilder
	tourScheduleID uuid.UUID
	tourSchedule   *TourSchedule
	participantID  uuid.UUID
	participant    *Participant
	paymentDate    time.Time
	payamentStatus string
	paymentAmount  float64
	paymentMethod  string
}

// Constructor for PaymentBuilder
func NewPaymentBuilder() *PaymentBuilder {
	o := new(PaymentBuilder)
	return o
}

// Build Method which creates Payment
func (b *PaymentBuilder) Build() *Payment {
	o := new(Payment)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.TourSchedule = b.tourSchedule
	o.TourScheduleID = b.tourScheduleID
	o.ParticipantID = b.participantID
	o.Participant = b.participant
	o.PaymentDate = b.paymentDate
	o.PaymentStatus = b.payamentStatus
	o.PaymentAmount = b.paymentAmount
	o.PaymentMethod = b.paymentMethod
	return o
}

// Setter method for the field tour of type *Tour in the object PaymentBuilder
func (t *PaymentBuilder) SetTour(tourSchedule *TourSchedule) {
	t.tourSchedule = tourSchedule
	t.tourScheduleID = tourSchedule.ID
}

// Setter method for the field participant of type *Participant in the object PaymentBuilder
func (t *PaymentBuilder) SetParticipant(participant *Participant) {
	t.participant = participant
	t.participantID = participant.ID
}

// Setter method for the field paymentDate of type time.Time in the object PaymentBuilder
func (p *PaymentBuilder) SetPaymentDate(paymentDate time.Time) {
	p.paymentDate = paymentDate
}

// Setter method for the field payamentStatus of type string in the object PaymentBuilder
func (p *PaymentBuilder) SetPayamentStatus(payamentStatus string) {
	p.payamentStatus = payamentStatus
}

// Setter method for the field paymentAmount of type string in the object PaymentBuilder
func (p *PaymentBuilder) SetPaymentAmount(paymentAmount float64) {
	p.paymentAmount = paymentAmount
}

// Setter method for the field paymentMethod of type string in the object PaymentBuilder
func (p *PaymentBuilder) SetPaymentMethod(paymentMethod string) {
	p.paymentMethod = paymentMethod
}
