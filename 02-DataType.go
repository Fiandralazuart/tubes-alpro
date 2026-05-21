package main

var lapangan = []Lapangan{
	{
		nama:   "Arena Futsal A",
		alamat: "Sidoarjo",
		harga:  120000,
		jenis:  "Futsal",
		// sewa: []Sewa{
		// 	{
		// 		penyewa: "Jamal",
		// 		tglMulai: "2026-05-21",
		// 		tglAkhir: "2026-05-21",
		// 		jamMulai: "08:00",
		// 		jamAkhir: "10:00",
		// 		durasi:   2,
		// 	},
		// 	{
		// 		penyewa: "Udin",
		// 		tglMulai: "2026-05-22",
		// 		tglAkhir: "2026-05-22",
		// 		jamMulai: "13:00",
		// 		jamAkhir: "15:00",
		// 		durasi:   2,
		// 	},
		// },
	},

	{
		nama:   "Soccer Center",
		alamat: "Surabaya",
		harga:  80000,
		jenis:  "Rumput Sintetis",
		// sewa: []Sewa{
		// 	{
		// 		penyewa: "Rahman",
		// 		tglMulai: "2026-05-23",
		// 		tglAkhir: "2026-05-23",
		// 		jamMulai: "19:00",
		// 		jamAkhir: "21:00",
		// 		durasi:   2,
		// 	},
		// },
	},
}


var database = []Database{
	{
		tanggal: "2026-05-23",
		lapangan: "Soccer Center",
		jadwal: defaultJadwal(),
	},
	{
		tanggal: "2026-05-23",
		lapangan: "Arena Futsal A",
		jadwal: defaultJadwal(),
	},
} 

type Database struct {
	tanggal string
	lapangan string
	jadwal []Jam
	reservasi []Sewa
}

type Sewa struct {
	penyewa string
	tglMulai string
	tglAkhir string 
	jamMulai string 
	jamAkhir string 
	durasi int
}



type Jam struct {
	waktu string
	isAvailable bool
}

type Lapangan struct {
	nama string
	alamat string
	harga int
	jenis string
	// sewa []Sewa
}