package main

import (
	"bufio"
	"fmt"
	"os"
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

// func mainPenyewa() {
// 	runPenyewa()
// }

// func runPenyewa() {
// 	menuPenyewa()
// }

func menuPenyewa() {
	var pilih int

	fmt.Println("")
	fmt.Println("=== MENU PENYEWA ===")
	fmt.Println("1. Tambah Penyewa")
	fmt.Println("2. Tampilkan Penyewa")
	fmt.Println("3. Update Penyewa")
	fmt.Println("4. Hapus Penyewa")
	fmt.Println("5. Sorting Total Booking")
	fmt.Println("6. Keluar")
	fmt.Print("Pilih Menu: ")
	fmt.Scanln(&pilih)

	switch pilih {
	case 1:
		tambahPenyewa()
	case 2:
		tampilkanPenyewa()
	case 3:
		updatePenyewa()
	case 4:
		hapusPenyewa()
	case 5:
		insertionSortBooking()
	case 6:
		fmt.Println("Program Selesai")
	default:
		fmt.Println("Menu Tidak Valid")
		menuPenyewa()
	}
}

func tampilkanPenyewa() {
	fmt.Println("")
	fmt.Println("=== DATA PENYEWA ===")

	if len(penyewa) == 0 {
		fmt.Println("Data Penyewa Kosong")
		menuPenyewa()
		return
	}

	for i := 0; i < len(penyewa); i++ {
		fmt.Printf("ID               : %d\n", penyewa[i].ID)
		fmt.Printf("Nama             : %s\n", penyewa[i].nama)
		fmt.Printf("No HP            : %s\n", penyewa[i].noHP)
		fmt.Printf("Total Booking    : %d\n", penyewa[i].totalBooking)

		if penyewa[i].punyaVoucher {
			fmt.Println("Voucher          : Ada")
		} else {
			fmt.Println("Voucher          : Tidak Ada")
		}

		fmt.Println("")
	}

	menuPenyewa()
}

func tambahPenyewa() {
	var nama string
	var noHP string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("")
	fmt.Println("=== TAMBAH PENYEWA ===")

	fmt.Print("Masukkan Nama : ")
	nama, _ = reader.ReadString('\n')

	nama = nama[:len(nama)-1]

	fmt.Print("Masukkan No HP : ")
	fmt.Scanln(&noHP)

	idBaru := 1

	if len(penyewa) > 0 {
		idBaru = penyewa[len(penyewa)-1].ID + 1
	}

	dataBaru := Penyewa{
		ID:           idBaru,
		nama:         nama,
		noHP:         noHP,
		totalBooking: 0,
		punyaVoucher: false,
	}

	penyewa = append(penyewa, dataBaru)

	fmt.Println("")
	fmt.Println("Berhasil Menambahkan Penyewa")
	fmt.Println("")

	menuPenyewa()
}

func cariIndexPenyewa(id int) int {
	for i := 0; i < len(penyewa); i++ {
		if penyewa[i].ID == id {
			return i
		}
	}

	return -1
}

func updatePenyewa() {
	var id int
	var namaBaru string
	var noHPBaru string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("")
	fmt.Println("=== UPDATE PENYEWA ===")

	tampilkanDataSingkat()

	fmt.Print("Masukkan ID Penyewa : ")
	fmt.Scanln(&id)

	index := cariIndexPenyewa(id)

	if index == -1 {
		fmt.Println("ID Tidak Ditemukan")
		menuPenyewa()
		return
	}

	fmt.Print("Masukkan Nama Baru : ")
	namaBaru, _ = reader.ReadString('\n')

	namaBaru = namaBaru[:len(namaBaru)-1]

	fmt.Print("Masukkan No HP Baru : ")
	fmt.Scanln(&noHPBaru)

	penyewa[index].nama = namaBaru
	penyewa[index].noHP = noHPBaru

	fmt.Println("")
	fmt.Println("Berhasil Update Penyewa")
	fmt.Println("")

	menuPenyewa()
}

func hapusPenyewa() {
	var id int

	fmt.Println("")
	fmt.Println("=== HAPUS PENYEWA ===")

	tampilkanDataSingkat()

	fmt.Print("Masukkan ID Penyewa : ")
	fmt.Scanln(&id)

	index := cariIndexPenyewa(id)

	if index == -1 {
		fmt.Println("ID Tidak Ditemukan")
		menuPenyewa()
		return
	}

	penyewa = append(penyewa[:index], penyewa[index+1:]...)

	fmt.Println("")
	fmt.Println("Berhasil Menghapus Penyewa")
	fmt.Println("")

	menuPenyewa()
}

func tampilkanDataSingkat() {
	fmt.Println("")

	for i := 0; i < len(penyewa); i++ {
		fmt.Printf("%d. %s\n", penyewa[i].ID, penyewa[i].nama)
	}

	fmt.Println("")
}

func insertionSortBooking() {
	var temp Penyewa
	var j int

	for i := 1; i < len(penyewa); i++ {
		temp = penyewa[i]
		j = i - 1

		for j >= 0 && penyewa[j].totalBooking < temp.totalBooking {
			penyewa[j+1] = penyewa[j]
			j--
		}

		penyewa[j+1] = temp
	}

	fmt.Println("")
	fmt.Println("Berhasil Sorting Berdasarkan Total Booking")

	tampilkanPenyewa()
}
