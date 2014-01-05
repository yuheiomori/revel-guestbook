package models

type Greeting struct {
	Id       int64
	Name     string `sql:"size:100"`
	Comment  string `sql:"size:200"`
	CreateAt int64
}
