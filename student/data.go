package student

import (
	"github.com/Amirghrb/KinderServiceBookKeeping/user"
	"github.com/google/uuid"
)

type StudentData struct {
	id       uuid.UUID
	user     user.UserData
	courseId uuid.UUID
	balance  int
}

func (s *StudentData) GetID() uuid.UUID {
	return s.id
}

func (s *StudentData) GetUser() user.UserData {
	return s.user
}

func (s *StudentData) GetClassroomID() uuid.UUID {
	return s.courseId
}

func (s *StudentData) GetBalance() int {
	return s.balance
}

// Setter methods

func (s *StudentData) SetID(id uuid.UUID) {
	s.id = id
}

func (s *StudentData) SetUser(userData user.UserData) {
	s.user = userData
}

func (s *StudentData) SetClassroomID(classroomId uuid.UUID) {
	s.courseId = classroomId
}

func (s *StudentData) SetBalance(balance int) {
	s.balance = balance
}
