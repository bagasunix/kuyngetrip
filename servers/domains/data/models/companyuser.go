package models

import "github.com/gofrs/uuid"

type CompanyUser struct {
	BaseModel
	UserID    uuid.UUID
	User      User `gorm:"column:user_id;foreignKey:UserID"`
	CompanyID uuid.UUID
	Company   Company `gorm:"column:user_id;foreignKey:CompanyID"`
}

// Builder Object for CompanyUser
type CompanyUserBuilder struct {
	BaseModelBuilder
	userID    uuid.UUID
	user      User
	companyID uuid.UUID
	company   Company
}

// Constructor for CompanyUserBuilder
func NewCompanyUserBuilder() *CompanyUserBuilder {
	o := new(CompanyUserBuilder)
	return o
}

// Build Method which creates CompanyUser
func (b *CompanyUserBuilder) Build() *CompanyUser {
	o := new(CompanyUser)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.User = b.user
	o.UserID = b.userID
	o.Company = b.company
	o.CompanyID = b.companyID
	return o
}

// Setter method for the field user of type User in the object CompanyUserBuilder
func (c *CompanyUserBuilder) SetUser(user User) {
	c.user = user
	c.userID = user.ID
}

// Setter method for the field company of type Company in the object CompanyUserBuilder
func (c *CompanyUserBuilder) SetCompany(company Company) {
	c.company = company
	c.companyID = company.ID
}
