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

func defaultJadwal() []Jam {
	return []Jam{
		{"08.00 - 09.00", true},
		{"09.00 - 10.00", true},
		{"10.00 - 11.00", true},
		{"11.00 - 12.00", true},
		{"12.00 - 13.00", true},
		{"13.00 - 14.00", true},
		{"14.00 - 15.00", true},
		{"15.00 - 16.00", true},
		{"16.00 - 17.00", true},
		{"17.00 - 18.00", true},
		{"18.00 - 19.00", true},
		{"19.00 - 20.00", true},
		{"20.00 - 21.00", true},
		{"21.00 - 22.00", true},
	}
}

func displayLapName() {
	for i := 0; i < len(lapangan); i++ {
		fmt.Printf("%d. Lapangan %s - %s \n", i+1, lapangan[i].nama, lapangan[i].alamat)
	}
}

func ambilJam(index int) string {
	switch index {
	case 0:
		return "08.00"
	case 1:
		return "09.00"
	case 2:
		return "10.00"
	case 3:
		return "11.00"
	case 4:
		return "12.00"
	case 5:
		return "13.00"
	case 6:
		return "14.00"
	case 7:
		return "15.00"
	case 8:
		return "16.00"
	case 9:
		return "17.00"
	case 10:
		return "18.00"
	case 11:
		return "19.00"
	case 12:
		return "20.00"
	case 13:
		return "21.00"
	default:
		return "Jam tidak ditemukan"
	}
}
