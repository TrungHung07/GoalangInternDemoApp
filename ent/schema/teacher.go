// Package schema filepath: ent/schema/teacher.go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Teacher holds the schema definition for the Teacher entity.
type Teacher struct {
	ent.Schema
}

// Fields of the Teacher.
func (Teacher) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("email"),
		field.Int("class_id"),
		field.Int("age").Positive(),
		field.Bool("is_deleted").Default(false),
	}
}

// TableName define name of table for orm
func (Teacher) TableName() string {
	return "teachers"
}

// Edges define relations of entity to each other.
func (Teacher) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("classes", Class.Type).Unique().Required().Field("class_id"),
	}
}
