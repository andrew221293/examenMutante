package usecase

import (
	"fmt"
	"strings"
)

func isMutant(dna[6]string) bool {
	adn := [6][6]string{}
	for a, values := range dna {
		newArray := strings.Split(values, "")
		for b, value := range newArray {
			adn[a][b] = value
		}
	}
	fmt.Println("  0  1  2  3  4  5")
	fmt.Println(" -------------------------------------------------")
	fmt.Println("0 "+adn[0][0]+ " " +adn[0][1]+" " + adn[0][2]+"    " + adn[0][3]+"    " + adn[0][4]+"    " + adn[0][5])
	fmt.Println(" -------------------------------------------------")
	fmt.Println("1 "+adn[1][0]+ " " +adn[1][1]+" " + adn[1][2]+"    " + adn[1][3]+"    " + adn[1][4]+"    " + adn[1][5])
	fmt.Println(" -------------------------------------------------")
	fmt.Println("2 "+adn[2][0]+ " " +adn[2][1]+" " + adn[2][2]+"    " + adn[2][3]+"    " + adn[2][4]+"    " + adn[2][5])
	fmt.Println(" -------------------------------------------------")
	fmt.Println("3 "+adn[3][0]+ " " +adn[3][1]+" " + adn[3][2]+"    " + adn[3][3]+"    " + adn[3][4]+"    " + adn[3][5])
	fmt.Println(" -------------------------------------------------")
	fmt.Println("4 "+adn[4][0]+ " " +adn[4][1]+" " + adn[4][2]+"    " + adn[4][3]+"    " + adn[4][4]+"    " + adn[4][5])
	fmt.Println(" -------------------------------------------------")
	fmt.Println("5 "+adn[5][0]+ " " +adn[5][1]+" " + adn[5][2]+"    " + adn[5][3]+"    " + adn[5][4]+"    " + adn[5][5])
	fmt.Println(" -------------------------------------------------")

	horizontal(dna)
	vertical(dna)
	oblique(dna)

	return true
}

func horizontal(dna[6]string) {
	for _, sequence := range dna {
		foundSequence(sequence)
	}
}
func vertical (dna[6]string) {
	for i:= 0; i < len(dna); i++ {
		var sequence string
		for _, values := range dna {
			sequence = sequence + string(values[i])
		}
		foundSequence(sequence)
	}
}
func oblique (dna[6]string) {
	var sequence string
	for i:= 0; i < len(dna); i++ {
		for index, values:= range dna {
			if i == index {
				sequence = sequence + string(values[i])
			}
		}
	}
	foundSequence(sequence)
}
func foundSequence (sequence string) bool {
	sequenceToFind := [4]string{"AAAA", "CCCC", "GGGG", "TTTT"}
	for _, n := range sequenceToFind {
		if sequence[0:4] == n {
			fmt.Println(sequence[0:4])
			fmt.Println("FOUND")
			return true
		}
	}
	return false
}