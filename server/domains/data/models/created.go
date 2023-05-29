package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Created struct {
	CreatedBy uuid.UUID `gorm:"type:uuid;index;<-:create"`
	CreatedAt time.Time `gorm:"type:uuid;index;sort:desc;<-:create"`
}

// Builder Object for Created
type CreatedBuilder struct {
	createBy uuid.UUID
	createAt time.Time
}

// Constructor for CreatedBuilder
func NewCreatedBuilder() *CreatedBuilder {
	o := new(CreatedBuilder)
	return o
}

// Build Method which creates Created
func (b *CreatedBuilder) Build() *Created {
	o := new(Created)
	o.CreatedBy = b.createBy
	o.CreatedAt = b.createAt
	return o
}

// Setter method for the field createBy of type uuid.UUID in the object CreatedBuilder
func (c *CreatedBuilder) SetCreateBy(createBy uuid.UUID) {
	c.createBy = createBy
}

// Setter method for the field createAt of type time.Time in the object CreatedBuilder
func (c *CreatedBuilder) SetCreateAt(createAt time.Time) {
	c.createAt = createAt
}
