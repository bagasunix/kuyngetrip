package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Company struct {
	BaseModel
	CompanyName            string `gorm:"column:company_name"`
	CompanyPhone           string `gorm:"column:company_phone;size:13"`
	CompanyEmail           string `gorm:"column:company_email"`
	CompanyUserPICId       *uuid.UUID
	CompanyUserPIC         *User  `gorm:"column:company_pic_user_id;foreignKey:CompanyUserPICId"`
	CompanyAddress         string `gorm:"column:company_address"`
	CompanyVillageId       int64
	CompanyVillage         Village `gorm:"column:company_village;foreignKey:CompanyVillageId"`
	CompanySubDistrictId   int32
	CompanySubDistrict     SubDistrict `gorm:"column:company_subDistrict;foreignKey:CompanySubDistrictId"`
	CompanyCityId          int16
	CompanyCity            City `gorm:"column:company_city;foreignKey:CompanyCityId"`
	CompanyProvinceId      int16
	CompanyProvince        Province `gorm:"column:company_province;foreignKey:CompanyProvinceId"`
	CompanyPostalCode      string   `gorm:"column:company_postal_code"`
	CompanyOwnerTitle      string   `gorm:"column:company_owner_title"`
	CompanyOwnerUserId     uuid.UUID
	CompanyOwnerUser       User       `gorm:"column:company_owner_user_id;foreignKey:CompanyOwnerUserId"`
	CompanyOwnerPostalCode string     `gorm:"column:company_owner_postal_code"`
	TokenVerifyEmail       string     `gorm:"column:token_verify_email"`
	VerifiedAT             *time.Time `gorm:"column:verified_at"`
}

// Builder Object for Company
type CompanyBuilder struct {
	BaseModelBuilder
	companyName            string
	companyPhone           string
	companyEmail           string
	companyUserPICId       *uuid.UUID
	companyUserPIC         *User
	companyAddress         string
	companyVillageId       int64
	companyVillage         Village
	companySubDistrictId   int32
	companySubDistrict     SubDistrict
	companyCityId          int16
	companyCity            City
	companyProvinceId      int16
	companyProvince        Province
	companyPostalCode      string
	companyOwnerTitle      string
	companyOwnerUserId     uuid.UUID
	companyOwnerUser       User
	companyOwnerPostalCode string
	tokenVerifyEmail       string
	verifiedAT             *time.Time
}

// Constructor for CompanyBuilder
func NewCompanyBuilder() *CompanyBuilder {
	o := new(CompanyBuilder)
	return o
}

// Build Method which creates Company
func (b *CompanyBuilder) Build() *Company {
	o := new(Company)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.CompanyName = b.companyName
	o.CompanyPhone = b.companyPhone
	o.CompanyEmail = b.companyEmail
	o.CompanyUserPICId = b.companyUserPICId
	o.CompanyUserPIC = b.companyUserPIC
	o.CompanyAddress = b.companyAddress
	o.CompanyVillage = b.companyVillage
	o.CompanyVillageId = b.companyVillageId
	o.CompanySubDistrict = b.companySubDistrict
	o.CompanySubDistrictId = b.companySubDistrictId
	o.CompanyCity = b.companyCity
	o.CompanyCityId = b.companyCityId
	o.CompanyProvince = b.companyProvince
	o.CompanyProvinceId = b.companyProvinceId
	o.CompanyPostalCode = b.companyPostalCode
	o.CompanyOwnerTitle = b.companyOwnerTitle
	o.CompanyOwnerUser = b.companyOwnerUser
	o.CompanyOwnerUserId = b.companyOwnerUserId
	o.CompanyOwnerPostalCode = b.companyOwnerPostalCode
	o.TokenVerifyEmail = b.tokenVerifyEmail
	o.VerifiedAT = b.verifiedAT
	return o
}

// Setter method for the field companyName of type string in the object CompanyBuilder
func (c *CompanyBuilder) SetCompanyName(companyName string) {
	c.companyName = companyName
}

// Setter method for the field companyPhone of type string in the object CompanyBuilder
func (c *CompanyBuilder) SetCompanyPhone(companyPhone string) {
	c.companyPhone = companyPhone
}

// Setter method for the field companyEmail of type string in the object CompanyBuilder
func (c *CompanyBuilder) SetCompanyEmail(companyEmail string) {
	c.companyEmail = companyEmail
}

// Setter method for the field companyUserPIC of type User in the object CompanyBuilder
func (c *CompanyBuilder) SetCompanyUserPIC(companyUserPIC *User) {
	c.companyUserPIC = companyUserPIC
	c.companyUserPICId = &companyUserPIC.ID
}

// Setter method for the field companyAddress of type string in the object CompanyBuilder
func (c *CompanyBuilder) SetCompanyAddress(companyAddress string) {
	c.companyAddress = companyAddress
}

// Setter method for the field companyVillage of type Village in the object CompanyBuilder
func (c *CompanyBuilder) SetCompanyVillage(companyVillage Village) {
	c.companyVillage = companyVillage
	c.companyVillageId = companyVillage.Id
}

// Setter method for the field companySubDistrictId of type int32 in the object CompanyBuilder
func (c *CompanyBuilder) SetCompanySubDistrict(companySubDistrict SubDistrict) {
	c.companySubDistrict = companySubDistrict
	c.companySubDistrictId = companySubDistrict.Id
}

// Setter method for the field companyCity of type City in the object CompanyBuilder
func (c *CompanyBuilder) SetCompanyCity(companyCity City) {
	c.companyCity = companyCity
	c.companyCityId = companyCity.Id
}

// Setter method for the field companyProvince of type Province in the object CompanyBuilder
func (c *CompanyBuilder) SetCompanyProvince(companyProvince Province) {
	c.companyProvince = companyProvince
	c.companyProvinceId = companyProvince.Id
}

// Setter method for the field companyPostalCode of type string in the object CompanyBuilder
func (c *CompanyBuilder) SetCompanyPostalCode(companyPostalCode string) {
	c.companyPostalCode = companyPostalCode
}

// Setter method for the field companyOwnerTitle of type string in the object CompanyBuilder
func (c *CompanyBuilder) SetCompanyOwnerTitle(companyOwnerTitle string) {
	c.companyOwnerTitle = companyOwnerTitle
}

// Setter method for the field companyOwnerUser of type User in the object CompanyBuilder
func (c *CompanyBuilder) SetCompanyOwnerUser(companyOwnerUser User) {
	c.companyOwnerUser = companyOwnerUser
	c.companyOwnerUserId = companyOwnerUser.ID
}

// Setter method for the field companyOwnerPostalCode of type string in the object CompanyBuilder
func (c *CompanyBuilder) SetCompanyOwnerPostalCode(companyOwnerPostalCode string) {
	c.companyOwnerPostalCode = companyOwnerPostalCode
}

// Setter method for the field tokenVerifyEmail of type string in the object CompanyBuilder
func (c *CompanyBuilder) SetTokenVerifyEmail(tokenVerifyEmail string) {
	c.tokenVerifyEmail = tokenVerifyEmail
}

// Setter method for the field verifiedAT of type *time.Time in the object CompanyBuilder
func (c *CompanyBuilder) SetVerifiedAT(verifiedAT *time.Time) {
	c.verifiedAT = verifiedAT
}
