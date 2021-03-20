package model

import "time"

type Post struct {
	BaseModel
	UrlSlug     string `gorm:"type:varchar(32); not null:true; unique:true; comment:url slug"`
	Title       string `gorm:"type:varchar(32); not null:true; comment:标题"`
	SubTitle    string `gorm:"type:varchar(32); not null:true; comment:子标题"`
	Content     string `gorm:"type:Text; not null:true; comment:内容"`
	Type        int8   `gorm:"not null:true; default:0; comment:0 为文章，1 为页面" json:"type"`
	PublishedAt *time.Time
	Category    *Category `gorm:"foreignKey:id"`
	Tags        []*Tag    `gorm:"-"`
}

func (Post) TableName() string {
	return "post"
}
