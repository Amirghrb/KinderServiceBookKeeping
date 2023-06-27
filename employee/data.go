package employee

import (
	"github.com/Amirghrb/KinderServiceBookKeeping/user"
)

type Title struct {
	val string
}

func (r Title) String() string {
	return r.val
}

var (
	MANAGER = Title{"MANAGER"}
	TEACHER = Title{"TEACHER"}
	CLEANER = Title{"CLEANER"}
	DEPUTY  = Title{"DEPUTY"}
)

type EmployeeData struct {
	user              user.UserData
	nationalId        int
	bankAccountNumber int
	salary            int
	title             Title
}

func (e *EmployeeData) SetUser(u user.UserData) {
	e.user = u
}

func (e *EmployeeData) GetUser() user.UserData {
	return e.user
}

func (e *EmployeeData) SetTitle(t Title) {
	e.title = t
}

func (e *EmployeeData) GetTitle() Title {
	return e.title
}

func (e *EmployeeData) GetNationalId() int {
	return e.nationalId
}

func (e *EmployeeData) GetBankAccountNumber() int {
	return e.bankAccountNumber
}

func (e *EmployeeData) GetSalary() int {
	return e.salary
}

// Setter methods

func (e *EmployeeData) SetNationalId(id int) {
	e.nationalId = id
}

func (e *EmployeeData) SetBankAccountNumber(num int) {
	e.bankAccountNumber = num
}

func (e *EmployeeData) SetSalary(sal int) {
	e.salary = sal
}
