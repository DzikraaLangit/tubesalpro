package main

import "fmt"

type Materi struct{
	judul string
	isi string
	level int
}

type tabMateri[7] Materi

func tampilMateri(){
	var data tabMateri

	data = tabMateri{
	{
		judul: "Tipe Data",
		isi: `Tipe data adalah cara untuk memberitahu komputer jenis nilai apa yang disimpan dalam sebuah variabel.

Contoh tipe data dalam bahasa Go:
- int     : untuk bilangan bulat (contoh: 10)
- float64 : untuk bilangan desimal (contoh: 3.14)
- string  : untuk teks (contoh: "Halo")
- bool    : untuk nilai benar atau salah (true/false)

Contoh kode:
var umur int = 20
var pi float64 = 3.14
var nama string = "Budi"
var aktif bool = true

fmt.Println(umur, pi, nama, aktif)`,
		level: 1,
	},
	{
		judul: "Perulangan",
		isi: `Perulangan (looping) digunakan untuk mengeksekusi blok kode secara berulang.

Go hanya punya satu kata kunci perulangan, yaitu 'for'.

Contoh perulangan dari 1 sampai 5:
for i := 1; i <= 5; i++ {
	fmt.Println("Angka ke", i)
}

Contoh perulangan while-style (tanpa kondisi akhir di awal):
i := 0
for i < 3 {
	fmt.Println("i =", i)
	i++
}`,
		level: 1,
	},
	{
		judul: "Percabangan",
		isi: `Percabangan digunakan untuk mengambil keputusan berdasarkan kondisi tertentu.

Contoh penggunaan if, else if, dan else:
var nilai int = 85

if nilai >= 90 {
	fmt.Println("Nilai A")
} else if nilai >= 80 {
	fmt.Println("Nilai B")
} else {
	fmt.Println("Nilai C")
}

Go juga mendukung switch:
switch hari := "Senin"; hari {
case "Senin":
	fmt.Println("Awal pekan")
case "Jumat":
	fmt.Println("Akhir pekan")
default:
	fmt.Println("Hari biasa")
}`,
		level: 1,
	},
	{
		judul: "Fungsi",
		isi: `Fungsi (function) adalah blok kode yang bisa dipanggil berkali-kali dan dapat mengembalikan nilai.

Contoh fungsi sederhana:
func tambah(a int, b int) int {
	return a + b
}

Contoh pemanggilan:
hasil := tambah(3, 5)
fmt.Println("Hasil penjumlahan:", hasil)

Fungsi juga bisa tanpa parameter atau tanpa return value:
func sapa() {
	fmt.Println("Halo semua!")
}`,
		level: 2,
	},
	{
		judul: "Prosedur",
		isi: `Prosedur pada Go sebenarnya adalah fungsi yang tidak mengembalikan nilai.

Contoh prosedur:
func tampilPesan(nama string) {
	fmt.Println("Selamat datang,", nama)
}

Pemanggilan:
tampilPesan("Andi")

Prosedur digunakan saat kita hanya ingin melakukan aksi tanpa memerlukan nilai kembali (return).`,
		level: 2,
	},
	{
		judul: "Rekursif",
		isi: `Rekursif adalah fungsi yang memanggil dirinya sendiri, biasanya digunakan untuk menyelesaikan masalah yang dapat dipecah menjadi submasalah yang sama.

Contoh fungsi rekursif untuk menghitung faktorial:
func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * faktorial(n-1)
}

Pemanggilan:
fmt.Println("5! =", faktorial(5)) // Output: 120

Catatan: Gunakan rekursif dengan hati-hati agar tidak terjadi stack overflow.`,
		level: 3,
	},
	{
		judul: "Array",
		isi: `Array adalah struktur data yang menyimpan sekumpulan elemen dengan tipe yang sama.

Contoh array:
var angka [5]int = [5]int{1, 2, 3, 4, 5}

Mengakses elemen:
fmt.Println(angka[0]) // Output: 1

Mengisi array dengan loop:
for i := 0; i < len(angka); i++ {
	angka[i] = i * 2
}

Menampilkan seluruh array:
for _, val := range angka {
	fmt.Println(val)
}`,
		level: 3,
	},
}

	lanjutLoop := true

	for lanjutLoop {
		clearScreen()
		fmt.Println("\nDaftar Materi:\n")
		for i := 0; i < len(data); i++ {
			fmt.Printf("%d. %s (Level %d)\n", i+1, data[i].judul, data[i].level)
		}

		var pilihan int
		fmt.Print("\nPilih nomor materi yang ingin ditampilkan: ")
		fmt.Scanln(&pilihan)
		clearScreen()

		if pilihan >= 1 && pilihan <= len(data) {
			fmt.Printf("\n=== %s ===\n", data[pilihan-1].judul)
			fmt.Println(data[pilihan-1].isi)

			var lanjut string
			fmt.Print("\nMau lihat materi lain? (y/n): ")
			fmt.Scanln(&lanjut)

			if lanjut != "y" && lanjut != "Y" {
				lanjutLoop = false
				clearScreen()
			}
		} else {
			fmt.Println("Pilihan tidak valid.")
			 
		}
	}
}