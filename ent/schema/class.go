// Package schema filepath: ent/schema/teacher.go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Class holds the schema definition for the Class entity.
type Class struct {
	ent.Schema
}

// Fields of the Class.
func (Class) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int64("grade"),
		field.Bool("is_deleted").Default(false),
	}
}

// TableName overide name
func (Class) TableName() string {
	return "class"
}

// Edges to map entity
func (Class) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("students", Student.Type).Ref("classes"),
		edge.From("teachers", Teacher.Type).Ref("classes"),
	}
}
