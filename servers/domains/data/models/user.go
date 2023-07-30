package models

import "time"

type User struct {
	BaseModel
	UserName              string `gorm:"size:100;uniqueIndex:idx_user_name_unique"`
	Email                 string `gorm:"uniqueIndex:idx_user_email_unique"`
	Password              string
	UserType              string `gorm:"size:1;comment:1 - Administrator, 2 - Admin, 3 - Guide, 4 - Participant"`
	Picture               *string
	Active                string `gorm:"size:1;comment:0 - NonActive, 1 - Active, 2 - Pending, 3 - Suspended"`
	LastLogin             *time.Time
	IpLogin               *string
	BrowserLogin          *string
	IpChangePassword      *string
	BrowserChangePassword *string
	LastChangePassword    *time.Time
	UserDetail            UserDetail `gorm:"foreignKey:UserID"`
}

// Builder Object for User
type UserBuilder struct {
	BaseModelBuilder
	userName              string
	email                 string
	password              string
	userType              string
	picture               *string
	activite              string
	lastLogin             *time.Time
	ipLogin               *string
	browserLogin          *string
	ipChangePassword      *string
	browserChangePassword *string
	lastChangePassword    *time.Time
	userDetail            UserDetail
}

// Constructor for UserBuilder
func NewUserBuilder() *UserBuilder {
	o := new(UserBuilder)
	return o
}

// Build Method which creates User
func (b *UserBuilder) Build() *User {
	o := new(User)
	o.UserName = b.userName
	o.Email = b.email
	o.Password = b.password
	o.UserType = b.userType
	o.Picture = b.picture
	o.Active = b.activite
	o.LastLogin = b.lastLogin
	o.IpLogin = b.ipLogin
	o.BrowserLogin = b.browserLogin
	o.IpChangePassword = b.ipChangePassword
	o.BrowserChangePassword = b.browserChangePassword
	o.LastChangePassword = b.lastChangePassword
	o.UserDetail = b.userDetail
	return o
}

// Setter method for the field userName of type string in the object UserBuilder
func (u *UserBuilder) SetUserName(userName string) {
	u.userName = userName
}

// Setter method for the field email of type string in the object UserBuilder
func (u *UserBuilder) SetEmail(email string) {
	u.email = email
}

// Setter method for the field password of type string in the object UserBuilder
func (u *UserBuilder) SetPassword(password string) {
	u.password = password
}

// Setter method for the field picture of type string in the object UserBuilder
func (u *UserBuilder) SetPicture(picture *string) {
	u.picture = picture
}

// Setter method for the field userDetail of type userDetail in the object UserBuilder
func (u *UserBuilder) SetUserDetail(userDetail UserDetail) {
	u.userDetail = userDetail
}

// Setter method for the field userType of type string in the object UserBuilder
func (u *UserBuilder) SetUserType(userType string) {
	u.userType = userType
}

// Setter method for the field activite of type string in the object UserBuilder
func (u *UserBuilder) SetActivite(activite string) {
	u.activite = activite
}

// Setter method for the field lastLogin of type *time.Time in the object UserBuilder
func (u *UserBuilder) SetLastLogin(lastLogin *time.Time) {
	u.lastLogin = lastLogin
}

// Setter method for the field ipLogin of type *string in the object UserBuilder
func (u *UserBuilder) SetIpLogin(ipLogin *string) {
	u.ipLogin = ipLogin
}

// Setter method for the field browserLogin of type *string in the object UserBuilder
func (u *UserBuilder) SetBrowserLogin(browserLogin *string) {
	u.browserLogin = browserLogin
}

// Setter method for the field ipChangePassword of type *string in the object UserBuilder
func (u *UserBuilder) SetIpChangePassword(ipChangePassword *string) {
	u.ipChangePassword = ipChangePassword
}

// Setter method for the field lastChangePassword of type *time.Time in the object UserBuilder
func (u *UserBuilder) SetLastChangePassword(lastChangePassword *time.Time) {
	u.lastChangePassword = lastChangePassword
}

// Setter method for the field browserChangePassword of type *string in the object UserBuilder
func (u *UserBuilder) SetBrowserChangePassword(browserChangePassword *string) {
	u.browserChangePassword = browserChangePassword
}
