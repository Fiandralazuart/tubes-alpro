package main

import (
	"fmt"
)

func mainReservation() {
	var n int

	fmt.Println("=== MENU RESERVASI LAPANGAN ===")
	fmt.Println("Pilih Menu:")
	fmt.Println("1. Tambah Member")
	fmt.Println("2. Tampilkan Jadwal")
	fmt.Println("3. Buat Jadwal")
	fmt.Println("4. Reservasi")
	fmt.Print("Menu: ")
	fmt.Scan(&n)

	switch {
		case n == 1:
			tampilkanJadwal()
		case n == 2:
			tampilkanJadwal()
		case n == 3:
			buatJadwal()
		case n == 4:
			hapusLapangan()
		default:
			fmt.Println("Perintah Tidak Valid")
			menuLain(mainReservation)
	}
}

func tampilkanJadwal() {
	var n int
	fmt.Println("=== Jadwal Lapangan ===")
	fmt.Println("Pilih Lapangan: ")
	displayLapName()

	fmt.Println("Pilih: ")
	fmt.Scan(&n)


	for i := 0; i < len(database); i++ {
		if database[i].lapangan == lapangan[n-1].nama {
			fmt.Println("==== Jadwal Tersedia ====")
			fmt.Printf("Lap. %s | %s\n", database[i].lapangan, database[i].tanggal)
			fmt.Println("")
			for j := 0; j < len(database[i].jadwal); j++ {
				if database[i].jadwal[j].isAvailable {
					fmt.Printf("%s | Tersedia \n", database[i].jadwal[j].waktu)
					} else {
					fmt.Printf("%s | Tersedia \n", database[i].jadwal[j].waktu)
				}
			}
			fmt.Println("")
		}
	}
	menuLain(mainReservation)
}

func buatJadwal() {
	var n int
	var tanggal string
	fmt.Println("=== Buat Jadwal ===")
	fmt.Println("Pilih Lapangan: ")
	displayLapName()

	fmt.Println("Pilih: ")
	fmt.Scan(&n)
	
	fmt.Print("Masukkan Tanggal: ")
	fmt.Scan(&tanggal)

	isAvailable := false
	for i := 0; i < len(database); i++ {
		if database[i].tanggal != tanggal && lapangan[n-1].nama == database[i].lapangan {
			isAvailable = true
		}
	}

	if isAvailable {
		database = append(database, Database{
			tanggal: tanggal,
			lapangan: lapangan[n-1].nama,
			jadwal: defaultJadwal(),
		})
		fmt.Println("Sukses Membuat Jadwal")
		menuLain(mainReservation)
	} else {
		fmt.Println("Gagal: Jadwal Sudah Ada")
		menuLain(mainReservation)
	}

}