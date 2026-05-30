package main

var lapangan = []Lapangan{
	{
		nama:   "Arena Futsal A",
		alamat: "Sidoarjo",
		harga:  120000,
		jenis:  "Futsal",
	},

	{
		nama:   "Soccer Center",
		alamat: "Surabaya",
		harga:  80000,
		jenis:  "Rumput Sintetis",
	},
}


var database = []Database{
	{
		tanggal:  "2026-05-23",
		lapangan: "Soccer Center",
		harga: 120000,
		jadwal:   defaultJadwal(),
		reservasi: []Sewa{
			{
				penyewa:  "Fathir",
				tglMulai: "2026-05-23",
				tglAkhir: "2026-05-23",
				jamMulai: "08.00",
				jamAkhir: "10.00",
				durasi:   2,
			},
			{
				penyewa:  "Budi",
				tglMulai: "2026-05-23",
				tglAkhir: "2026-05-23",
				jamMulai: "15.00",
				jamAkhir: "16.00",
				durasi:   1,
			},
		},
	},
	{
		tanggal:  "2026-05-23",
		lapangan: "Arena Futsal A",
		harga: 80000,
		jadwal:   defaultJadwal(),
		reservasi: []Sewa{
			{
				penyewa:  "Budi",
				tglMulai: "2026-05-23",
				tglAkhir: "2026-05-23",
				jamMulai: "09.00",
				jamAkhir: "11.00",
				durasi:   2,
			},
			{
				penyewa:  "Andi",
				tglMulai: "2026-05-23",
				tglAkhir: "2026-05-23",
				jamMulai: "18.00",
				jamAkhir: "20.00",
				durasi:   2,
			},
		},
	},
}

type Database struct {
	tanggal string
	lapangan string
	harga int
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
}