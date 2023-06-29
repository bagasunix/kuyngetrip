package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type UserDetail struct {
	UserDetailID         uuid.UUID `gorm:"primaryKey;type:uuid;<-:create"`
	UserID               uuid.UUID `gorm:"not null;type:uuid;"`
	User                 *User     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	FullName             string    `gorm:"size:50;not null;index:,sort:asc,option:CONCURRENTLY"`
	DOB                  *time.Time
	POB                  *string
	Phone                string `gorm:"size:12;not null;uniqueIndex:idx_user_phone_unique"`
	Sex                  int    `gorm:"size:1;comment:1-male,2-female"`
	IdentificationType   *string
	IdentificationNumber *string
	IdentificationImage  *string
	VillageId            int64
	Village              Village `gorm:"foreignKey:VillageId"`
	SubDistrictId        int32
	SubDistrict          SubDistrict `gorm:"foreignKey:SubDistrictId"`
	CityId               int16
	City                 City `gorm:"foreignKey:CityId"`
	ProvinceId           int16
	Province             Province `gorm:"foreignKey:ProvinceId"`
	Verified             *int8    `gorm:"size:1"`
	Created
	Updated
}

// Builder Object for UserDetail
type UserDetailBuilder struct {
	userDetailID         uuid.UUID
	userID               uuid.UUID
	user                 *User
	fullName             string
	dOB                  *time.Time
	pOB                  *string
	phone                string
	sex                  int
	identificationType   *string
	identificationNumber *string
	identificationImage  *string
	villageId            int64
	village              Village
	subDistrictId        int32
	subDistrict          SubDistrict
	cityId               int16
	city                 City
	provinceId           int16
	province             Province
	verified             *int8
	CreatedBuilder
	UpdatedBuilder
}

// Constructor for UserDetailBuilder
func NewUserDetailBuilder() *UserDetailBuilder {
	o := new(UserDetailBuilder)
	return o
}

// Build Method which creates UserDetail
func (b *UserDetailBuilder) Build() *UserDetail {
	o := new(UserDetail)
	o.UserDetailID = b.userDetailID
	o.UserID = b.userID
	o.User = b.user
	o.FullName = b.fullName
	o.Phone = b.phone
	o.Sex = b.sex
	o.DOB = b.dOB
	o.POB = b.pOB
	o.IdentificationType = b.identificationType
	o.IdentificationImage = b.identificationImage
	o.IdentificationNumber = b.identificationNumber
	o.VillageId = b.villageId
	o.Village = b.village
	o.SubDistrict = b.subDistrict
	o.SubDistrictId = b.subDistrictId
	o.City = b.city
	o.CityId = b.cityId
	o.Province = b.province
	o.ProvinceId = b.provinceId
	o.Verified = b.verified
	o.Created = *b.CreatedBuilder.Build()
	o.Updated = *b.UpdatedBuilder.Build()
	return o
}

// Setter method for the field userDetailID of type uuid.UUID in the object UserDetailBuilder
func (i *UserDetailBuilder) SetuserDetailID(userDetailID uuid.UUID) {
	i.userDetailID = userDetailID
}

// Setter method for the field user of type *User in the object UserDetailBuilder
func (i *UserDetailBuilder) SetUser(user *User) {
	i.user.ID = i.userID
	i.user = user
}

// Setter method for the field fullName of type string in the object UserBuilder
func (u *UserDetailBuilder) SetFullName(fullName string) {
	u.fullName = fullName
}

// Setter method for the field identificationType of type string in the object UserDetailBuilder
func (i *UserDetailBuilder) SetIdentificationType(identificationType *string) {
	i.identificationType = identificationType
}

// Setter method for the field identificationNumber of type string in the object UserDetailBuilder
func (i *UserDetailBuilder) SetIdentificationNumber(identificationNumber *string) {
	i.identificationNumber = identificationNumber
}

// Setter method for the field identificationImage of type string in the object UserDetailBuilder
func (i *UserDetailBuilder) SetIdentificationImage(identificationImage *string) {
	i.identificationImage = identificationImage
}

// Setter method for the field dOB of type *time.Time in the object UserDetailBuilder
func (u *UserDetailBuilder) SetDOB(dOB *time.Time) {
	u.dOB = dOB
}

// Setter method for the field pOB of type *string in the object UserDetailBuilder
func (u *UserDetailBuilder) SetPOB(pOB *string) {
	u.pOB = pOB
}

// Setter method for the field phone of type string in the object UserDetailBuilder
func (u *UserDetailBuilder) SetPhone(phone string) {
	u.phone = phone
}

// Setter method for the field sex of type int in the object UserDetailBuilder
func (u *UserDetailBuilder) SetSex(sex int) {
	u.sex = sex
}

// Setter method for the field village of type Village in the object UserDetailBuilder
func (u *UserDetailBuilder) SetVillage(village Village) {
	u.village = village
	u.villageId = village.Id
}

// Setter method for the field subDistrict of type SubDistrict in the object UserDetailBuilder
func (u *UserDetailBuilder) SetSubDistrict(subDistrict SubDistrict) {
	u.subDistrict = subDistrict
	u.subDistrictId = subDistrict.Id
}

// Setter method for the field city of type City in the object UserDetailBuilder
func (u *UserDetailBuilder) SetCity(city City) {
	u.city = city
	u.cityId = city.Id
}

// Setter method for the field province of type Province in the object UserDetailBuilder
func (u *UserDetailBuilder) SetProvince(province Province) {
	u.province = province
	u.provinceId = province.Id
}

// Setter method for the field verified of type *int8 in the object UserDetailBuilder
func (u *UserDetailBuilder) SetVerified(verified *int8) {
	u.verified = verified
}
