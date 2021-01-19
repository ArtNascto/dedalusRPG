package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive(),
		field.String("name").
			Default("unknown"),
		field.String("surname").
			Default("unknown"),
		field.String("username").
			Default("unknown"),
		field.String("email").
			Default("unknown"),
		field.String("phone_number").
			Default("unknown"),
		field.String("days_off").
			Default("unknown"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
