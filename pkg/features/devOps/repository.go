package devOps

type Repository interface {
	DevOps(u *RequestDevops) error
}
