package main

import (
	"cms-lab/function"
	"fmt"
)

func main() {
	testPolyndrome1 := function.Polyndrome("no lemon, no melon")
	testPolyndrome2 := function.Polyndrome("SAIPPUAKIVIKAUPPIAS")
	testPolyndrome3 := function.Polyndrome("Aibohphobia")
	testPolyndrome4 := function.Polyndrome("Anna")
	testPolyndrome5 := function.Polyndrome("My gym  ")
	testPolyndrome6 := function.Polyndrome("No lemon, no melon")
	testPolyndrome7 := function.Polyndrome("INI SALAH")

	fmt.Println(testPolyndrome1)
	fmt.Println(testPolyndrome2)
	fmt.Println(testPolyndrome3)
	fmt.Println(testPolyndrome4)
	fmt.Println(testPolyndrome5)
	fmt.Println(testPolyndrome6)
	fmt.Println(testPolyndrome7)
	function.FizzBuzz()
}

