package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Deleted struct {
	Delete    *int       `gorm:"size:1"`
	DeletedBy *uuid.UUID `gorm:"type:uuid;index"`
	DeletedAt *time.Time `gorm:"autoUpdateTime;<-:update"`
}

// Builder Object for Deleted
type DeletedBuilder struct {
	delete    *int
	deletedBy *uuid.UUID
	deletedAt *time.Time
}

// Constructor for DeletedBuilder
func NewDeletedBuilder() *DeletedBuilder {
	o := new(DeletedBuilder)
	return o
}

// Build Method which creates Deleted
func (b *DeletedBuilder) Build() *Deleted {
	o := new(Deleted)
	o.Delete = b.delete
	o.DeletedBy = b.deletedBy
	o.DeletedAt = b.deletedAt
	return o
}

// Setter method for the field deteled of type *int in the object DeletedBuilder
func (d *DeletedBuilder) SetDeteled(delete *int) {
	d.delete = delete
}

// Setter method for the field deletedBy of type *uuid.UUID in the object DeletedBuilder
func (d *DeletedBuilder) SetDeletedBy(deletedBy *uuid.UUID) {
	d.deletedBy = deletedBy
}

// Setter method for the field deletedAt of type *time.Time in the object DeletedBuilder
func (d *DeletedBuilder) SetDeletedAt(deletedAt *time.Time) {
	d.deletedAt = deletedAt
}
