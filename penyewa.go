package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Penyewa struct {
	ID           int
	nama         string
	noHP         string
	totalBooking int
	punyaVoucher bool
}

var penyewa = []Penyewa{
	{
		ID:           1,
		nama:         "Fathir",
		noHP:         "08123456789",
		totalBooking: 3,
		punyaVoucher: false,
	},
	{
		ID:           2,
		nama:         "Budi",
		noHP:         "08234567890",
		totalBooking: 7,
		punyaVoucher: true,
	},
	{
		ID:           3,
		nama:         "Andi",
		noHP:         "08345678901",
		totalBooking: 2,
		punyaVoucher: false,
	},
}

func mainPenyewa() {
	runPenyewa()
}

func runPenyewa() {
	menuPenyewa()
}

func menuPenyewa() {
	var n int

	fmt.Println("=== MENU MANAJEMEN PENYEWA ===")
	fmt.Println("Pilih Menu")
	fmt.Println("1. Tampilkan Penyewa")
	fmt.Println("2. Tambah Penyewa")
	fmt.Println("3. Update Penyewa")
	fmt.Println("4. Hapus Penyewa")
	fmt.Println("5. Sorting Total Booking")
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
	default:
		fmt.Println("Perintah Tidak Valid")
		menuLain(menuPenyewa)
	}
}

// displayPenyewa menampilkan data penyewa, jika all=true tampilkan semua, jika false tampilkan satu berdasarkan nomor urut
func displayPenyewa(data []Penyewa, all bool, n int, label string) {
	if label != "" {
		fmt.Printf("=== %s ===\n", label)
	}

	if all {
		// menampilkan seluruh data penyewa 
		for i, p := range data {
			voucher := "Tidak Ada"
			if p.punyaVoucher {
				voucher = "Ada"
			}
			fmt.Printf("%d. Nama: %s | No HP: %s | Total Booking: %d | Voucher: %s\n",
				i+1, p.nama, p.noHP, p.totalBooking, voucher)
		}
	} else {

		if n < 1 || n > len(data) {
			fmt.Println("Nomor Tidak Valid")
			return
		}
		p := data[n-1]
		voucher := "Tidak Ada"
		if p.punyaVoucher {
			voucher = "Ada"
		}
		fmt.Printf("Nama: %s | No HP: %s | Total Booking: %d | Voucher: %s\n",
			p.nama, p.noHP, p.totalBooking, voucher)
	}

	fmt.Println()
}

// menuLain mengarahkan kembali ke menu yang diberikan atau keluar jika ditolak
func menuLain(callback func()) {
	var pilihan string
	fmt.Print("Kembali ke menu? (yes/no): ")
	fmt.Scan(&pilihan)
	fmt.Println()

	switch strings.ToLower(pilihan) {
	case "yes", "y":
		callback()
	case "no", "n":
		fmt.Println("Terima kasih!")
		os.Exit(0)
	default:
		fmt.Println("Perintah Tidak Valid")
		menuLain(callback)
	}
}

// tampilkanPenyewa menampilkan seluruh data penyewa
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

// tambahPenyewa menambahkan data penyewa baru ke dalam slice 
func tambahPenyewa() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== MENU TAMBAH PENYEWA ===")

	fmt.Print("Masukkan Nama: ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama) // trim newline dan spasi di kiri kanan (gabriel edbert)

	fmt.Print("Masukkan No HP: ")
	var noHP string
	fmt.Scanln(&noHP)

	idBaru := 1
	if len(penyewa) > 0 {
		idBaru = penyewa[len(penyewa)-1].ID + 1
	}

	penyewa = append(penyewa, Penyewa{
		ID:           idBaru,
		nama:         nama,
		noHP:         noHP,
		totalBooking: 0,
		punyaVoucher: false,
	})

	fmt.Printf("Berhasil Menambahkan Penyewa Baru: %s!\n\n", nama)
	displayPenyewa(penyewa, true, 0, "")
	menuLain(menuPenyewa)
}

// updatePenyewa mengupdate data penyewa
func updatePenyewa() {
	var n int
	reader := bufio.NewReader(os.Stdin)

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
	reader.ReadString('\n') // membersihkan sisa newline di buffer sebelum ReadString berikutnya 

	if isAll == "yes" || isAll == "y" {
		fmt.Print("Ubah Nama: ")
		namaBaru, _ := reader.ReadString('\n')
		namaBaru = strings.TrimSpace(namaBaru)

		fmt.Print("Ubah No HP: ")
		var noHPBaru string
		fmt.Scanln(&noHPBaru)

		penyewa[n-1].nama = namaBaru
		penyewa[n-1].noHP = noHPBaru

		fmt.Println("Berhasil Update Semua Data")

	} else if isAll == "no" || isAll == "n" {
		ubah := ""
		fmt.Print("Masukkan field yang ingin diubah (nama/noHP): ")
		fmt.Scan(&ubah)
		reader.ReadString('\n') // membersihkan sisa newline di buffer 

		switch {
		case ubah == "nama":
			fmt.Print("Masukkan Nama Baru: ")
			namaBaru, _ := reader.ReadString('\n')
			namaBaru = strings.TrimSpace(namaBaru)
			penyewa[n-1].nama = namaBaru
			fmt.Println("Berhasil Update Nama")

		case ubah == "noHP":
			fmt.Print("Masukkan No HP Baru: ")
			var noHPBaru string
			fmt.Scanln(&noHPBaru)
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

// hapusPenyewa menghapus data penyewa berdasarkan nomor urut 
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

	// validasi nomor urut yang dipilih 
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
		// menghapus elemen dari slice dengan cara menggabungkan elemen sebelum dan sesudahnya 
		penyewa = append(penyewa[:n-1], penyewa[n:]...)
		fmt.Println("Berhasil Hapus Penyewa")
	} else if cond == "no" || cond == "n" {
		fmt.Println("Hapus Penyewa Dibatalkan")
	} else {
		fmt.Println("Perintah Tidak Valid")
	}

	menuLain(menuPenyewa)
}

// insertionSortBooking mengurutkan slice penyewa berdasarkan totalBooking dari besar ke kecil 
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