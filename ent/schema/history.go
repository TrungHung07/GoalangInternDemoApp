package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// History holds the schema definition for the History entity.
type History struct {
	ent.Schema
}

// Fields of the History.
func (History) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			Immutable(),
		field.String("table_name").
			Comment("Tên bảng được thao tác"),
		field.String("record_id").
			Comment("ID của record được thao tác"),
		field.String("action").
			Comment("Loại thao tác: INSERT, UPDATE, DELETE"),
		field.JSON("old_data", map[string]interface{}{}).
			Optional().
			Comment("Dữ liệu cũ (cho UPDATE/DELETE)"),
		field.JSON("new_data", map[string]interface{}{}).
			Optional().
			Comment("Dữ liệu mới (cho INSERT/UPDATE)"),
		field.String("user_id").
			Optional().
			Comment("ID user thực hiện thao tác"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.JSON("metadata", map[string]interface{}{}).
			Optional().
			Comment("Thông tin bổ sung"),
	}
}

// Indexes of the History.
func (History) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("table_name", "record_id"),
		index.Fields("action"),
		index.Fields("created_at"),
		index.Fields("user_id"),
	}
}
