package main

type Materi struct{
	judul string
	isi string
	level int
}

type tabMateri[100] Materi

func materi