package entity

type TaskType int

const (
	Special TaskType = iota + 1
	Leagues
	RefTasks
	LevelUp
)

type Task struct {
	Id       int64  `json:"id"`
	Name     string `json:"name" gorm:"unique"`
	Cost     int64  `json:"cost" `
	Type     TaskType
	UniqueId string `json:"unique_id" gorm:"unique"`
}

type TaskUser struct {
	Id     int64 `json:"id"`
	TaskId int64
	Task   Task `gorm:"foreignKey:Id;references:TaskId"`
	IdUser int64
	Cost   int64 `json:"cost"`
}
