package main
import "fmt"

func main() {
	var n int
	fmt.Println("=== Menu Sport ===")
	fmt.Println("Pilih Menu:")
	fmt.Println("1. CRUD Lapangan")
	fmt.Println("2. Lihat Jadwal")
	fmt.Scan(&n)

	switch {
		case n == 1:
			mainCrud()
		case n == 2:
			mainReservation()
		default:
			fmt.Println("Perintah Tidak Valid")
			menuLain(main)
	}
}