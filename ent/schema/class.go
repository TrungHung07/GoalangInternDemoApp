// filepath: ent/schema/teacher.go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Teacher holds the schema definition for the Teacher entity.
type Class struct {
	ent.Schema
}

// Fields of the Teacher.
func (Class) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int64("grade"),
		field.Bool("is_deleted").Default(false),
	}
}

func (Class) TableName() string {
	return "class"
}

func (Class) Edges() []ent.Edge {
	return []ent.Edge{edge.From("students", Student.Type).Ref("classes")}
}
