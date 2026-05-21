package main
import "fmt"

type Sewa struct {
	tglMulai string
	tglAkhir string 
	jamMulai string 
	jamAkhir string 
	durasi int
}

type Lapangan struct {
	nama string
	alamat string
	harga int
	jenis string
	sewa []Sewa
}

var lapangan = []Lapangan{
	{
		nama:   "Arena Futsal A",
		alamat: "Sidoarjo",
		harga:  120000,
		jenis:  "Futsal",
		sewa: []Sewa{
			{
				tglMulai: "2026-05-21",
				tglAkhir: "2026-05-21",
				jamMulai: "08:00",
				jamAkhir: "10:00",
				durasi:   2,
			},
			{
				tglMulai: "2026-05-22",
				tglAkhir: "2026-05-22",
				jamMulai: "13:00",
				jamAkhir: "15:00",
				durasi:   2,
			},
		},
	},

	{
		nama:   "Soccer Center",
		alamat: "Surabaya",
		harga:  80000,
		jenis:  "Rumput Sintetis",
		sewa: []Sewa{
			{
				tglMulai: "2026-05-23",
				tglAkhir: "2026-05-23",
				jamMulai: "19:00",
				jamAkhir: "21:00",
				durasi:   2,
			},
		},
	},
}
func mainCrud() {

	var n int
	fmt.Println("Selamat Datang")
	
	fmt.Println("Pilih Menu")
	fmt.Println("1. Tampilkan Lapangan")
	fmt.Println("2. Tambahkan Lapangan")
	fmt.Println("3. Edit Lapangan")
	fmt.Println("4. Hapus Lapangan")
	fmt.Scan(&n)

	switch {
		case n == 1:
			tampilkanLapangan(lapangan)
		case n == 3:
			updateLapangan(lapangan)
		default:
			fmt.Println("TITID")
	}

}

func tampilkanLapangan(lapangan []Lapangan) {
	displayLap(lapangan, true, 0, "")

	menuLain()
}

func updateLapangan(lapangan []Lapangan) {
	var n, harga int
	var nama, alamat, jenis string

	fmt.Println("=== MENU UPDATE LAPANGAN ===")
	displayLap(lapangan, true, 0, "")

	fmt.Println("Pilih Lapangan Untuk Update:")
	fmt.Scan(&n)

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
	} else {
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
		menuLain()
	}
}