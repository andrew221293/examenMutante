package store

const (
	mutantStatement = `
	INSERT INTO mutants(
		dna,
		name
	) VALUES (
		$1,
		$2
	);`

	humanStatement = `
	INSERT INTO humans(
		dna,
		name
	) VALUES (
		$1,
		$2
	);`

	mutantsTotalStatement = `
	SELECT count(*) FROM mutants;`

	humansTotalStatement = `
	SELECT count(*) FROM humans;`
)
