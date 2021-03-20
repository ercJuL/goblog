package model

type Category struct {
	BaseModel
}

func (Category) TableName() string {
	return "category"
}
