package module

import "fmt"

type Active func(string) bool

func User(name string, isAktif Active) {

	if isAktif(name) {
		fmt.Println(name, " Status Aktif")
	} else {
		fmt.Println(name, " Status Non-AKtif")
	}
}
