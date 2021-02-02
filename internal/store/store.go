package store

type Store interface {
	User() UserRepo
	Course() CourseRepo
	Module() ModuleRepo
	Card() CardRepo
}
