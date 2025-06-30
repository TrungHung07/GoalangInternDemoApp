package filter

//TeacherFilter store filter field to list all teachers
type TeacherFilter struct {
	MinAge *int32
	MaxAge *int32
	Email  *string
}
