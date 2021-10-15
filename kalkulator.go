package module

import "fmt"

func About() {
	fmt.Println("Ini adalah fungsi kalkulator")
}

func Tambah(a int, b int) int {
	return a + b
}

func Persegi(a int, b int) (int, int) {
	keliling := (a + b) * 2
	luas := a * b
	return keliling, luas
}
