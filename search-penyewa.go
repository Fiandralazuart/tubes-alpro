package main

import (
	"fmt"
	"strings"
)

// ==================== MENU UTAMA SEARCH ====================

// menampilkan pilihan antara sequential search dan binary search
func menuSearchPenyewa() {
	var n int

	fmt.Println("=== MENU PENCARIAN PENYEWA ===")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search")
	fmt.Print("Pilih: ")
	fmt.Scan(&n)

	switch {
	case n == 1:
		menuSeqPenyewa()
	case n == 2:
		menuBinPenyewa()
	default:
		fmt.Println("Perintah Tidak Valid")
		menuLain(menuSearchPenyewa)
	}
}

// ==================== SEQUENTIAL SEARCH ====================

// menampilkan pilihan field untuk sequential search
func menuSeqPenyewa() {
	var n int

	fmt.Println("=== SEQUENTIAL SEARCH PENYEWA ===")
	fmt.Println("Cari berdasarkan:")
	fmt.Println("1. Nama")
	fmt.Println("2. No HP")
	fmt.Println("3. Status Voucher")
	fmt.Print("Pilih: ")
	fmt.Scan(&n)

	switch {
	case n == 1:
		seqCariNama()
	case n == 2:
		seqCariNoHP()
	case n == 3:
		seqCariVoucher()
	default:
		fmt.Println("Perintah Tidak Valid")
		menuLain(menuSeqPenyewa)
	}
}

// mencari penyewa berdasarkan nama secara linear
func seqCariNama() {
	var key string

	fmt.Print("Masukkan Nama: ")
	fmt.Scan(&key)

	keyLower := strings.ToLower(key)

	// melakukan perulangan untuk mencari penyewa yang namanya mengandung keyword
	var hasil []Penyewa
	for i := 0; i < len(penyewa); i++ {
		if strings.Contains(strings.ToLower(penyewa[i].nama), keyLower) {
			hasil = append(hasil, penyewa[i])
		}
	}

	fmt.Println("")
	if len(hasil) == 0 {
		fmt.Printf("Penyewa dengan nama \"%s\" tidak ditemukan.\n", key)
		menuLain(menuSeqPenyewa)
		return
	}

	fmt.Printf("Ditemukan %d penyewa:\n", len(hasil))
	displayPenyewa(hasil, true, 0, "Hasil Sequential Search")

	// jika hanya 1 hasil ditemukan, tawarkan aksi langsung ke owner
	if len(hasil) == 1 {
		aksiSetelahDitemukan(hasil[0].ID)
	} else {
		menuLain(menuSeqPenyewa)
	}
}

// mencari penyewa berdasarkan nomor HP secara linear
func seqCariNoHP() {
	var key string

	fmt.Print("Masukkan No HP: ")
	fmt.Scan(&key)

	// melakukan perulangan untuk mencari penyewa yang nomornya mengandung keyword
	var hasil []Penyewa
	for i := 0; i < len(penyewa); i++ {
		if strings.Contains(penyewa[i].noHP, key) {
			hasil = append(hasil, penyewa[i])
		}
	}

	fmt.Println("")
	if len(hasil) == 0 {
		fmt.Printf("Penyewa dengan No HP \"%s\" tidak ditemukan.\n", key)
		menuLain(menuSeqPenyewa)
		return
	}

	fmt.Printf("Ditemukan %d penyewa:\n", len(hasil))
	displayPenyewa(hasil, true, 0, "Hasil Sequential Search")

	// jika hanya 1 hasil ditemukan, tawarkan aksi langsung ke owner
	if len(hasil) == 1 {
		aksiSetelahDitemukan(hasil[0].ID)
	} else {
		menuLain(menuSeqPenyewa)
	}
}

// mencari penyewa berdasarkan status kepemilikan voucher secara linear
func seqCariVoucher() {
	var n int

	fmt.Println("Tampilkan penyewa:")
	fmt.Println("1. Yang punya voucher")
	fmt.Println("2. Yang tidak punya voucher")
	fmt.Print("Pilih: ")
	fmt.Scan(&n)

	if n != 1 && n != 2 {
		fmt.Println("Perintah Tidak Valid")
		menuLain(menuSeqPenyewa)
		return
	}

	cariVoucher := n == 1

	// melakukan perulangan untuk memfilter penyewa berdasarkan status voucher
	var hasil []Penyewa
	for i := 0; i < len(penyewa); i++ {
		if penyewa[i].punyaVoucher == cariVoucher {
			hasil = append(hasil, penyewa[i])
		}
	}

	statusLabel := "tidak punya voucher"
	if cariVoucher {
		statusLabel = "punya voucher"
	}

	fmt.Println("")
	if len(hasil) == 0 {
		fmt.Printf("Tidak ada penyewa yang %s.\n", statusLabel)
	} else {
		fmt.Printf("Ditemukan %d penyewa yang %s:\n", len(hasil), statusLabel)
		displayPenyewa(hasil, true, 0, "Hasil Sequential Search - Voucher")
	}

	menuLain(menuSeqPenyewa)
}

