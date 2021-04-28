package model

type Literature struct {
	LitType string `json:"type"`
	Name    string `json:"name"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
