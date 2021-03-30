package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_name").Unique().NotEmpty().Comment("用户名"),
		field.String("password").Sensitive().NotEmpty().Comment("密码"),
		field.String("nick_name").NotEmpty().Comment("用户昵称"),
		field.String("header_img").Comment("头像"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
