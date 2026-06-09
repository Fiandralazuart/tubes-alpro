package main

var lapangan = []Lapangan{
	{
		nama:   "Arena Futsal A",
		alamat: "Sidoarjo",
		harga:  harga{
			hargaDefault: 120000,
			happyHour: 150000,
		},
		jenis:  "Futsal",
	},

	{
		nama:   "Soccer Center",
		alamat: "Surabaya",
		harga:  harga{
			hargaDefault: 80000,
			happyHour: 120000,
		},
		jenis:  "Rumput Sintetis",
	},
}


var database = []Database{
	{
		tanggal:  tanggal{
			hari: "23",
			bulan: "05",
			tahun: "2026",
		},
		lapangan: "Soccer Center",
		harga:  harga{
			hargaDefault: 120000,
			happyHour: 150000,
		},
		jadwal:   defaultJadwal(120000, 150000),
		reservasi: []Sewa{
			{
				penyewa:  "Fathir",
				tglMulai: "2026-05-23",
				tglAkhir: "2026-05-23",
				jamMulai: "08.00",
				jamAkhir: "10.00",
				durasi:   2,
				total: 240000,
			},
			{
				penyewa:  "Budi",
				tglMulai: "2026-05-23",
				tglAkhir: "2026-05-23",
				jamMulai: "15.00",
				jamAkhir: "16.00",
				durasi:   1,
				total: 120000,
			},
		},
	},
	{
		tanggal:  tanggal{
			hari: "23",
			bulan: "05",
			tahun: "2026",
		},
		lapangan: "Arena Futsal A",
		harga:  harga{
			hargaDefault: 80000,
			happyHour: 120000,
		},
		jadwal:   defaultJadwal(80000, 120000),
		reservasi: []Sewa{
			{
				penyewa:  "Budi",
				tglMulai: "2026-05-23",
				tglAkhir: "2026-05-23",
				jamMulai: "09.00",
				jamAkhir: "11.00",
				durasi:   2,
				total: 240000,
			},
			{
				penyewa:  "Andi",
				tglMulai: "2026-05-23",
				tglAkhir: "2026-05-23",
				jamMulai: "18.00",
				jamAkhir: "20.00",
				durasi:   2,
				total: 240000,
			},
		},
	},
}

type Database struct {
	tanggal tanggal
	lapangan string
	harga harga
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
	total int
}



type Jam struct {
	waktu string
	isAvailable bool
	harga int
}

type Lapangan struct {
	nama string
	alamat string
	harga harga
	jenis string
}

type tanggal struct {
	hari string
	bulan string
	tahun string
}

type harga struct {
	hargaDefault int
	happyHour int
}

type stats struct {
	bulan string
	reservasi int
	jam int
	total int
	hari [14]int
}