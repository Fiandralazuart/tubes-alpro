package main

import "fmt"

func menuPenyewa() {
	var n int

	fmt.Println("=== MENU MANAJEMEN PENYEWA ===")
	fmt.Println("Pilih Menu")
	fmt.Println("1. Tampilkan Penyewa")
	fmt.Println("2. Tambah Penyewa")
	fmt.Println("3. Update Penyewa")
	fmt.Println("4. Hapus Penyewa")
	fmt.Println("5. Sorting Total Booking")
	fmt.Println("6. Sequential Search Penyewa")
	fmt.Println("7. Binary Search Penyewa")
	fmt.Println("8. Statistik Penyewa")
	fmt.Println("9. Kembali")
	fmt.Print("Pilih: ")
	fmt.Scan(&n)

	switch {
	case n == 1:
		tampilkanPenyewa()
	case n == 2:
		tambahPenyewa()
	case n == 3:
		updatePenyewa()
	case n == 4:
		hapusPenyewa()
	case n == 5:
		insertionSortBooking()
	case n == 6:
		menuSeqPenyewa()
	case n == 7:
		menuBinPenyewa()
	case n == 8:
		statistikPenyewa()
	case n == 9:
		main()
	default:
		fmt.Println("Perintah Tidak Valid")
		menuLain(menuPenyewa)
	}
}

func displayPenyewa(data []Penyewa, all bool, n int, label string) {
	if label != "" {
		fmt.Printf("=== %s ===\n", label)
	}

	if all {
		for i, p := range data {
			fmt.Printf("%d. Nama: %s | No HP: %s | Total Booking: %d | Total Pengeluaran: Rp.%d\n",
				i+1, p.nama, p.noHP, p.totalBooking, p.totalPengeluaran)
		}
	} else {
		if n < 1 || n > len(data) {
			fmt.Println("Nomor Tidak Valid")
			return
		}
		p := data[n-1]
		fmt.Printf("Nama: %s | No HP: %s | Total Booking: %d | Total Pengeluaran: Rp.%d\n",
			p.nama, p.noHP, p.totalBooking, p.totalPengeluaran)
	}

	fmt.Println()
}


func tampilkanPenyewa() {
	fmt.Println("=== TAMPILKAN PENYEWA ===")

	if len(penyewa) == 0 {
		fmt.Println("Data Penyewa Kosong")
		menuLain(menuPenyewa)
		return
	}

	displayPenyewa(penyewa, true, 0, "")
	menuLain(menuPenyewa)
}

func tambahPenyewa() {
	var nama, noHP string

	fmt.Println("=== MENU TAMBAH PENYEWA ===")

	fmt.Print("Masukkan Nama: ")
	fmt.Scan(&nama)

	fmt.Print("Masukkan No HP: ")
	fmt.Scan(&noHP)

	idBaru := 1
	if len(penyewa) > 0 {
		idBaru = penyewa[len(penyewa)-1].ID + 1
	}

	penyewa = append(penyewa, Penyewa{
		ID:               idBaru,
		nama:             nama,
		noHP:             noHP,
		totalBooking:     0,
		totalPengeluaran: 0,
	})

	fmt.Printf("Berhasil Menambahkan Penyewa Baru: %s!\n\n", nama)
	displayPenyewa(penyewa, true, 0, "")
	menuLain(menuPenyewa)
}

func updatePenyewa() {
	var n int

	fmt.Println("=== MENU UPDATE PENYEWA ===")

	if len(penyewa) == 0 {
		fmt.Println("Data Penyewa Kosong")
		menuLain(menuPenyewa)
		return
	}

	displayPenyewa(penyewa, true, 0, "")

	fmt.Print("Pilih Penyewa Untuk Update: ")
	fmt.Scan(&n)

	if n < 1 || n > len(penyewa) {
		fmt.Println("Nomor Tidak Valid")
		menuLain(menuPenyewa)
		return
	}

	displayPenyewa(penyewa, false, n, "Update Penyewa")

	isAll := ""
	fmt.Print("Update semua data? (yes/no): ")
	fmt.Scan(&isAll)

	if isAll == "yes" || isAll == "y" {
		var namaBaru, noHPBaru string

		fmt.Print("Ubah Nama: ")
		fmt.Scan(&namaBaru)

		fmt.Print("Ubah No HP: ")
		fmt.Scan(&noHPBaru)

		penyewa[n-1].nama = namaBaru
		penyewa[n-1].noHP = noHPBaru

		fmt.Println("Berhasil Update Semua Data")

	} else if isAll == "no" || isAll == "n" {
		ubah := ""
		fmt.Print("Masukkan field yang ingin diubah (nama/noHP): ")
		fmt.Scan(&ubah)

		switch {
		case ubah == "nama":
			var namaBaru string
			fmt.Print("Masukkan Nama Baru: ")
			fmt.Scan(&namaBaru)
			penyewa[n-1].nama = namaBaru
			fmt.Println("Berhasil Update Nama")

		case ubah == "noHP":
			var noHPBaru string
			fmt.Print("Masukkan No HP Baru: ")
			fmt.Scan(&noHPBaru)
			penyewa[n-1].noHP = noHPBaru
			fmt.Println("Berhasil Update No HP")

		default:
			fmt.Println("Field Tidak Valid")
		}

	} else {
		fmt.Println("Perintah Tidak Valid")
		menuLain(menuPenyewa)
		return
	}

	fmt.Println()
	displayPenyewa(penyewa, false, n, "Data Penyewa Terbaru")
	menuLain(menuPenyewa)
}

