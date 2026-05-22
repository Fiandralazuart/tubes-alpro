package main

import "fmt"

func runProject() {
	fmt.Println("Selamat Datang di Tubes Alpro")
}

func main() {
	var n int
	fmt.Println("=== Menu Sport ===")
	fmt.Println("Pilih Menu:")
	fmt.Println("1. Data Lapangan")
	fmt.Println("2. Member Lapangan")
	fmt.Println("3. Lihat Jadwal")
	fmt.Scan(&n)

	switch {
	case n == 1:
		mainCrud()
	case n == 2:
		menuPenyewa()
	case n == 3:
		mainReservation()
	default:
		fmt.Println("Perintah Tidak Valid")
		menuLain(main)
	}
}
