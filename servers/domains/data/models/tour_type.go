package models

type TourType struct {
	BaseModel
	TourTypeName        string
	TourTypeDescription string
}

// Builder Object for TourType
type TourTypeBuilder struct {
	BaseModelBuilder
	tourTypeName        string
	tourTypeDescription string
}

// Constructor for TourTypeBuilder
func NewTourTypeBuilder() *TourTypeBuilder {
	o := new(TourTypeBuilder)
	return o
}

// Build Method which creates TourType
func (b *TourTypeBuilder) Build() *TourType {
	o := new(TourType)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.TourTypeDescription = b.tourTypeDescription
	o.TourTypeName = b.tourTypeName
	return o
}

// Setter method for the field tourTypeName of type string in the object TourTypeBuilder
func (t *TourTypeBuilder) SetTourTypeName(tourTypeName string) {
	t.tourTypeName = tourTypeName
}

// Setter method for the field tourTypeDescription of type string in the object TourTypeBuilder
func (t *TourTypeBuilder) SetTourTypeDescription(tourTypeDescription string) {
	t.tourTypeDescription = tourTypeDescription
}
