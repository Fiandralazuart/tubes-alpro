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
	fmt.Println("6. Statistik Pendapatan")
	fmt.Println("7. Kembali")
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
		statistikPendapatan()
		// fmt.Println("helo")
	case n == 7:
		main()
	default:
		fmt.Println("Perintah Tidak Valid")
		menuLain(mainReservation)
	}
}

func tampilkanJadwal() {
	var n int
	var pilih string
	var metode, kriteria int

	fmt.Println("=== Jadwal Lapangan ===")
	fmt.Println("Pilih Lapangan: ")
	displayLapName()

	fmt.Print("Pilih: ")
	fmt.Scan(&n)

	fmt.Println("Urutkan Data? (yes/no)")
	fmt.Scan(&pilih)

	if pilih == "yes" {
		fmt.Println("Pilih Kriteria Sorting")
		fmt.Println("1. Status Booking")
		fmt.Println("2. Harga")
		fmt.Scan(&kriteria)

		if kriteria < 1 || kriteria > 2 {
			fmt.Println("Input Tidak Valid")
			menuLain(mainReservation)
		}
		
		fmt.Println("Pilih Metode Sorting")
		fmt.Println("1. SelectionSort")
		fmt.Println("2. InsertionSort")
		fmt.Scan(&metode)

		if metode < 1 || metode > 2  {
			fmt.Println("Input Tidak Valid")
			menuLain(mainReservation)
		}

	} else if pilih != "no" {
		fmt.Println("Perintah Tidak Valid")
		menuLain(mainReservation)
		return
	}

	for i := 0; i < len(database); i++ {
		if database[i].lapangan == lapangan[n-1].nama {

			jadwalTampil := append([]Jam{}, database[i].jadwal...)

			if pilih == "yes" {
				if kriteria == 1 && metode == 1 {
					selectionSort(jadwalTampil)
				} else if kriteria == 1 && metode == 2 {
					insertionSort(jadwalTampil)
				} else if kriteria == 2 && metode == 1 {
					selectionSortHarga(jadwalTampil)
				} else if kriteria == 2 && metode == 2 {
					insertionSortHarga(jadwalTampil)
				}
			}

			fmt.Println("==== Jadwal Tersedia ====")
			fmt.Printf("Lap. %s | %d-%d-%d YYYY-M-D\n",
				database[i].lapangan,
				database[i].tanggal.tahun,
				database[i].tanggal.bulan,
				database[i].tanggal.hari)

			fmt.Println()
			displayJadwal(jadwalTampil)

			fmt.Println()
		}
	}

	menuLain(mainReservation)
}

func buatJadwal() {
	var n int
	var hari, bulan, tahun int
	fmt.Println("=== Buat Jadwal ===")
	fmt.Println("Pilih Lapangan: ")
	displayLapName()

	fmt.Println("Pilih: ")
	fmt.Scan(&n)

	x := 1
	fmt.Println("Masukkan Tanggal: (numerik | D-M-YYYY)")
	fmt.Print("tanggal: ")
	for x != 0 {
		fmt.Scan(&hari)
		x = validateDate(hari, 30)
		if x == 0 {
			fmt.Println("Input tidak valid")
			return
		}

		fmt.Print("Bulan: ")
		fmt.Scan(&bulan)
		x = validateDate(bulan, 12)
		if x == 0 {
			fmt.Println("Input tidak valid")
			return
		}
		
		fmt.Print("Tahun: ")
		fmt.Scan(&tahun)
		x = validateDate(tahun, 3000)
		if x == 0 {
			fmt.Println("Input tidak valid")
			return
		}
	}

	isAvailable := true

	for i := 0; i < len(database); i++ {
		if database[i].tanggal.hari == hari &&
			database[i].tanggal.bulan == bulan &&
			database[i].tanggal.tahun == tahun &&
			database[i].lapangan == lapangan[n-1].nama {

			isAvailable = false
			break
		}
	}

	if isAvailable {
		database = append(database, Database{
			tanggal: tanggal{
				hari,
				bulan,
				tahun,
			},
			lapangan: lapangan[n-1].nama,
			harga: harga{
				happyHour:    lapangan[n-1].harga.happyHour,
				hargaDefault: lapangan[n-1].harga.hargaDefault,
			},
			jadwal: defaultJadwal(lapangan[n-1].harga.hargaDefault, lapangan[n-1].harga.happyHour),
		})
		fmt.Println("Sukses Membuat Jadwal")
		menuLain(mainReservation)
	} else {
		fmt.Println("Gagal: Jadwal Sudah Ada")
		menuLain(mainReservation)
	}

}

