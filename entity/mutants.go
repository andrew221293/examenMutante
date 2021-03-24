package entity

type (
	MutantsRequest struct {
		dna [6]string `json:"dna"`
		name string `json:"name"`
	}
)