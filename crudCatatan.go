package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)


func main() {
	var data tabCatatan
	var lastID int
	menuCatatan(&data, &lastID)
}

func menuCatatan(catatan *tabCatatan, lastID *int) {
	for {
		fmt.Println("\n--- MENU ---")
		fmt.Println("1. Buat Catatan")
		fmt.Println("2. Edit Catatan")
		fmt.Println("3. Hapus Catatan")
		fmt.Println("4. Cari Catatan")
		fmt.Println("5. Tampilkan Semua Catatan")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			buatCatatan(catatan, lastID)
		case 2:
			editCatatan(catatan, *lastID)
		case 3:
			hapusCatatan(catatan, lastID)
		case 4:
			munculcariCatatan(catatan, *lastID)
		case 5:
			tampilCatatan(catatan, *lastID)
		case 6:
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func buatCatatan(catatan *tabCatatan, lastID *int) {
	if *lastID >= len(*catatan) {
		fmt.Println("Kapasitas catatan penuh.")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	 

	fmt.Print("Masukkan judul: ")
	judul, _ := reader.ReadString('\n')
	judul = strings.TrimSpace(judul)

	fmt.Print("Masukkan isi: ")
	isi, _ := reader.ReadString('\n')
	isi = strings.TrimSpace(isi)

	(*catatan)[*lastID] = Catatan{Judul: judul, Isi: isi, tanggal: time.Now()}
	(*lastID)++

	fmt.Println("Catatan berhasil dibuat.")
}

func munculcariCatatan(catatan *tabCatatan, lastID int){
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan judul catatan yang ingin diedit: ")
	judulCari, _ := reader.ReadString('\n')
	judulCari = strings.TrimSpace(judulCari)
	index := cariCatatan(*catatan, lastID, judulCari)

	if index == -1 {
		fmt.Println("Catatan tidak ditemukan.")
		return
	}

	fmt.Println("Judul   :", (*catatan)[index].Judul)
    fmt.Println("Tanggal :", (*catatan)[index].tanggal.Format("02-01-2006"))
    fmt.Println("Isi     :", (*catatan)[index].Isi)

}

func cariCatatan(catatan tabCatatan, lastID int, judulCari string) int {
	judulCari = strings.ToLower(judulCari)
	for i := 0; i < lastID; i++ {
		if strings.ToLower(catatan[i].Judul) == judulCari {
			return i
		}
	}
	return -1
}

func editCatatan(catatan *tabCatatan, lastID int) {
	var tanggallama time.Time
	
	if lastID == 0 {
		fmt.Println("Belum ada catatan.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	 

	fmt.Print("Masukkan judul catatan yang ingin diedit: ")
	judulCari, _ := reader.ReadString('\n')
	judulCari = strings.TrimSpace(judulCari)

	index := cariCatatan(*catatan, lastID, judulCari)
	if index == -1 {
		fmt.Println("Catatan tidak ditemukan.")
		return
	}

	fmt.Print("Masukkan judul baru: ")
	judulBaru, _ := reader.ReadString('\n')
	judulBaru = strings.TrimSpace(judulBaru)

	fmt.Print("Masukkan isi baru: ")
	isiBaru, _ := reader.ReadString('\n')
	isiBaru = strings.TrimSpace(isiBaru)

	tanggallama = catatan[index].tanggal

	(*catatan)[index] = Catatan{Judul: judulBaru, Isi: isiBaru, tanggal: tanggallama}
	fmt.Println("Catatan berhasil diedit.")
}

func hapusCatatan(catatan *tabCatatan, lastID *int) {
	if *lastID == 0 {
		fmt.Println("Belum ada catatan.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	 

	fmt.Print("Masukkan judul catatan yang ingin dihapus: ")
	judulCari, _ := reader.ReadString('\n')
	judulCari = strings.TrimSpace(judulCari)

	index := cariCatatan(*catatan, *lastID, judulCari)
	if index == -1 {
		fmt.Println("Catatan tidak ditemukan.")
		return
	}

	for i := index; i < *lastID-1; i++ {
		(*catatan)[i] = (*catatan)[i+1]
	}
	*lastID--

	fmt.Println("Catatan berhasil dihapus.")
}

func tampilCatatan(catatan *tabCatatan, lastID int) {
    var Pilihan string

    if lastID == 0 {
        fmt.Println("Belum ada catatan.")
        return
    }

    fmt.Print("Tampilkan berdasarkan terbaru/terlama?: ")
    fmt.Scan(&Pilihan)

	bufio.NewReader(os.Stdin).ReadString('\n')

    if Pilihan == "terbaru" {
        SelectionSortCatatanTurun(catatan, lastID)
    } else if Pilihan == "terlama" {
        SelectionSortCatatanNaik(catatan, lastID)
    } else {
        fmt.Println("Jawab cing baleg hama")
        return
    }

    fmt.Println("\nDaftar Catatan:")
    for i := 0; i < lastID; i++ {
        fmt.Printf("\nCatatan #%d\n", i+1)
        fmt.Println("Judul   :", (*catatan)[i].Judul)
        fmt.Println("Tanggal :", (*catatan)[i].tanggal.Format("02-01-2006"))
        fmt.Println("Isi     :", (*catatan)[i].Isi)
    }
}


func SelectionSortCatatanNaik(catatan *tabCatatan, lastID int) {
    for i := 0; i < lastID-1; i++ {
        minIdx := i
        for j := i + 1; j < lastID; j++ {
            if (*catatan)[j].tanggal.Before((*catatan)[minIdx].tanggal) {
                minIdx = j
            }
        }
        (*catatan)[i], (*catatan)[minIdx] = (*catatan)[minIdx], (*catatan)[i]
    }
}

func SelectionSortCatatanTurun(catatan *tabCatatan, lastID int) {
    for i := 0; i < lastID-1; i++ {
        maxIdx := i
        for j := i + 1; j < lastID; j++ {
            if (*catatan)[j].tanggal.After((*catatan)[maxIdx].tanggal) {
                maxIdx = j
            }
        }
        (*catatan)[i], (*catatan)[maxIdx] = (*catatan)[maxIdx], (*catatan)[i]
    }
}