func bookingJadwal() {
	var n, idx, hari, bulan, tahun int
	var nama string
	// var data Database
	fmt.Println("=== Booking Jadwal ===")
	jml := displayLapName()

	fmt.Println("Pilih: ")
	fmt.Scan(&n)

	if n > jml {
		fmt.Println("Perintah Tidak Valid")
		menuLain(mainReservation)
	}

	x := 1
	fmt.Println("Masukkan Tanggal: (numerik | D-M-YYYY)")
	fmt.Print("tanggal: ")
	for x != 0 {
		fmt.Scan(&hari)
		x = validateDate(hari, 30)
		if x == 0 {
			fmt.Println("Input tidak valid")
			return
		}

		fmt.Print("Bulan: ")
		fmt.Scan(&bulan)
		x = validateDate(bulan, 12)
		if x == 0 {
			fmt.Println("Input tidak valid")
			return
		}
		
		fmt.Print("Tahun: ")
		fmt.Scan(&tahun)
		x = validateDate(tahun, 3000)
		if x == 0 {
			fmt.Println("Input tidak valid")
			return
		}
		break
	}


	isThere := false
	for i := 0; i < len(database); i++ {
		if database[i].lapangan == lapangan[n-1].nama && database[i].tanggal.hari == hari && database[i].tanggal.bulan == bulan && database[i].tanggal.tahun == tahun {
			// data = database[i]
			idx = i
			fmt.Println("==== Jadwal Tersedia ====")
			fmt.Printf("Lap. %s | %d-%d-%d YYYY-M-D\n", database[i].lapangan, database[i].tanggal.tahun, database[i].tanggal.bulan, database[i].tanggal.hari)
			fmt.Println("")
			displayJadwal(database[i].jadwal)
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

		var x, totalHarga int
		for i := 0; i < jmlJam; i++ {
			fmt.Println("Pilih Jadwal: ")
			fmt.Scan(&x)

			database[idx].jadwal[x-1].isAvailable = false
			totalHarga += database[idx].jadwal[x-1].harga
		}
		
		tgl := fmt.Sprintf("%d-%d-%d", tahun, bulan, hari)

		database[idx].reservasi = append(database[idx].reservasi, Sewa{
			penyewa:  nama,
			jamMulai: ambilJam(x - 1),
			jamAkhir: ambilJam(x),
			tglMulai: tgl,
			tglAkhir: tgl,
			durasi:   n,
			total: totalHarga,
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


func statistikPendapatan() {
	fmt.Println("=== Dashboard Pendapatan ===")
	var dataStats []stats

	for i := 0; i < len(database); i++ {
		ketemu := false

		for j := 0; j < len(dataStats); j++ {
			if database[i].tanggal.bulan == dataStats[j].bulan {
				ketemu = true
				break
			}
		}

		if !ketemu {
			dataStats = append(dataStats, stats{
				bulan: database[i].tanggal.bulan,
				reservasi: 0,
				jam: 0,
				total: 0,
				hari: [14]int{},
			})
		}
	}

	for i := 0; i < len(dataStats); i++ {
		for j := 0; j < len(database); j++ {
			if dataStats[i].bulan == database[j].tanggal.bulan {
				for k := 0; k < len(database[j].reservasi); k++{
					dataStats[i].reservasi++
					dataStats[i].jam += database[j].reservasi[k].durasi
					dataStats[i].total += database[j].reservasi[k].total
					x := ambilIndexJam(database[j].reservasi[k].jamMulai)
					dataStats[i].hari[x] += 1
				}
			}
		}
		maxIdx := 0
		for k := 1; k < len(dataStats[i].hari); k++ {
			if dataStats[i].hari[k] > dataStats[i].hari[maxIdx] {
				maxIdx = k
			}
		}
		fmt.Println("------------------")
		fmt.Println("Bulan:", dataStats[i].bulan)
		fmt.Println("Total Reservasi:", dataStats[i].reservasi)
		fmt.Println("Total Jam :", dataStats[i].jam)
		fmt.Printf("Total Revenue: Rp.%d\n", dataStats[i].total)
		fmt.Printf("Jam Teramai: %s\n", ambilJam(maxIdx))
		fmt.Println("------------------")
	}

	menuLain(mainReservation)
}

