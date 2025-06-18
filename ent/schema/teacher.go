// filepath: ent/schema/teacher.go
package schema

import (
	"entgo.io/ent"
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
		field.Int("grade"),
	}
}

func (Teacher) TableName() string {
	return "teachers"
}
