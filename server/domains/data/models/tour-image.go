package models

import "github.com/gofrs/uuid"

type TourImage struct {
	BaseModel
	DestinationID    uuid.UUID    `gorm:"not null;type:uuid;"`
	Destination      *Destination `gorm:"foreignKey:DestinationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ImageName        string
	ImagePath        string
	ImageDescription string
}

// Builder Object for TourImage
type TourImageBuilder struct {
	BaseModelBuilder
	destinationID    uuid.UUID
	destination      *Destination
	imageName        string
	imagePath        string
	imageDescription string
}

// Constructor for TourImageBuilder
func NewTourImageBuilder() *TourImageBuilder {
	o := new(TourImageBuilder)
	return o
}

// Build Method which creates TourImage
func (b *TourImageBuilder) Build() *TourImage {
	o := new(TourImage)
	o.Destination = b.destination
	o.DestinationID = b.destinationID
	o.ImageName = b.imageName
	o.ImagePath = b.imagePath
	o.ImageDescription = b.imageDescription
	return o
}

// Setter method for the field tour of type *Tour in the object TourImageBuilder
func (t *TourImageBuilder) SetTour(destination *Destination) {
	t.destination = destination
	t.destinationID = destination.ID
}

// Setter method for the field imageName of type string in the object TourImageBuilder
func (t *TourImageBuilder) SetImageName(imageName string) {
	t.imageName = imageName
}

// Setter method for the field imagePath of type string in the object TourImageBuilder
func (t *TourImageBuilder) SetImagePath(imagePath string) {
	t.imagePath = imagePath
}

// Setter method for the field imageDescription of type string in the object TourImageBuilder
func (t *TourImageBuilder) SetImageDescription(imageDescription string) {
	t.imageDescription = imageDescription
}
