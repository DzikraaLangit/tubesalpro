package main

import "fmt"

func soal() {
	var menu=1
	var nomer int
	fmt.Println("\n===================================================")
	fmt.Printf("%30s\n", "PILIH MATERI SOAL")
	fmt.Println("===================================================")
	fmt.Println("\nALPRO 1")
	fmt.Println("1. Tipe Data")
	fmt.Println("2. Perulangan")
	fmt.Println("3. Percabangan")
	fmt.Println("\nALPRO 2")
	fmt.Println("4. Fungsi")
	fmt.Println("5. Prosedur")
	fmt.Println("6. Rekursif")
	fmt.Println("7. Array")
	fmt.Println()
	fmt.Println("8. Back To Menu")
	fmt.Println()
	fmt.Print("Pilih Nomer: ")
	fmt.Scan(&nomer)
	fmt.Println()
	clearScreen()

	switch nomer {
	case 1:
		tipedata(&latsol)
	case 2:
		perulangan(&latsol)
	case 3:
		percabangan(&latsol)
	case 4:
		fungsi(&latsol)
	case 5:
		prosedur(&latsol)
	case 6:
		rekursif(&latsol)
	case 7:
		array(&latsol)
	case 8:
		inter(&menu)
		clearScreen()
	default:
		fmt.Println("Soal belum tersedia untuk pilihan ini.")
	}
}

func tipedata(l *tabnilai) {
	var ans string
	var temp int
	temp=l[0].nilai
	l[0].nilai=0
	l[0].tipe = "Tipe Data"
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println("1. Tipe data yang hanya diisi dengan bilangan bulat adalah?")
	fmt.Println()
	fmt.Println("a. string\nb. integer\nc. float\nd. boolean")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "b" {
		l[0].nilai += 20
	}
	ans=""

	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("2. Tipe data yang menghasilkan true atau false adalah?")
	fmt.Println()
	fmt.Println("a. string\nb. integer\nc. float\nd. boolean")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "d" {
		l[0].nilai += 20
	}
	ans=""
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("3. Tipe data yang bisa diisi dengan bilangan real adalah?")
	fmt.Println()
	fmt.Println("a. string\nb. integer\nc. float\nd. boolean")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "c" {
		l[0].nilai += 20
	}
	ans=""
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("4. Tipe data yang digunakan untuk menyimpan nama adalah?")
	fmt.Println()
	fmt.Println("a. string\nb. integer\nc. float\nd. boolean")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "a" {
		l[0].nilai += 20
	}
	ans=""
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("5. Tipe data yang bisa melakukan operasi aritmatika adalah?")
	fmt.Println()
	fmt.Println("a. string\nb. integer\nc. char\nd. boolean")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "b" {
		l[0].nilai += 20
	}
	ans=""
	clearScreen()
	fmt.Println()
	fmt.Printf("Total Nilai Kamu : %d\n", l[0].nilai)
	if l[0].nilai<temp{
		l[0].nilai=temp
	}
}

func perulangan(l *tabnilai) {
	var ans string
	var temp int
	temp=l[1].nilai
	l[1].nilai=0
	l[1].tipe = "Perulangan"
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println("1. Perulangan dengan jumlah iterasi tidak pasti?")
	fmt.Println()
	fmt.Println("a. for i := 0; i < 10; i++\nb. for i := 1; i <= 5; i++\nc. for i < 100 { ... }\nd. for i := 0; i <= 100; i += 2")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "c" {
		l[1].nilai += 20
	}
	ans=""
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("2. Fungsi dari break dalam perulangan?")
	fmt.Println()
	fmt.Println("a. Melanjutkan ke iterasi berikutnya\nb. Keluar dari perulangan\nc. Mengulang dari awal\nd. Menampilkan hasil")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "b" {
		l[1].nilai += 20
	}
	ans=""
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("3. i=1")
	fmt.Println("for i < 4 {")
	fmt.Println("fmt.Print(i)")
	fmt.Println("}")
	fmt.Println("Apa Output kode di atas?")
	fmt.Println()
	fmt.Println("a. 123\nb. 0123\nc. 234\nd. 135")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "a" {
		l[1].nilai += 20
	}
	ans=""
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("4. for i=2;i<=5;i++{")
	fmt.Println("fmt.Print(i)")
	fmt.Println("}")
	fmt.Println("Apa Output kode di atas?")
	fmt.Println()
	fmt.Println("a. 123\nb. 0123\nc. 234\nd. 135")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "c" {
		l[1].nilai += 20
	}
	ans=""
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("5. for i=10;i>5;i-=2{")
	fmt.Println("fmt.Print(i)")
	fmt.Println("}")
	fmt.Println("Apa Output kode di atas?")
	fmt.Println()
	fmt.Println("a. 109876\nb. 1086\nc. 108\nd. 579")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "b" {
		l[1].nilai += 20
	}
	ans=""
	clearScreen()
	fmt.Println()
	fmt.Printf("Total Nilai Kamu : %d\n", l[1].nilai)
	if l[1].nilai<temp{
		l[1].nilai=temp
	}
	
}

