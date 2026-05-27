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

// hapus film berdasarkan nomor urut
func hapusFilm() {
	fmt.Println()
	garis()
	fmt.Println("  >> HAPUS FILM")
	garis()

	if jumlahFilm == 0 {
		fmt.Println("  Tidak ada film yang bisa dihapus.")
		garis()
		return
	}

	// tampilkan dulu filmnya biar user tau nomor berapa
	for i := 0; i < jumlahFilm; i++ {
		fmt.Printf("  [%d] %s (%.1f)\n", i+1, daftarFilm[i].judul, daftarFilm[i].rating)
	}
	fmt.Println()

	nomor := bacaInt("  Masukkan nomor film yang dihapus: ")

	if nomor < 1 || nomor > jumlahFilm {
		fmt.Println("  [!] Nomor tidak valid.")
		garis()
		return
	}

	judulDihapus := daftarFilm[nomor-1].judul

	// geser elemen ke kiri buat nutupin yang dihapus
	for i := nomor - 1; i < jumlahFilm-1; i++ {
		daftarFilm[i] = daftarFilm[i+1]
	}
	daftarFilm = daftarFilm[:jumlahFilm-1]
	jumlahFilm--

	fmt.Printf("\n  [OK] Film \"%s\" berhasil dihapus.\n", judulDihapus)
	garis()
}

// sequential search cari film berdasarkan judul
func cariFilm() {
	fmt.Println()
	garis()
	fmt.Println("  >> CARI FILM (Sequential Search)")
	garis()

	if jumlahFilm == 0 {
		fmt.Println("  Tidak ada film untuk dicari.")
		garis()
		return
	}

	keyword := bacaInput("  Masukkan judul film: ")
	keyword = strings.ToLower(keyword)

	ketemu := false
	var hasil []int

	// loop dari awal sampai akhir, cek satu-satu
	for i := 0; i < jumlahFilm; i++ {
		judulLower := strings.ToLower(daftarFilm[i].judul)
		if strings.Contains(judulLower, keyword) {
			hasil = append(hasil, i)
			ketemu = true
		}
	}

	fmt.Println()
	if !ketemu {
		fmt.Printf("  Film dengan kata kunci \"%s\" tidak ditemukan.\n", keyword)
	} else {
		fmt.Printf("  Ditemukan %d hasil:\n\n", len(hasil))
		for _, idx := range hasil {
			f := daftarFilm[idx]
			fmt.Printf("  - %s | %s | %d | Rating: %.1f\n", f.judul, f.genre, f.tahun, f.rating)
			fmt.Printf("    %s\n", f.deskripsi)
			fmt.Println()
		}
	}
	garis()
}

// selection sort urutkan film dari rating tertinggi ke terendah
func selectionSort() {
	fmt.Println()
	garis()
	fmt.Println("  >> URUTKAN FILM (Selection Sort by Rating)")
	garis()

	if jumlahFilm == 0 {
		fmt.Println("  Tidak ada film untuk diurutkan.")
		garis()
		return
	}

	// proses selection sort - cari yang paling besar terus tukar
	for i := 0; i < jumlahFilm-1; i++ {
		idxMax := i
		for j := i + 1; j < jumlahFilm; j++ {
			if daftarFilm[j].rating > daftarFilm[idxMax].rating {
				idxMax = j
			}
		}

		// tukar posisi data film
		if idxMax != i {
			tmp := daftarFilm[i]
			daftarFilm[i] = daftarFilm[idxMax]
			daftarFilm[idxMax] = tmp
		}
	}

	fmt.Println("  [OK] Film berhasil diurutkan dari rating tertinggi!")
	fmt.Println()

	// tampilkan hasil sorting
	for i := 0; i < jumlahFilm; i++ {
		f := daftarFilm[i]
		fmt.Printf("  %d. %-25s %.1f/10\n", i+1, f.judul, f.rating)
	}
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
