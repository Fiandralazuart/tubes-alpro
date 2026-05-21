package main
import "fmt"

func displayLap(lapangan []Lapangan, isLoop bool, n int, message string) {
	if isLoop {
		for i := 0; i < len(lapangan); i++ {
			fmt.Printf("Lapangan-%d \n", i+1)
			fmt.Printf("Nama Lap: %s \n", lapangan[i].nama)
			fmt.Printf("Alamat Lap: %s \n", lapangan[i].alamat)
			fmt.Printf("Jenis Lap: %s \n", lapangan[i].jenis)
			fmt.Printf("Harga Sewa: %d \n", lapangan[i].harga)
			fmt.Println(" ")
		}
	} else {
		fmt.Println(" ")
		fmt.Printf("%s-%d \n", message, n)
		fmt.Printf("Nama Lap: %s \n", lapangan[n-1].nama)
		fmt.Printf("Alamat Lap: %s \n", lapangan[n-1].alamat)
		fmt.Printf("Jenis Lap: %s \n", lapangan[n-1].jenis)
		fmt.Printf("Harga Sewa: %d \n", lapangan[n-1].harga)
		fmt.Println(" ")
	}
}

func menuLain() {
	var cond string

	fmt.Println("Buka Menu Lain? (yes/no)")
	fmt.Scan(&cond)

	if cond == "yes" {
		mainCrud()
	} else if cond == "no" {
		fmt.Println("Selesai")
	} else {
		fmt.Println("Perintah Tidak Valid")
	}
}