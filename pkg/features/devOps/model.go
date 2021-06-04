package devOps

// devOps.RequestDevops representa el modelo para una peticion del api devOps
type RequestDevops struct {
	Message       string  `json:"message”" example:"This is a test"`
	To            string  `json:"to" example:"Juan Perez"`
	From          string  `json:"from" example:"Rita Asturia"`
	TimeToLifeSec float64 `json:"timeToLifeSec" example:"45"`
}
