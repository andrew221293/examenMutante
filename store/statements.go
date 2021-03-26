package store

const (
	mutantStatement = `
	INSERT INTO mutants(
		dna,
		name,
	) VALUES (
		$1,
		$2,
	);`
)
