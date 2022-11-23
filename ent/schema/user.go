package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

// Fields of the Todo.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").NotEmpty(),
		field.Text("email").NotEmpty(),
		field.Text("password").NotEmpty(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Enum("status").
			NamedValues(
				"Enable", "ENABLE",
				"Disable", "DISABLE",
			).
			Default("ACTIVE"),
	}
}

// Edges of the Todo.
func (User) Edges() []ent.Edge {
	return nil
}