func hapusPenyewa() {
	var n int

	fmt.Println("=== MENU HAPUS PENYEWA ===")

	if len(penyewa) == 0 {
		fmt.Println("Data Penyewa Kosong")
		menuLain(menuPenyewa)
		return
	}

	fmt.Println("Penyewa Tersedia:")
	displayPenyewa(penyewa, true, 0, "")

	fmt.Print("Pilih penyewa untuk dihapus: ")
	fmt.Scan(&n)

	if n < 1 || n > len(penyewa) {
		fmt.Println("Nomor Tidak Valid")
		menuLain(menuPenyewa)
		return
	}

	displayPenyewa(penyewa, false, n, "Data Penyewa")

	cond := ""
	fmt.Print("Anda yakin menghapusnya? (yes/no): ")
	fmt.Scan(&cond)

	if cond == "yes" || cond == "y" {
		penyewa = append(penyewa[:n-1], penyewa[n:]...)
		fmt.Println("Berhasil Hapus Penyewa")
	} else if cond == "no" || cond == "n" {
		fmt.Println("Hapus Penyewa Dibatalkan")
	} else {
		fmt.Println("Perintah Tidak Valid")
	}

	menuLain(menuPenyewa)
}

func insertionSortBooking() {
	// melakukan insertion sort dari index 1 hingga akhir slice 
	for i := 1; i < len(penyewa); i++ {
		temp := penyewa[i]
		j := i - 1

		// menggeser elemen ke kanan selama totalBooking lebih kecil dari temp 
		for j >= 0 && penyewa[j].totalBooking < temp.totalBooking {
			penyewa[j+1] = penyewa[j]
			j--
		}
		penyewa[j+1] = temp
	}

	fmt.Println("Berhasil Sorting Berdasarkan Total Booking")
	fmt.Println()
	displayPenyewa(penyewa, true, 0, "")
	menuLain(menuPenyewa)
}

func updateStatsPenyewa(nama string, total int) {
	for i := 0; i < len(penyewa); i++ {
		if penyewa[i].nama == nama {
			penyewa[i].totalBooking++
			penyewa[i].totalPengeluaran += total
			break
		}
	}
}

func statistikPenyewa() {
	fmt.Println("=== STATISTIK PENYEWA ===")

	if len(penyewa) == 0 {
		fmt.Println("Data Penyewa Kosong")
		menuLain(menuPenyewa)
		return
	}

	idxTertinggi := 0
	for i := 1; i < len(penyewa); i++ {
		if penyewa[i].totalPengeluaran > penyewa[idxTertinggi].totalPengeluaran {
			idxTertinggi = i
		}
	}

	idxTerbanyak := 0
	for i := 1; i < len(penyewa); i++ {
		if penyewa[i].totalBooking > penyewa[idxTerbanyak].totalBooking {
			idxTerbanyak = i
		}
	}

	totalKeseluruhan := 0
	for i := 0; i < len(penyewa); i++ {
		totalKeseluruhan += penyewa[i].totalPengeluaran
	}

	fmt.Println("------------------")
	fmt.Printf("Total Penyewa Terdaftar : %d\n", len(penyewa))
	fmt.Printf("Total Pengeluaran Semua : Rp.%d\n", totalKeseluruhan)
	fmt.Println("------------------")
	fmt.Printf("Penyewa Teraktif        : %s (%dx booking)\n",
		penyewa[idxTerbanyak].nama, penyewa[idxTerbanyak].totalBooking)
	fmt.Printf("Penyewa Terbesar        : %s (Rp.%d)\n",
		penyewa[idxTertinggi].nama, penyewa[idxTertinggi].totalPengeluaran)
	fmt.Println("------------------")
	fmt.Println()

	displayPenyewa(penyewa, true, 0, "")
	menuLain(menuPenyewa)
}