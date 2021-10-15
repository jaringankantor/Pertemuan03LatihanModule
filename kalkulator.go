package module

import "fmt"

type PersegiInterface interface {
	Keliling() int
	Luas() int
}

type InputPersegi struct {
	a, b int
}

func (in InputPersegi) Keliling() int {
	return (in.a + in.b) * 2
}

func (in InputPersegi) Luas() int {
	return in.a * in.b
}

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
