package services

type Data struct {
	Status Status `json:"status"`
}

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type Response struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Value  int    `json:"value"`
}