// ==================== BINARY SEARCH ====================

// menampilkan pilihan field untuk binary search
func menuBinPenyewa() {
	var n int

	fmt.Println("=== BINARY SEARCH PENYEWA ===")
	fmt.Println("Cari berdasarkan (exact match):")
	fmt.Println("1. Nama")
	fmt.Println("2. No HP")
	fmt.Println("3. ID")
	fmt.Print("Pilih: ")
	fmt.Scan(&n)

	switch {
	case n == 1:
		binCariNama()
	case n == 2:
		binCariNoHP()
	case n == 3:
		binCariID()
	default:
		fmt.Println("Perintah Tidak Valid")
		menuLain(menuBinPenyewa)
	}
}

// mencari penyewa berdasarkan nama dengan binary search
func binCariNama() {
	var key string

	fmt.Print("Masukkan Nama (exact): ")
	fmt.Scan(&key)

	// menyalin slice agar data asli tidak berubah saat diurutkan
	data := make([]Penyewa, len(penyewa))
	copy(data, penyewa)

	sortNama(data)
	idx := binNama(data, key)

	fmt.Println("")
	if idx == -1 {
		fmt.Printf("Penyewa dengan nama \"%s\" tidak ditemukan.\n", key)
		menuLain(menuBinPenyewa)
		return
	}

	fmt.Println("Penyewa ditemukan:")
	displayPenyewa(data, false, idx+1, "Hasil Binary Search")

	// setelah ditemukan, tawarkan aksi langsung ke owner
	aksiSetelahDitemukan(data[idx].ID)
}

// mencari penyewa berdasarkan nomor HP dengan binary search
func binCariNoHP() {
	var key string

	fmt.Print("Masukkan No HP (exact): ")
	fmt.Scan(&key)

	// menyalin slice agar data asli tidak berubah saat diurutkan
	data := make([]Penyewa, len(penyewa))
	copy(data, penyewa)

	sortNoHP(data)
	idx := binNoHP(data, key)

	fmt.Println("")
	if idx == -1 {
		fmt.Printf("Penyewa dengan No HP \"%s\" tidak ditemukan.\n", key)
		menuLain(menuBinPenyewa)
		return
	}

	fmt.Println("Penyewa ditemukan:")
	displayPenyewa(data, false, idx+1, "Hasil Binary Search")

	// setelah ditemukan, tawarkan aksi langsung ke owner
	aksiSetelahDitemukan(data[idx].ID)
}

// mencari penyewa berdasarkan ID dengan binary search
func binCariID() {
	var targetID int

	fmt.Print("Masukkan ID: ")
	fmt.Scan(&targetID)

	// menyalin slice agar data asli tidak berubah saat diurutkan
	data := make([]Penyewa, len(penyewa))
	copy(data, penyewa)

	sortID(data)
	idx := binID(data, targetID)

	fmt.Println("")
	if idx == -1 {
		fmt.Printf("Penyewa dengan ID %d tidak ditemukan.\n", targetID)
		menuLain(menuBinPenyewa)
		return
	}

	fmt.Println("Penyewa ditemukan:")
	displayPenyewa(data, false, idx+1, "Hasil Binary Search")

	// setelah ditemukan, tawarkan aksi langsung ke owner
	aksiSetelahDitemukan(data[idx].ID)
}

// ==================== AKSI SETELAH DITEMUKAN ====================

// menawarkan aksi update atau hapus setelah penyewa berhasil ditemukan
func aksiSetelahDitemukan(id int) {
	var n int

	fmt.Println("=== AKSI ===")
	fmt.Println("1. Update Penyewa Ini")
	fmt.Println("2. Hapus Penyewa Ini")
	fmt.Println("3. Kembali")
	fmt.Print("Pilih: ")
	fmt.Scan(&n)

	// mencari nomor urut penyewa di slice asli berdasarkan ID
	nomorUrut := -1
	for i := 0; i < len(penyewa); i++ {
		if penyewa[i].ID == id {
			nomorUrut = i + 1
			break
		}
	}

	if nomorUrut == -1 {
		fmt.Println("Penyewa tidak ditemukan")
		menuLain(menuSearchPenyewa)
		return
	}

	switch {
	case n == 1:
		updatePenyewaByIndex(nomorUrut)
	case n == 2:
		hapusPenyewaByIndex(nomorUrut)
	case n == 3:
		menuSearchPenyewa()
	default:
		fmt.Println("Perintah Tidak Valid")
		menuLain(menuSearchPenyewa)
	}
}

