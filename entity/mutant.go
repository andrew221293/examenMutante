package entity

type (
	Request struct {
		Dna  []string `json:"dna"`
		Name string    `json:"name"`
	}

	Response struct {
		Dna  []string `json:"dna"`
		Name string    `json:"name"`
		Type string    `json:"type"`
	}
)
