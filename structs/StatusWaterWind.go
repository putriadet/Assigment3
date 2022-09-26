package structs

type StatusWaterWind struct {
	Status struct {
		Water int `json:"water"`
		Wind  int `json:"Wind"`
	} `json:"status"`
}
