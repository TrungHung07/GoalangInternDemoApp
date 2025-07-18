// Code generated by ent, DO NOT EDIT.

package ent

import (
	"DemoApp/ent/class"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Class is the model entity for the Class schema.
type Class struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Grade holds the value of the "grade" field.
	Grade int64 `json:"grade,omitempty"`
	// IsDeleted holds the value of the "is_deleted" field.
	IsDeleted bool `json:"is_deleted,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ClassQuery when eager-loading is set.
	Edges        ClassEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ClassEdges holds the relations/edges for other nodes in the graph.
type ClassEdges struct {
	// Students holds the value of the students edge.
	Students []*Student `json:"students,omitempty"`
	// Teachers holds the value of the teachers edge.
	Teachers []*Teacher `json:"teachers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// StudentsOrErr returns the Students value or an error if the edge
// was not loaded in eager-loading.
func (e ClassEdges) StudentsOrErr() ([]*Student, error) {
	if e.loadedTypes[0] {
		return e.Students, nil
	}
	return nil, &NotLoadedError{edge: "students"}
}

// TeachersOrErr returns the Teachers value or an error if the edge
// was not loaded in eager-loading.
func (e ClassEdges) TeachersOrErr() ([]*Teacher, error) {
	if e.loadedTypes[1] {
		return e.Teachers, nil
	}
	return nil, &NotLoadedError{edge: "teachers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Class) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case class.FieldIsDeleted:
			values[i] = new(sql.NullBool)
		case class.FieldID, class.FieldGrade:
			values[i] = new(sql.NullInt64)
		case class.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Class fields.
func (c *Class) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case class.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case class.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case class.FieldGrade:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field grade", values[i])
			} else if value.Valid {
				c.Grade = value.Int64
			}
		case class.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				c.IsDeleted = value.Bool
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Class.
// This includes values selected through modifiers, order, etc.
func (c *Class) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryStudents queries the "students" edge of the Class entity.
func (c *Class) QueryStudents() *StudentQuery {
	return NewClassClient(c.config).QueryStudents(c)
}

// QueryTeachers queries the "teachers" edge of the Class entity.
func (c *Class) QueryTeachers() *TeacherQuery {
	return NewClassClient(c.config).QueryTeachers(c)
}

// Update returns a builder for updating this Class.
// Note that you need to call Class.Unwrap() before calling this method if this Class
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Class) Update() *ClassUpdateOne {
	return NewClassClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Class entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Class) Unwrap() *Class {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Class is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Class) String() string {
	var builder strings.Builder
	builder.WriteString("Class(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("grade=")
	builder.WriteString(fmt.Sprintf("%v", c.Grade))
	builder.WriteString(", ")
	builder.WriteString("is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", c.IsDeleted))
	builder.WriteByte(')')
	return builder.String()
}

// Classes is a parsable slice of Class.
type Classes []*Class
