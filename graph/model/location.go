package model

type Location struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Note      string  `json:"note"`
	MapUrl    string  `json:"mapUrl"`
}

type CreateInput struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Note      string  `json:"note"`
	MapUrl    string  `json:"mapUrl"`
}

type UpdateInput struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Note      string  `json:"note"`
	MapUrl    string  `json:"mapUrl"`
}
