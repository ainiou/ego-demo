package db

type HelloWorld struct {
	Id    int64  `json:"id" gorm:"column:id"`       // id
	Hello string `json:"hello" gorm:"column:hello"` // hello
	World string `json:"world" gorm:"column:world"` // hello
}

func (m *HelloWorld) TableName() string {
	return "hello_world"
}
