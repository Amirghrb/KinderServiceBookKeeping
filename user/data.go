package user

import "github.com/google/uuid"

type UserData struct {
	name     string
	id       uuid.UUID
	birthday string
}

func (u *UserData) SetName(name string) {
	u.name = name
}
func (u *UserData) Setid() {
	u.id = uuid.New()
}
func (u *UserData) Setbirthday(b string) {
	u.birthday = b
}
func (u *UserData) GetName() string {
	return u.name
}
func (u *UserData) Getid() uuid.UUID {
	return u.id
}
func (u *UserData) Getbirthday() string {
	return u.birthday
}
