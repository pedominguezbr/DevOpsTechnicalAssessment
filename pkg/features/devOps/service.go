package devOps

type Service interface {
	DevOps(s *RequestDevops) error
}

type devOpsService struct {
	repo Repository
}

// NewActualizaEstadoService new Service for Status
func NewdevOpsService(repo Repository) Service {
	return &devOpsService{
		repo: repo,
	}
}

// DevOps funcion services
func (svc *devOpsService) DevOps(srv *RequestDevops) error {

	//if len(srv.Details) != 0 {
	return svc.repo.DevOps(srv)
	//} else {
	//var auxInterface interface{}

	//	return errors.New("Required Fields Details")
	//}

}
