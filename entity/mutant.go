package entity

type (
	Request struct {
		Dna  [6]string `json:"dna"`
		Name string    `json:"name"`
	}

	Response struct {
		Dna  [6]string `json:"dna"`
		Name string    `json:"name"`
		Type string    `json:"type"`
	}
)
