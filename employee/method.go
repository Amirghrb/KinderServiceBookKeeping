package employee

import "github.com/Amirghrb/KinderServiceBookKeeping/user"

type EmployeeUsecase interface {
	SetUser(u user.UserData)
	GetUser() user.UserData
	GetSalary() int
	GetNationalId() int
	GetBankAccountNumber(int)
	SetSalary(int)
	SetNationalId(int)
	SetBankAccountNumber(int)
}
