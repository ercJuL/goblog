package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("name_slug").Unique().NotEmpty().Comment("分类slug"),
		field.String("name").Unique().Comment("分类"),
		field.String("desc").Comment("描述"),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
	}
}
