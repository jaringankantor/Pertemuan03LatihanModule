package module

import "fmt"

type Active func(string) bool

type Pegawai struct {
	LamaBekerja  int8
	Nama, Divisi string
	Aktif        bool
}

func User(name string, isAktif Active) {

	if isAktif(name) {
		fmt.Println(name, " Status Aktif")
	} else {
		fmt.Println(name, " Status Non-AKtif")
	}
}

func TampilData() {
	pegawai := Pegawai{
		LamaBekerja: 5,
		Nama:        "Anwars",
		Divisi:      "Keuangan",
		Aktif:       true,
	}

	fmt.Println(pegawai)
}

func (pegawai Pegawai) SelamatDatang(nama string) {
	fmt.Println("Selamat datang", nama, "giat bekerja ya pak", pegawai.Nama)
}

func AnonymousDataOrtu() {
	OrangTua := struct {
		Usia     int8
		NamaOrtu string
		Hidup    bool
	}{
		Usia:     63,
		NamaOrtu: "Syahroni",
		Hidup:    true,
	}

	fmt.Println(OrangTua)
}
