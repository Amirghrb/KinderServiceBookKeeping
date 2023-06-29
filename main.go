package main

import (
	"fmt"

	"github.com/Amirghrb/KinderServiceBookKeeping/employee"
	"github.com/Amirghrb/KinderServiceBookKeeping/user"
)

func main() {
	u := user.UserData{}
	u.SetName("amir")
	u.Setbirthday("2000/2/9")
	u.Setid()
	e := employee.EmployeeData{}
	e.SetUser(u)
	e.SetTitle(employee.MANAGER)
	fmt.Println(u.GetName())
	fmt.Println(u.Getid())
	//TODO - add handler
}
