package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "time"
)

type Playlist struct {
	ent.Schema
}

func (Playlist) Fields() []ent.Field {
	return []ent.Field{
        field.String("title").NotEmpty(),
        field.String("description").Optional(),
        field.Int("user_id").Optional(),
        field.Time("created_at").Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (Playlist) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("user", User.Type).
            Ref("playlists").
            Field("user_id").
            Unique(),
        edge.To("videos", Videos.Type),
    }
}