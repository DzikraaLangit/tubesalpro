package main

import (
	"fmt"
	"os"
    "os/exec"
    "runtime"
	"math/rand"
    "sort"
    "time"
	"bufio"
)

const nmax int = 1000

var latsol tabnilai
var totalHours int

func clearScreen() {
    switch runtime.GOOS {
    case "windows":
        cmd := exec.Command("cmd", "/c", "cls")
        cmd.Stdout = os.Stdout
        cmd.Run()
    default:
        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}
func early(logreg *int) {
	fmt.Println("\n===================================================")
	fmt.Printf("%30s\n", "SELAMAT DATANG")
	fmt.Printf("%47s\n", "Silahkan Log in atau Register terlebih dahulu")
	fmt.Println("===================================================")
	fmt.Println("1. Log In")
	fmt.Println("2. Register")
	fmt.Print("Pilih (1/2): ")
	fmt.Scan(logreg)
	fmt.Println()
	clearScreen()
}

func register(usn, pass *string) {
	fmt.Println("\n===================================================")
	fmt.Printf("%30s\n", "REGISTRASI")
	fmt.Println("===================================================")
	fmt.Print("Masukkan Username : ")
	fmt.Scan(usn)
	fmt.Print("Masukkan Password : ")
	fmt.Scan(pass)
	fmt.Println()
	clearScreen()
}

func login(usn, pass string) {
	var lusn, lpass string
	fmt.Println("\n===================================================")
	fmt.Printf("%27s\n", "LOGIN")
	fmt.Println("===================================================")
	for {
		fmt.Print("Masukkan Username : ")
		fmt.Scan(&lusn)
		fmt.Print("Masukkan Password : ")
		fmt.Scan(&lpass)

		if lusn == usn && lpass == pass{
			break
		}
		fmt.Println("Username atau password salah. Silakan coba lagi.\n")
	}
		fmt.Println("\nLogin berhasil!\n")
		clearScreen()
}

func inter(menu *int) {
	fmt.Println("\n===================================================")
	fmt.Printf("%35s\n", "SELAMAT DATANG DI AIPRO")
	fmt.Println("===================================================")
	fmt.Println("Pilih Menu:")
	fmt.Println("1. Catatan")
	fmt.Println("2. Materi")
	fmt.Println("3. Soal")
	fmt.Println("4. Jadwal")
	fmt.Println("5. Nilai")
	fmt.Println("0. Exit")
	fmt.Print("Masukkan nomor dari menu: ")
}

func inmenu(menu int) {
	var materi tabstring
	var data tabCatatan
	var lastID int
	switch menu{
		case 0:
			fmt.Print("")
		case 1:
			menuCatatan(&data, &lastID)
		case 2:
			tampilMateri()
		case 3:
			soal()
		case 4:
			materiLen := IsiMateri(&materi)
			fmt.Println("Masukkan Total Jam Dalam Seminggu Untuk Belajar:")
			fmt.Scan(&totalHours)
    		Schedule(totalHours, materi, materiLen)
		case 5:
			selectionSortNilai(&latsol)
			cetaknilai(latsol)
		default:
			fmt.Println("Menu belum tersedia untuk pilihan ini.")
	}
}
func backtomenu(){
	fmt.Print("Kembali Ke Menu (Enter)")
	fmt.Scanln()
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	clearScreen()
}

func cetaknilai(l tabnilai) {
	fmt.Println("\n=== Nilai Materi Dari Terbesar Hingga Terkecil ===")
	for i := 0; i < 7; i++ {
		if l[i].tipe != "" {
			fmt.Printf("%d. %s: %d\n", i+1, l[i].tipe, l[i].nilai)
		}
	}
}

func selectionSortNilai(l *tabnilai) {
	for i := 0; i < 7; i++ {
		maxIdx := i
		for j := i + 1; j < 7; j++ {
			if l[j].nilai > l[maxIdx].nilai {
				maxIdx = j
			}
		}
		l[i], l[maxIdx] = l[maxIdx], l[i]
	}
}

func Schedule(totalHours int, materi tabstring, materiLen int) {
    rand.Seed(time.Now().UnixNano())
    const startHour = 8
    const endHour = 22
    days := [7]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
    var hours [endHour - startHour]int
    for i := 0; i < len(hours); i++ {
        hours[i] = startHour + i
    }

    maxSlots := len(days) * len(hours)
    if totalHours > maxSlots {
        fmt.Printf("Jumlah jam terlalu banyak. Maksimal tersedia %d jam.\n", maxSlots)
        return
    }

    var jadwal [nmax]Jadwalsche
    jadwalLen := 0

    isDuplicate := func(day string, hour int) bool {
        for i := 0; i < jadwalLen; i++ {
            if jadwal[i].Hari == day && jadwal[i].Jam == hour {
                return true
            }
        }
        return false
    }

    for count := 0; count < totalHours; {
        day := days[rand.Intn(len(days))]
        hour := hours[rand.Intn(len(hours))]
        if !isDuplicate(day, hour) {
            topic := materi[rand.Intn(materiLen)]
            jadwal[jadwalLen] = Jadwalsche{Hari: day, Jam: hour, Materi: topic}
            jadwalLen++
            count++
        }
    }

    sort.Slice(jadwal[:jadwalLen], func(i, j int) bool {
        dayIndex := func(day string) int {
            for idx, d := range days {
                if d == day {
                    return idx
                }
            }
            return -1
        }
        di := dayIndex(jadwal[i].Hari)
        dj := dayIndex(jadwal[j].Hari)
        if di != dj {
            return di < dj
        }
        return jadwal[i].Jam < jadwal[j].Jam
    })

    fmt.Println("\n====================================")
    fmt.Println("     Jadwal Belajar Mingguan     ")
    fmt.Println("====================================")

    currentDay := ""
    for i := 0; i < jadwalLen; i++ {
        if jadwal[i].Hari != currentDay {
            currentDay = jadwal[i].Hari
            fmt.Printf("\n%s:\n", currentDay)
        }
        fmt.Printf("  • %02d:00–%02d:00 : %s\n", jadwal[i].Jam, jadwal[i].Jam+1, jadwal[i].Materi)
    }

    fmt.Println("\n====================================")
    fmt.Println("Selamat belajar! Tetap semangat! ")
    fmt.Println("====================================")
}

func IsiMateri(materi *tabstring) int {
    materi[0] = "Tipe Data"
    materi[1] = "Perulangan"
    materi[2] = "Percabangan"
    materi[3] = "Fungsi"
    materi[4] = "Prosedur"
    materi[5] = "Rekursif"
    materi[6] = "Array"
    return 7
}










