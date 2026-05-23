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
	fmt.Println("4. Booking Jadwal")
	fmt.Println("5. History Reservasi")
	fmt.Println("6. Kembali")
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
		bookingJadwal()
	case n == 5:
		tampilkanReservasi()
	case n == 6:
		main()
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
					fmt.Printf("%s | Terbooking \n", database[i].jadwal[j].waktu)
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
			tanggal:  tanggal,
			lapangan: lapangan[n-1].nama,
			jadwal:   defaultJadwal(),
		})
		fmt.Println("Sukses Membuat Jadwal")
		menuLain(mainReservation)
	} else {
		fmt.Println("Gagal: Jadwal Sudah Ada")
		menuLain(mainReservation)
	}

}

func bookingJadwal() {
	var n, idx int
	var tanggal, nama string
	// var data Database
	fmt.Println("=== Booking Jadwal ===")
	jml := displayLapName()

	fmt.Println("Pilih: ")
	fmt.Scan(&n)

	if n > jml {
		fmt.Println("Perintah Tidak Valid")
		menuLain(mainReservation)
	}

	fmt.Println("Masukkan Tanggal: ")
	fmt.Scan(&tanggal)

	isThere := false
	for i := 0; i < len(database); i++ {
		if database[i].lapangan == lapangan[n-1].nama && database[i].tanggal == tanggal {
			// data = database[i]
			idx = i
			fmt.Println("==== Jadwal Tersedia ====")
			fmt.Printf("Lap. %s | %s\n", database[i].lapangan, database[i].tanggal)
			fmt.Println("")
			for j := 0; j < len(database[i].jadwal); j++ {
				if database[i].jadwal[j].isAvailable {
					fmt.Printf("%d. %s | Tersedia \n", j+1, database[i].jadwal[j].waktu)
				} else {
					fmt.Printf("%d. %s | Terbooking \n", j+1, database[i].jadwal[j].waktu)
				}
			}
			isThere = true
			fmt.Println("")
		}
	}
	if !isThere {
		cond := ""
		fmt.Println("Jadwal Belum Ada, Buat Jadwal? (yes/no)")
		fmt.Scan(&cond)

		if cond == "yes" {
			buatJadwal()
		} else {
			menuLain(mainReservation)
		}
	}

	fmt.Println("Nama Member: ")
	fmt.Scan(&nama)

	isMember := false
	for i := 0; i < len(penyewa); i++ {
		if penyewa[i].nama == nama {
			isMember = true
		}
	}

	jmlJam := 0
	if isMember {

		fmt.Println("Jumlah Jam: ")
		fmt.Scan(&jmlJam)

		var x int
		for i := 0; i < jmlJam; i++ {
			fmt.Println("Pilih Jadwal: ")
			fmt.Scan(&x)

			database[idx].jadwal[x-1].isAvailable = false
		}
		database[idx].reservasi = append(database[idx].reservasi, Sewa{
			penyewa:  nama,
			jamMulai: ambilJam(x-n),
			jamAkhir: ambilJam(x),
			tglMulai: tanggal,
			tglAkhir: tanggal,
			durasi: n,
		})
	} else {
		fmt.Println("Member Tidak Terdaftar, Tambahkan Member")
		tambahPenyewa()
	}

	fmt.Println("Sukses Booking Lapangan")

	menuLain(mainReservation)
}

func tampilkanReservasi() {
	var n int
	fmt.Println("Pilih Opsi:")
	fmt.Println("1. Tiap Lapangan")
	fmt.Println("2. Semua Lapangan")
	fmt.Print("Pilih: ")
	fmt.Scan(&n)

	if n == 1 {
		displayLapName()

		fmt.Print("Pilih Lapangan: ")
		fmt.Scan(&n)
	
		for i := 0; i < len(database); i++ {
			if database[i].lapangan == lapangan[n-1].nama {
				loopReservasi(i)
			}
		}
		menuLain(mainReservation)
	} else if n == 2 {
		for i := 0; i < len(database); i++ {
			loopReservasi(i)
		}
		menuLain(mainReservation)
	} else {
		fmt.Println("Perintah Tidak Valid")
		menuLain(mainReservation)
	}

}

// Catatan Reservasi
// Tampilkan Data Reservasi Per Lapangan (Done)
// Tampilkan Data Reservasi Semua Lapangan (Done)

// Catatan Booking
// Booking Jam an (Done)