func percabangan(l *tabnilai){
	var ans string
	var temp int
	temp=l[2].nilai
	l[2].nilai=0
	l[2].tipe="Percabangan"
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("1. Pernyataan if dalam Go digunakan untuk?")
	fmt.Println()
	fmt.Println("a. Mengulang proses\nb. Menyimpan nilai\nc. Mengevaluasi kondisi dan mengeksekusi kode berdasarkan hasilnya\nd. Mencetak ke layar")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "c" {
		l[2].nilai += 20
	}
	ans=""
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("2. Bagaimana bentuk umum pernyataan if-else dalam Go?")
	fmt.Println()
	fmt.Println("a. if condition then doSomething()\nb. if (condition) { } else { }\nc. when condition { ... }\nd. if: condition => ...")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "b" {
		l[2].nilai += 20
	}
	ans=""
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("3. Fungsi else if digunakan untuk?")
	fmt.Println()
	fmt.Println("a. Menjalankan dua kondisi if secara paralel\nb. Menambahkan kondisi tambahan jika kondisi sebelumnya salah\nc. Mencetak hasil dari kondisi\nd. Menyimpan hasil kondisi")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "b" {
		l[2].nilai += 20
	}
	ans=""
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("3. x=10\nif x>5{\nfmt.Println('lebih besar dari 5')\n}else{\nfmt.Println('lebih kecil atau sama dengan 5')}")
	fmt.Println("Apa Output kode di atas?")
	fmt.Println()
	fmt.Println("a.  Error\nb. Lebih besar dari 5\nc. Lebih kecil atau sama dengan 5\nd. Tidak mencetak apapun")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "b" {
		l[2].nilai += 20
	}
	ans=""
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("3. nilai := 85\nif nilai >= 90 {\nfmt.Println('a')\n} else if nilai >= 80 {\nfmt.Println('b')\n} else {\nfmt.Println('c')\n}")
	fmt.Println("Apa Output kode di atas?")
	fmt.Println()
	fmt.Println("a. a\nb. b\nc. c\nd. d")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "b" {
		l[2].nilai += 20
	}
	ans=""
	clearScreen()
	fmt.Println()
	fmt.Printf("Total Nilai Kamu : %d\n", l[2].nilai)
	if l[2].nilai<temp{
		l[2].nilai=temp
	}
}
func fungsi(l *tabnilai){
	var ans string
	var temp int
	temp=l[3].nilai
	l[3].nilai=0
	l[3].tipe="Fungsi"
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("1. Apa yang dimaksud dengan fungsi dalam pemrograman?")
	fmt.Println()
	fmt.Println("a. Perintah untuk mencetak data\nb. Struktur penyimpanan data\nc. Blok kode yang digunakan untuk melakukan tugas tertentu\nd. Variabel untuk menyimpan angka")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "c" {
		l[3].nilai += 20
	}
	ans=""
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("2. Bagaimana cara mendeklarasikan fungsi di Go?")
	fmt.Println()
	fmt.Println("a. function namaFungsi() { ... }\nb. def namaFungsi() { ... }\nc. void namaFungsi() { ... }\nd. func namaFungsi() { ... }")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "d" {
		l[3].nilai += 20
	}
	ans=""
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("3. Fungsi bisa mengembalikan nilai menggunakan?")
	fmt.Println()
	fmt.Println("a. send\nb. get\nc. return\nd. output")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "c" {
		l[3].nilai += 20
	}
	ans=""
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("4. func tambah(a int, b int) int {\nreturn a + b\n}\nfunc main() {\nfmt.Println(tambah(3, 4))\n}")
	fmt.Println("Apa Output kode di atas?")
	fmt.Println()
	fmt.Println("a. 7\nb. 34\nc. tambah\nd. error")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "a" {
		l[3].nilai += 20
	}
	ans=""
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("5. func sapa() {\nfmt.Println('Halo!')\n}\nfunc main() {\nsapa()\n}")
	fmt.Println("Apa Output kode di atas?")
	fmt.Println()
	fmt.Println("a. Halo!\nb. sapa\nc. Halo\nd. error")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "a" {
		l[3].nilai += 20
	}
	ans=""
	clearScreen()
	fmt.Println()
	fmt.Printf("Total Nilai Kamu : %d\n", l[3].nilai)
	if l[3].nilai<temp{
		l[3].nilai=temp
	}

	
}
func prosedur(l *tabnilai){
	var ans string
	var temp int
	temp=l[4].nilai
	l[4].nilai=0
	l[4].tipe="Prosedur"
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("1. Apa yang membedakan prosedur dari fungsi dalam pemrograman?")
	fmt.Println()
	fmt.Println("a. Prosedur tidak menggunakan parameter\nb.Prosedur tidak mengembalikan nilai\nc. Prosedur hanya bisa dipanggil satu kali\nd. Prosedur dijalankan otomatis tanpa dipanggil")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "b" {
		l[4].nilai += 20
	}
	ans=""
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("2. Perintah fmt.Println() di Go merupakan contoh dari:")
	fmt.Println()
	fmt.Println("a. Fungsi rekursif\nb. Prosedur karena tidak mengembalikan nilai\nc. Fungsi yang mengembalikan string\nd. Operator logika")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "b" {
		l[4].nilai += 20
	}
	ans=""
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("3. Apa manfaat utama dari menggunakan prosedur dalam program?")
	fmt.Println()
	fmt.Println("a. Untuk memperlambat proses komputasi\nb. Untuk membingungkan pembaca kode\nc. Untuk membagi program menjadi bagian-bagian modular\nd. Untuk membuat program hanya bisa jalan satu kali")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "c" {
		l[4].nilai += 20
	}
	ans=""
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("4. Bagaimana cara memanggil prosedur bernama tampilkan() di Go?")
	fmt.Println()
	fmt.Println("a. tampilkan\nb. call tampilkan\nc. tampilkan()\nd. call tampilkan()")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "c" {
		l[4].nilai += 20
	}
	ans=""
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("5.func cetakPesan() {\nfmt.Println('Halo dunia!')\n}")
	fmt.Println("Apa yang dilakukan prosedur cetakPesan di atas?")
	fmt.Println()
	fmt.Println("a. Mengembalikan nilai string\nb. Menampilkan pesan ke layar\nc. Menyimpan input pengguna\nd. Mengubah nilai variabel global")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "b" {
		l[4].nilai += 20
	}
	ans=""
	clearScreen()
	fmt.Println()
	fmt.Printf("Total Nilai Kamu : %d\n", l[4].nilai)
	if l[4].nilai<temp{
		l[4].nilai=temp
	}
}
func rekursif(l *tabnilai){
	var ans string
	var temp int
	temp=l[5].nilai
	l[5].nilai=0
	l[5].tipe="Rekursif"
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("1. Apa itu fungsi rekursif?")
	fmt.Println()
	fmt.Println("a. Fungsi yang memanggil prosedur lain\nb. Fungsi yang dijalankan satu kali saja\nc. Fungsi yang memanggil dirinya sendiri\nd. Fungsi yang tidak memiliki parameter")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "c" {
		l[5].nilai += 20
	}
	ans=""
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("2. Apa syarat penting dalam membuat fungsi rekursif?")
	fmt.Println()
	fmt.Println("a. Harus menggunakan dua parameter\nb. Harus menggunakan for\nc. Harus memiliki kondisi berhenti (base case)\nd. Harus memanggil fungsi main")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "c" {
		l[5].nilai += 20
	}
	ans=""
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("3. Apa kelemahan utama penggunaan fungsi rekursif yang tidak efisien?")
	fmt.Println()
	fmt.Println("a. Program menjadi lebih pendek\nb. Fungsi tidak bisa dijalankan\nc. Dapat menyebabkan stack overflow\nd. Menghasilkan angka acak")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "c" {
		l[5].nilai += 20
	}
	ans=""
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("4.func faktorial(n int) int {\nif n == 1 {\nreturn 1\n}\nreturn n * faktorial(n-1)\n}")
	fmt.Println("Apa hasil dari pemanggilan faktorial(3)?")
	fmt.Println()
	fmt.Println("a.6\nb.9\nc.3\nd.1")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "a" {
		l[5].nilai += 20
	}
	ans=""
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("5.func cetak(n int) {\nif n == 0 {\nreturn\n}\nfmt.Println(n)\ncetak(n-1)\n}")
	fmt.Println("Apa hasil dari pemanggilan cetak(3)?")
	fmt.Println()
	fmt.Println("a. 1 2 3\nb. 3 2 1\nc. 3 3 3\nd. 1 1 1")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "b" {
		l[5].nilai += 20
	}
	ans=""
	clearScreen()
	fmt.Println()
	fmt.Printf("Total Nilai Kamu : %d\n", l[5].nilai)
	if l[5].nilai<temp{
		l[5].nilai=temp
	}
}
func array(l *tabnilai){
	var ans string
	var temp int
	temp=l[6].nilai
	l[6].nilai=0
	l[6].tipe="Array"
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("1. Apa yang dimaksud dengan array?")
	fmt.Println()
	fmt.Println("a. Tipe data untuk menyimpan karakter\nb. Variabel tunggal yang menyimpan banyak nilai sejenis\nc. Fungsi untuk mencetak nilai\nd. Variabel yang hanya menyimpan angka 1 dan 0")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "b" {
		l[6].nilai += 20
	}
	ans=""
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("2. Bagaimana cara mengakses elemen pertama dari array data di Go?")
	fmt.Println()
	fmt.Println("a. data.1\nb. data(0)\nc. data[0]\nd. data{0}")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "c" {
		l[6].nilai += 20
	}
	ans=""
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("3. Manakah yang merupakan pernyataan yang benar tentang array di Go?")
	fmt.Println()
	fmt.Println("a. Panjang array bisa berubah saat runtime\nb. Array menyimpan nilai bertipe berbeda\nc. Elemen array diakses dengan indeks\nd. Indeks array dimulai dari 1")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "c" {
		l[6].nilai += 20
	}
	ans=""
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("4.var angka [5]int")
	fmt.Println("Deklarasi array berikut menyimpan berapa elemen?")
	fmt.Println()
	fmt.Println("a. 4\nb. 5\nc. 6\nd. 7")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "b" {
		l[6].nilai += 20
	}
	ans=""
	for ans!="a"&&ans!="b"&&ans!="c"&&ans!="d"{
	clearScreen()
	fmt.Println()
	fmt.Println("5.data := [3]int{2, 4, 6}\nfmt.Println(data[1])")
	fmt.Println("Apa output dari kode berikut?")
	fmt.Println()
	fmt.Println("a. 2\nb.4\nc. 6\nd. 3")
	fmt.Print("Jawaban: ")
	fmt.Scan(&ans)
	}
	if ans == "b" {
		l[6].nilai += 20
	}
	ans=""
	clearScreen()
	fmt.Println()
	fmt.Printf("Total Nilai Kamu : %d\n", l[6].nilai)
	if l[6].nilai<temp{
		l[6].nilai=temp
	}
}