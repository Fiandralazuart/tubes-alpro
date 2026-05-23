package main
import "fmt"

func mainCrud() {

	var n int
	fmt.Println("=== MENU MANAJEMEN LAPANGAN ===")
	
	fmt.Println("Pilih Menu")
	fmt.Println("1. Tampilkan Lapangan")
	fmt.Println("2. Tambahkan Lapangan")
	fmt.Println("3. Edit Lapangan")
	fmt.Println("4. Hapus Lapangan")
	fmt.Println("5. Kembali")

	fmt.Print("Menu: ")
	fmt.Scan(&n)

	switch {
		case n == 1:
			tampilkanLapangan()
		case n == 2:
			tambahLapangan()
		case n == 3:
			updateLapangan()
		case n == 4:
			hapusLapangan()
		case n == 5:
			main()
		default:
			fmt.Println("Perintah Tidak Valid")
			menuLain(mainCrud)
	}

}

func tampilkanLapangan() {
	displayLap(lapangan, true, 0, "")
	
	menuLain(mainCrud)
}

func tambahLapangan() {
	var n int
	fmt.Println("=== MENU TAMBAH LAPANGAN ===")
	
	fmt.Println("Ingin menambahkan berapa lapangan?")
	fmt.Scan(&n)
	
	for i := 0; i < n; i++ {
		var nama, alamat, jenis string
		var harga int
		
		fmt.Println("Lapangan ke", i+1)
		fmt.Print("Masukkan Nama: ")
		fmt.Scan(&nama)
		fmt.Print("Masukkan Alamat: ")
		fmt.Scan(&alamat)
		fmt.Print("Masukkan Jenis: ")
		fmt.Scan(&jenis)
		fmt.Print("Masukkan Harga: ")
		fmt.Scan(&harga)
		
		lapangan = append(lapangan, Lapangan{
			nama: nama,
			alamat: alamat,
			jenis: jenis,
			harga: harga,
		})
	}
	fmt.Printf("Berhasil Menambahkan %d Lapangan Baru! \n", n)
	displayLap(lapangan, true, 0, "")
	menuLain(mainCrud)
}

func updateLapangan() {
	var n, harga int
	var nama, alamat, jenis string

	fmt.Println("=== MENU UPDATE LAPANGAN ===")
	jml := displayLap(lapangan, true, 0, "")

	fmt.Println("Pilih Lapangan Untuk Update:")
	fmt.Scan(&n)

	if n > jml {
		fmt.Println("Perintah Tidak Valid")
		menuLain(mainCrud)
	}

	displayLap(lapangan, false, n, "Update Lapangan")

	isAll := ""
	fmt.Println("Update semua data? (yes/no)")
	fmt.Scan(&isAll)

	if isAll == "yes" {
		fmt.Print("Ubah Nama: ")
		fmt.Scan(&nama)
		fmt.Print("Ubah Alamat: ")
		fmt.Scan(&alamat)
		fmt.Print("Ubah Jenis: ")
		fmt.Scan(&jenis)
		fmt.Print("Ubah Harga: ")
		fmt.Scan(&harga)

		lapangan[n-1].nama = nama
		lapangan[n-1].alamat = alamat
		lapangan[n-1].jenis = jenis
		lapangan[n-1].harga = harga 

		displayLap(lapangan, false, n, "Berhasil Update Lapangan")
		mainCrud()
	} else if isAll == "no" {
		var ubah string
		
		fmt.Print("Masukkan field yang ingin diubah: ")
		fmt.Scan(&ubah)
		
		switch {
			case ubah == "nama":
				nama := ""
				fmt.Print("Masukkan Nama: ")
				fmt.Scan(&nama)
				
				lapangan[n-1].nama = nama
			case ubah == "alamat":
				alamat := ""
				fmt.Print("Masukkan alamat: ")
				fmt.Scan(&alamat)
				
				lapangan[n-1].alamat = alamat
			case ubah == "jenis":
				jenis := ""
				fmt.Print("Masukkan jenis: ")
				fmt.Scan(&jenis)
				
				lapangan[n-1].jenis = jenis
			case ubah == "harga":
				harga := 0
				fmt.Print("Masukkan harga: ")
				fmt.Scan(&harga)
				
				lapangan[n-1].harga = harga
		}
		displayLap(lapangan, false, n, "Berhasil Update Lapangan")
		menuLain(mainCrud)
	} else {
		fmt.Println("Perintah Tidak Valid")
		fmt.Println(" ")
		menuLain(mainCrud)
	}
}

func hapusLapangan() {
	var n int

	fmt.Println("=== MENU HAPUS LAPANGAN ===")
	
	fmt.Println("Lapangan Tersedia:")
	jml := displayLap(lapangan, true, 0, "")
	
	fmt.Print("Pilih lapangan untuk dihapus: ")
	fmt.Scan(&n)

	if n > jml {
		fmt.Println("Perintah Tidak Valid")
		menuLain(mainReservation)
	}

	displayLap(lapangan, false, n, "Data Lapangan")
	fmt.Println("Anda yakin menghapusnya? (yes/no)")

	cond := ""
	fmt.Scan(&cond)

	if cond == "yes" {
		lapangan = append(lapangan[:n-1], lapangan[n:]...)
		fmt.Println("Berhasil hapus lapangan")
		menuLain(mainCrud)
	} else if cond == "no" {
		menuLain(mainCrud)
	} else {
		fmt.Println("Perintah Tidak Valid")
		menuLain(mainCrud)
	}
}