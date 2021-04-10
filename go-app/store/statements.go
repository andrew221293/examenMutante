package store

const (
	createTableMutants = `CREATE TABLE IF NOT EXISTS mutants
	(dna varchar unique, name varchar);`

	createTableHumans = `CREATE TABLE IF NOT EXISTS humans
	( dna varchar unique, name varchar);`

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
