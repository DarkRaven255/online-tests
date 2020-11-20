package domainmodel

import (
	"online-tests/delivery/command"
	"time"

	"github.com/google/uuid"
)

type Test struct {
	ID             uint64     `json:"id" gorm:"primary_key"`
	CreatedAt      time.Time  `json:"-"`
	UpdatedAt      time.Time  `json:"updatedAt"`
	DeletedAt      *time.Time `json:"-" sql:"index"`
	Title          string     `json:"title"`
	UserID         uint64     `json:"userID"`
	NumOfQuestions uint       `json:"numOfQuestions"`
	TestUUID       string     `json:"testUUID" qorm:"unique"`
	Questions      []Question `json:"questions" gorm:"foreignKey:TestID"`
}

func (Test) TableName() string {
	return "onlinetests.tests"
}

func NewTestModel(cmd *command.AddTestCmd) Test {
	tuuid := uuid.New().String()
	return Test{
		Title:          cmd.Test.Title,
		UserID:         cmd.Test.UserID,
		NumOfQuestions: cmd.Test.NumOfQuestions,
		TestUUID:       tuuid,
		Questions:      *NewQuestionsArray(&cmd.Test.Questions),
	}
}
