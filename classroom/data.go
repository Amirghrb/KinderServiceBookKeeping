package classroom

import (
	"github.com/Amirghrb/KinderServiceBookKeeping/employee"
	"github.com/google/uuid"
)

type ClassroomData struct {
	id       uuid.UUID
	capacity int
	teachers []employee.EmployeeData
}
