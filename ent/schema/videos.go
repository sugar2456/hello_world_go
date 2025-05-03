package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// Videos holds the schema definition for the Videos entity.
type Videos struct {
	ent.Schema
}

// Fields of the Videos.
func (Videos) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			NotEmpty(),
		field.String("description").
			Optional(),
		field.String("url").
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").
			Optional(),
		field.String("thumbnail").
			Optional(),
		field.String("category").
			Optional(),
		field.String("tags").
			Optional(),
	}
}

// Edges of the Videos.
func (Videos) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("videos").
			Unique(),
	}
}