// mengupdate penyewa berdasarkan nomor urut, dipanggil setelah search
func updatePenyewaByIndex(n int) {
	displayPenyewa(penyewa, false, n, "Update Penyewa")

	isAll := ""
	fmt.Print("Update semua data? (yes/no): ")
	fmt.Scan(&isAll)

	if isAll == "yes" {
		var nama, noHP string

		fmt.Print("Ubah Nama: ")
		fmt.Scan(&nama)
		fmt.Print("Ubah No HP: ")
		fmt.Scan(&noHP)

		penyewa[n-1].nama = nama
		penyewa[n-1].noHP = noHP

		displayPenyewa(penyewa, false, n, "Berhasil Update Penyewa")
		menuSearchPenyewa()

	} else if isAll == "no" {
		ubah := ""
		fmt.Print("Masukkan field yang ingin diubah (nama/noHP): ")
		fmt.Scan(&ubah)

		switch {
		case ubah == "nama":
			var nama string
			fmt.Print("Masukkan Nama Baru: ")
			fmt.Scan(&nama)
			penyewa[n-1].nama = nama

		case ubah == "noHP":
			var noHP string
			fmt.Print("Masukkan No HP Baru: ")
			fmt.Scan(&noHP)
			penyewa[n-1].noHP = noHP

		default:
			fmt.Println("Field Tidak Valid")
		}

		displayPenyewa(penyewa, false, n, "Berhasil Update Penyewa")
		menuLain(menuSearchPenyewa)

	} else {
		fmt.Println("Perintah Tidak Valid")
		menuLain(menuSearchPenyewa)
	}
}

// menghapus penyewa berdasarkan nomor urut, dipanggil setelah search
func hapusPenyewaByIndex(n int) {
	displayPenyewa(penyewa, false, n, "Data Penyewa")

	cond := ""
	fmt.Print("Anda yakin menghapusnya? (yes/no): ")
	fmt.Scan(&cond)

	if cond == "yes" {
		// menghapus elemen dari slice dengan cara menggabungkan elemen sebelum dan sesudahnya
		penyewa = append(penyewa[:n-1], penyewa[n:]...)
		fmt.Println("Berhasil Hapus Penyewa")
		menuLain(menuSearchPenyewa)
	} else if cond == "no" {
		menuLain(menuSearchPenyewa)
	} else {
		fmt.Println("Perintah Tidak Valid")
		menuLain(menuSearchPenyewa)
	}
}

// ==================== HELPER SORT ====================

// mengurutkan data penyewa berdasarkan nama A-Z menggunakan insertion sort
func sortNama(data []Penyewa) {
	for i := 1; i < len(data); i++ {
		temp := data[i]
		j := i - 1

		// menggeser elemen ke kanan selama nama lebih besar dari temp
		for j >= 0 && strings.ToLower(data[j].nama) > strings.ToLower(temp.nama) {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
}

//  mengurutkan data penyewa berdasarkan nomor HP menggunakan insertion sort
func sortNoHP(data []Penyewa) {
	for i := 1; i < len(data); i++ {
		temp := data[i]
		j := i - 1

		// menggeser elemen ke kanan selama noHP lebih besar dari temp
		for j >= 0 && data[j].noHP > temp.noHP {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
}

//  mengurutkan data penyewa berdasarkan ID menggunakan insertion sort
func sortID(data []Penyewa) {
	for i := 1; i < len(data); i++ {
		temp := data[i]
		j := i - 1

		// menggeser elemen ke kanan selama ID lebih besar dari temp
		for j >= 0 && data[j].ID > temp.ID {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
}

// ==================== HELPER BINARY SEARCH ====================

// mencari penyewa berdasarkan nama dengan binary search, data harus sudah terurut
func binNama(data []Penyewa, key string) int {
	awal := 0
	akhir := len(data) - 1
	key = strings.ToLower(key)

	// melakukan binary search dengan membagi rentang pencarian setiap iterasi
	for awal <= akhir {
		tengah := (awal + akhir) / 2
		namaTengah := strings.ToLower(data[tengah].nama)

		if namaTengah == key {
			return tengah
		} else if namaTengah < key {
			awal = tengah + 1
		} else {
			akhir = tengah - 1
		}
	}
	return -1
}

// mencari penyewa berdasarkan nomor HP dengan binary search, data harus sudah terurut
func binNoHP(data []Penyewa, key string) int {
	awal := 0
	akhir := len(data) - 1

	// melakukan binary search dengan membagi rentang pencarian setiap iterasi
	for awal <= akhir {
		tengah := (awal + akhir) / 2
		noTengah := data[tengah].noHP

		if noTengah == key {
			return tengah
		} else if noTengah < key {
			awal = tengah + 1
		} else {
			akhir = tengah - 1
		}
	}
	return -1
}

// mencari penyewa berdasarkan ID dengan binary search, data harus sudah terurut
func binID(data []Penyewa, targetID int) int {
	awal := 0
	akhir := len(data) - 1

	// melakukan binary search dengan membagi rentang pencarian setiap iterasi
	for awal <= akhir {
		tengah := (awal + akhir) / 2

		if data[tengah].ID == targetID {
			return tengah
		} else if data[tengah].ID < targetID {
			awal = tengah + 1
		} else {
			akhir = tengah - 1
		}
	}
	return -1
}