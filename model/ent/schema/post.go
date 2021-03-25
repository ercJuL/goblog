package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"regexp"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("title_slug").Unique().NotEmpty().Match(regexp.MustCompile("^[a-z][A-Z-]+$")).Comment("url slug"),
		field.String("title").Comment("标题"),
		field.String("sub_title").Comment("子标题"),
		field.String("content").Comment("内容"),
		field.Enum("content_type").Values("markdown", "rich_text", "html").Comment("内容类型"),
		field.Enum("type").Values("page", "blog").Comment("发布类型"),
		field.Time("published_at").Optional().Nillable().Comment("发布时间"),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("categories", Category.Type).Ref("posts").Unique(),
		edge.From("tags", Tag.Type).Ref("posts"),
	}
}
