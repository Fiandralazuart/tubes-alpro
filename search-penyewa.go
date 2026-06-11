package main

import (
	"fmt"
	"strings"
)

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

func menuSeqPenyewa() {
	var n int

	fmt.Println("=== SEQUENTIAL SEARCH PENYEWA ===")
	fmt.Println("Cari berdasarkan:")
	fmt.Println("1. Nama")
	fmt.Println("2. No HP")
	fmt.Println("3. Total Pengeluaran Terbesar")
	fmt.Print("Pilih: ")
	fmt.Scan(&n)

	switch {
	case n == 1:
		seqCariNama()
	case n == 2:
		seqCariNoHP()
	case n == 3:
		seqCariPengeluaranTerbesar()
	default:
		fmt.Println("Perintah Tidak Valid")
		menuLain(menuSeqPenyewa)
	}
}

func seqCariNama() {
	var key string

	fmt.Print("Masukkan Nama: ")
	fmt.Scan(&key)

	keyLower := strings.ToLower(key)

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

	if len(hasil) == 1 {
		aksiSetelahDitemukan(hasil[0].ID)
	} else {
		menuLain(menuSeqPenyewa)
	}
}

func seqCariNoHP() {
	var key string

	fmt.Print("Masukkan No HP: ")
	fmt.Scan(&key)

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

	if len(hasil) == 1 {
		aksiSetelahDitemukan(hasil[0].ID)
	} else {
		menuLain(menuSeqPenyewa)
	}
}

func seqCariPengeluaranTerbesar() {
	if len(penyewa) == 0 {
		fmt.Println("Data Penyewa Kosong")
		menuLain(menuSeqPenyewa)
		return
	}

	idxTerbesar := 0
	for i := 1; i < len(penyewa); i++ {
		if penyewa[i].totalPengeluaran > penyewa[idxTerbesar].totalPengeluaran {
			idxTerbesar = i
		}
	}

	fmt.Println("")
	fmt.Printf("Penyewa dengan pengeluaran terbesar:\n")
	displayPenyewa(penyewa, false, idxTerbesar+1, "Hasil Sequential Search - Pengeluaran Terbesar")

	aksiSetelahDitemukan(penyewa[idxTerbesar].ID)
}

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

func binCariNama() {
	var key string

	fmt.Print("Masukkan Nama (exact): ")
	fmt.Scan(&key)

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
	aksiSetelahDitemukan(data[idx].ID)
}

func binCariNoHP() {
	var key string

	fmt.Print("Masukkan No HP (exact): ")
	fmt.Scan(&key)

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
	aksiSetelahDitemukan(data[idx].ID)
}

func binCariID() {
	var targetID int

	fmt.Print("Masukkan ID: ")
	fmt.Scan(&targetID)

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
	aksiSetelahDitemukan(data[idx].ID)
}

func aksiSetelahDitemukan(id int) {
	var n int

	fmt.Println("=== AKSI ===")
	fmt.Println("1. Update Penyewa Ini")
	fmt.Println("2. Hapus Penyewa Ini")
	fmt.Println("3. Kembali")
	fmt.Print("Pilih: ")
	fmt.Scan(&n)

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

func hapusPenyewaByIndex(n int) {
	displayPenyewa(penyewa, false, n, "Data Penyewa")

	cond := ""
	fmt.Print("Anda yakin menghapusnya? (yes/no): ")
	fmt.Scan(&cond)

	if cond == "yes" {
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

func sortNama(data []Penyewa) {
	for i := 1; i < len(data); i++ {
		temp := data[i]
		j := i - 1

		for j >= 0 && strings.ToLower(data[j].nama) > strings.ToLower(temp.nama) {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
}

func sortNoHP(data []Penyewa) {
	for i := 1; i < len(data); i++ {
		temp := data[i]
		j := i - 1

		for j >= 0 && data[j].noHP > temp.noHP {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
}

func sortID(data []Penyewa) {
	for i := 1; i < len(data); i++ {
		temp := data[i]
		j := i - 1

		for j >= 0 && data[j].ID > temp.ID {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
}

func binNama(data []Penyewa, key string) int {
	awal := 0
	akhir := len(data) - 1
	key = strings.ToLower(key)

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

func binNoHP(data []Penyewa, key string) int {
	awal := 0
	akhir := len(data) - 1

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

func binID(data []Penyewa, targetID int) int {
	awal := 0
	akhir := len(data) - 1

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