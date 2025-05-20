package main

import "time"

type Catatan struct {
	Judul   string
	Isi     string
	tanggal time.Time
}

type tabCatatan [100]Catatan