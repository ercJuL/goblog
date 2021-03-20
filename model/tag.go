package model

type Tag struct {
	BaseModel
}

func (Tag) TableName() string {
	return "tag"
}
