package user

type UserUsecase interface {
	setName(name string)
	setbirthday(birthday string)
	setid()
	getName() string
	getbirthday() string
	getid() int
}
