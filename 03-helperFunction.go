package main

import "fmt"

func displayLap(lapangan []Lapangan, isLoop bool, n int, message string) int {
	x := 0
	if isLoop {
		for i := 0; i < len(lapangan); i++ {
			fmt.Printf("Lapangan-%d \n", i+1)
			fmt.Printf("Nama Lap: %s \n", lapangan[i].nama)
			fmt.Printf("Alamat Lap: %s \n", lapangan[i].alamat)
			fmt.Printf("Jenis Lap: %s \n", lapangan[i].jenis)
			fmt.Println("--- Harga Sewa ---")
			fmt.Printf("Harga Standar: Rp %d \n", lapangan[i].harga.hargaDefault)
			fmt.Printf("Harga Happy Hour: Rp %d \n", lapangan[i].harga.happyHour)
			fmt.Println(" ")
			x++
		}
	} else {
		fmt.Println(" ")
		fmt.Printf("%s-%d \n", message, n)
		fmt.Printf("Nama Lap: %s \n", lapangan[n-1].nama)
		fmt.Printf("Alamat Lap: %s \n", lapangan[n-1].alamat)
		fmt.Printf("Jenis Lap: %s \n", lapangan[n-1].jenis)
		fmt.Println("--- Harga Sewa ---")
		fmt.Printf("Harga Standar: Rp %d \n", lapangan[n-1].harga.hargaDefault)
		fmt.Printf("Harga Happy Hour: Rp %d \n", lapangan[n-1].harga.happyHour)
		fmt.Println(" ")
		x++
	}
	return x
}
func menuLain(fungsi func()) {
	var cond string

	fmt.Println("Buka Menu Lain? (yes/no)")
	fmt.Scan(&cond)

	if cond == "yes" {
		fungsi()
	} else if cond == "no" {
		fmt.Println("Selesai")
	} else {
		fmt.Println("Perintah Tidak Valid")
	}
}

func defaultJadwal(hargaDefault, happyHour int) []Jam {
	return []Jam{
		{"08.00 - 09.00", true, happyHour},
		{"09.00 - 10.00", true, happyHour},
		{"10.00 - 11.00", true, happyHour},
		{"11.00 - 12.00", true, hargaDefault},
		{"12.00 - 13.00", true, hargaDefault},
		{"13.00 - 14.00", true, hargaDefault},
		{"14.00 - 15.00", true, hargaDefault},
		{"15.00 - 16.00", true, hargaDefault},
		{"16.00 - 17.00", true, happyHour},
		{"17.00 - 18.00", true, happyHour},
		{"18.00 - 19.00", true, happyHour},
		{"19.00 - 20.00", true, happyHour},
		{"20.00 - 21.00", true, hargaDefault},
		{"21.00 - 22.00", true, hargaDefault},
	}
}

func displayLapName() int {
	n := 0
	for i := 0; i < len(lapangan); i++ {
		fmt.Printf("%d. Lapangan %s - %s \n", i+1, lapangan[i].nama, lapangan[i].alamat)
		n++
	}
	return n
}
// halo

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

func ambilIndexJam(jam string) int {
	switch jam {
	case "08.00":
		return 0
	case "09.00":
		return 1
	case "10.00":
		return 2
	case "11.00":
		return 3
	case "12.00":
		return 4
	case "13.00":
		return 5
	case "14.00":
		return 6
	case "15.00":
		return 7
	case "16.00":
		return 8
	case "17.00":
		return 9
	case "18.00":
		return 10
	case "19.00":
		return 11
	case "20.00":
		return 12
	case "21.00":
		return 13
	default:
		return -1
	}
}

func loopReservasi(i int) {
	for j := 0; j < len(database[i].reservasi); j++ {
		fmt.Println("")
		fmt.Printf("Nama: %s\n", database[i].reservasi[j].penyewa)
		fmt.Printf("Tanggal: %s - %s\n", database[i].reservasi[j].tglMulai, database[i].reservasi[j].tglAkhir)
		fmt.Printf("Jam: %s - %s\n", database[i].reservasi[j].jamMulai, database[i].reservasi[j].jamAkhir)
		fmt.Printf("Durasi: %d jam\n", database[i].reservasi[j].durasi)
		fmt.Printf("Harga: Rp. %d\n", database[i].reservasi[j].total)
		fmt.Println("")
	}
}

func displayJadwal(data []Jam) {
	for j := 0; j < len(data); j++ {
		if data[j].isAvailable {
			fmt.Printf("%d. %s | Tersedia | Rp. %d\n", j+1, data[j].waktu, data[j].harga)
		} else {
			fmt.Printf("%d. %s | Terbooking | Rp. %d\n", j+1, data[j].waktu, data[j].harga)
		}
	}
}

func selectionSort(jadwal []Jam) {
	for a := 0; a < len(jadwal)-1; a++ {
		max := a
		for b := a + 1; b < len(jadwal); b++ {
			if !jadwal[b].isAvailable &&
				jadwal[max].isAvailable {
				max = b
			}
		}
		jadwal[a], jadwal[max] =
			jadwal[max], jadwal[a]
	}
}

func insertionSort(jadwal []Jam) {
	for a := 1; a < len(jadwal); a++ {
		key := jadwal[a]
		b := a - 1
		for b >= 0 &&
			jadwal[b].isAvailable &&
			!key.isAvailable {
			jadwal[b+1] = jadwal[b]
			b--
		}
		jadwal[b+1] = key
	}
}

func selectionSortHarga(jadwal []Jam) {
	for i := 0; i < len(jadwal)-1; i++ {
		min := i
		for j := i + 1; j < len(jadwal); j++ {
			if jadwal[j].harga < jadwal[min].harga {
				min = j
			}
		}
		jadwal[i], jadwal[min] = jadwal[min], jadwal[i]
	}
}

func insertionSortHarga(jadwal []Jam) {
	for i := 1; i < len(jadwal); i++ {

		key := jadwal[i]
		j := i - 1
		for j >= 0 &&
			jadwal[j].harga > key.harga {
			jadwal[j+1] = jadwal[j]
			j--
		}
		jadwal[j+1] = key
	}
}

func validateDate(x int, b int) int {
	if x == 0 {
		return 0
	}

	if x < 1 || x > b {
		fmt.Println("Data Tidak Valid")
	}

	return 1
}