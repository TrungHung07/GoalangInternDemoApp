// Package schema filepath: ent/schema/teacher.go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Student holds the schema definition for the Teacher entity.
type Student struct {
	ent.Schema
}

// Fields of the Teacher.
func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("class_id"),
		field.Bool("is_deleted").Default(false),
	}
}

// TableName define Name of entity
func (Student) TableName() string {
	return "students"
}

// Edges define relation of entities to each other
func (Student) Edges() []ent.Edge {
	return []ent.Edge{edge.To("classes", Class.Type).Unique().Required().Field("class_id")}
}
