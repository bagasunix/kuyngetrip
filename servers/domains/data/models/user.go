package models

type User struct {
	BaseModel
	UserName   string
	Email      string `gorm:"size:100;uniqueIndex:idx_user_email_unique"`
	Password   string
	UserType   string `gorm:"size:1;comment:1 - Owner, 2 - Admin, 3 - Guide, 4 - Participant, 5 - Administrator"`
	Picture    string
	Active     string     `gorm:"size:1;comment:0 - NonActive, 1 - Active, 2 - Pending, 3 - Suspended"`
	UserDetail UserDetail `gorm:"foreignKey:UserID"`
}

// Builder Object for User
type UserBuilder struct {
	BaseModelBuilder
	userName   string
	email      string
	password   string
	userType   string
	picture    string
	activite   string
	userDetail UserDetail
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
func (u *UserBuilder) SetPicture(picture string) {
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
