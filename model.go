package main

import "time"

type Catatan struct {
	Judul   string
	Isi     string
	tanggal time.Time
}

type tabstring [nmax]string
type tabnilai [nmax]nilaisoal
type Jadwalsche struct {
    Hari   string
    Jam    int
    Materi string
}
type nilaisoal struct {
	tipe  string
	nilai int
}

type tabCatatan [100]Catatan