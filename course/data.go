package course

import (
	"time"

	"github.com/google/uuid"
)

type CourseData struct {
	id          uuid.UUID
	name        string
	description string
	startDate   time.Time
	endDate     time.Time
	classroomId uuid.UUID
	studentsId  []uuid.UUID
	teachersId  []uuid.UUID
}

// Getters methods
func (c *CourseData) GetID() uuid.UUID {
	return c.id
}

func (c *CourseData) GetName() string {
	return c.name
}

func (c *CourseData) GetDescription() string {
	return c.description
}

func (c *CourseData) GetStartDate() time.Time {
	return c.startDate
}

func (c *CourseData) GetEndDate() time.Time {
	return c.endDate
}

func (c *CourseData) GetStudentsID() []uuid.UUID {
	return c.studentsId
}

func (c *CourseData) GetTeachersID() []uuid.UUID {
	return c.teachersId
}

// Setter methods

func (c *CourseData) SetID(id uuid.UUID) {
	c.id = id
}

func (c *CourseData) SetName(name string) {
	c.name = name
}

func (c *CourseData) SetDescription(description string) {
	c.description = description
}

func (c *CourseData) SetStartDate(startDate time.Time) {
	c.startDate = startDate
}

func (c *CourseData) SetEndDate(endDate time.Time) {
	c.endDate = endDate
}

func (c *CourseData) SetStudentsID(studentsId []uuid.UUID) {
	c.studentsId = studentsId
}

func (c *CourseData) SetTeachersID(teachersId []uuid.UUID) {
	c.teachersId = teachersId
}
