package domainmodel

import "time"

type Results struct {
	ID        uint64     `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" sql:"index"`
	TestID    uint64     `json:"testID"`
	UserID    uint64     `json:"userID"`
	Result    float32    `json:"result"`
}

func (Results) TableName() string {
	return "onlinetests.results"
}
