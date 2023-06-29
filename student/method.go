package student

import (
	"github.com/Amirghrb/KinderServiceBookKeeping/user"
	"github.com/google/uuid"
)

type StudentInterface interface {
	GetID() uuid.UUID
	GetUser() user.UserData
	GetClassroomID() uuid.UUID
	GetBalance() int

	SetID(id uuid.UUID)
	SetUser(userData user.UserData)
	SetClassroomID(classroomId uuid.UUID)
	SetBalance(balance int)
}
