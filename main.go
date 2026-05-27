package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// struct untuk nyimpen data film
type Film struct {
	judul     string
	genre     string
	tahun     int
	deskripsi string
	rating    float64
}

// variabel global daftar film dan jumlahnya
var daftarFilm []Film
var jumlahFilm int = 0

var reader = bufio.NewReader(os.Stdin)

// baca input string dari user
func bacaInput(prompt string) string {
	fmt.Print(prompt)
	teks, _ := reader.ReadString('\n')
	teks = strings.TrimRight(teks, "\r\n")
	return teks
}

// baca input angka bulat
func bacaInt(prompt string) int {
	for {
		str := bacaInput(prompt)
		angka, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("  [!] Input harus berupa angka, coba lagi.")
		} else {
			return angka
		}
	}
}

// baca input angka desimal buat rating
func bacaFloat(prompt string) float64 {
	for {
		str := bacaInput(prompt)
		angka, err := strconv.ParseFloat(str, 64)
		if err != nil || angka < 0 || angka > 10 {
			fmt.Println("  [!] Rating harus angka 0-10, coba lagi.")
		} else {
			return angka
		}
	}
}

// nampilin garis pemisah
func garis() {
	fmt.Println("--------------------------------------------------")
}

// menampilkan menu utama
func tampilMenu() {
	fmt.Println()
	fmt.Println("++++++++++++++ CineReview ++++++++++++++")
	fmt.Println("  [1] Tambah Film")
	fmt.Println("  [2] Tampilkan Semua Film")
	fmt.Println("  [0] Keluar")
	fmt.Println("========================================")
}

// fungsi buat nambahin film baru ke slice
func tambahFilm() {
	fmt.Println()
	garis()
	fmt.Println("  >> TAMBAH FILM BARU")
	garis()

	var f Film

	f.judul = bacaInput("  Judul Film     : ")
	if f.judul == "" {
		fmt.Println("  [!] Judul tidak boleh kosong.")
		return
	}

	f.genre = bacaInput("  Genre          : ")
	f.tahun = bacaInt("  Tahun Rilis    : ")
	f.deskripsi = bacaInput("  Deskripsi      : ")
	f.rating = bacaFloat("  Rating (0-10)  : ")

	// masukkan film ke slice
	daftarFilm = append(daftarFilm, f)
	jumlahFilm++

	fmt.Println()
	fmt.Println("  [OK] Film berhasil ditambahkan!")
	garis()
}

// procedure buat nampilin semua data film yang ada
func tampilSemuaFilm() {
	fmt.Println()
	garis()
	fmt.Println("  >> DAFTAR FILM")
	garis()

	if jumlahFilm == 0 {
		fmt.Println("  Belum ada film yang tersimpan.")
		garis()
		return
	}

	for i := 0; i < jumlahFilm; i++ {
		f := daftarFilm[i]
		fmt.Printf("  [%d] %s\n", i+1, f.judul)
		fmt.Printf("      Genre   : %s\n", f.genre)
		fmt.Printf("      Tahun   : %d\n", f.tahun)
		fmt.Printf("      Rating  : %.1f / 10\n", f.rating)
		fmt.Printf("      Sinopsis: %s\n", f.deskripsi)
		fmt.Println()
	}

	fmt.Printf("  Total: %d film\n", jumlahFilm)
	garis()
}

// main - titik mulai program
func main() {
	// tampilkan header waktu program pertama jalan
	fmt.Println()
	fmt.Println("  +++ CineReview +++")
	fmt.Println("  Aplikasi Katalog dan Rating Film")
	fmt.Println("  Tugas Besar - Algoritma Pemrograman 2")
	fmt.Println()

	// loop menu utama
	for {
		tampilMenu()
		pilihan := bacaInput("  Pilih menu: ")

		switch pilihan {
		case "1":
			tambahFilm()
		case "2":
			tampilSemuaFilm()
		case "0":
			fmt.Println()
			fmt.Println("  Terima kasih sudah menggunakan CineReview!")
			fmt.Println("  Program selesai.")
			fmt.Println()
			os.Exit(0)
		default:
			fmt.Println("  [!] Menu tidak tersedia, pilih 0-2.")
		}
	}
}
