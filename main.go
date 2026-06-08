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

// kapasitas maksimum film dan genre yang bisa disimpan
const MAX_FILM = 100
const MAX_GENRE = 50

// variabel global daftar film dan jumlahnya (array statis)
var daftarFilm [MAX_FILM]Film
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
	fmt.Println("  [3] Hapus Film")
	fmt.Println("  [4] Cari Film (Sequential Search)")
	fmt.Println("  [5] Urutkan Film by Rating (Selection Sort)")
	fmt.Println("  [6] Edit Data Film")
	fmt.Println("  [7] Cari Film (Binary Search)")
	fmt.Println("  [8] Urutkan Film by Tahun (Insertion Sort)")
	fmt.Println("  [9] Statistik Film")
	fmt.Println("  [0] Keluar")
	fmt.Println("========================================")
}

// fungsi buat nambahin film baru ke array statis
func tambahFilm() {
	fmt.Println()
	garis()
	fmt.Println("  >> TAMBAH FILM BARU")
	garis()

	if jumlahFilm >= MAX_FILM {
		fmt.Printf("  [!] Daftar film sudah penuh (maks %d film).\n", MAX_FILM)
		garis()
		return
	}

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

	// masukkan film ke array statis
	daftarFilm[jumlahFilm] = f
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
	// kosongkan elemen terakhir dan kurangi jumlah
	daftarFilm[jumlahFilm-1] = Film{}
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
	var hasil [MAX_FILM]int
	jumlahHasil := 0

	// cari film satu per satu dari awal sampai akhir
	for i := 0; i < jumlahFilm; i++ {
		judulLower := strings.ToLower(daftarFilm[i].judul)
		if strings.Contains(judulLower, keyword) {
			hasil[jumlahHasil] = i
			jumlahHasil++
			ketemu = true
		}
	}

	fmt.Println()
	if !ketemu {
		fmt.Printf("  Film dengan kata kunci \"%s\" tidak ditemukan.\n", keyword)
	} else {
		fmt.Printf("  Ditemukan %d hasil:\n\n", jumlahHasil)
		for i := 0; i < jumlahHasil; i++ {
			f := daftarFilm[hasil[i]]
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

// edit data film yang sudah ada
func editFilm() {
	fmt.Println()
	garis()
	fmt.Println("  >> EDIT DATA FILM")
	garis()

	if jumlahFilm == 0 {
		fmt.Println("  Tidak ada film yang bisa diedit.")
		garis()
		return
	}

	// tampilkan dulu semua film biar user tau mau edit nomor berapa
	for i := 0; i < jumlahFilm; i++ {
		fmt.Printf("  [%d] %s (%d) - %.1f/10\n", i+1, daftarFilm[i].judul, daftarFilm[i].tahun, daftarFilm[i].rating)
	}
	fmt.Println()

	nomor := bacaInt("  Masukkan nomor film yang ingin diedit: ")

	if nomor < 1 || nomor > jumlahFilm {
		fmt.Println("  [!] Nomor tidak valid.")
		garis()
		return
	}

	// ambil data film yang mau diedit
	idx := nomor - 1
	fmt.Printf("\n  Editing film: %s\n", daftarFilm[idx].judul)
	fmt.Println("  (Kosongkan input untuk tidak mengubah)")
	fmt.Println()

	// edit judul
	judulBaru := bacaInput("  Judul baru     : ")
	if judulBaru != "" {
		daftarFilm[idx].judul = judulBaru
	}

	// edit genre
	genreBaru := bacaInput("  Genre baru     : ")
	if genreBaru != "" {
		daftarFilm[idx].genre = genreBaru
	}

	// edit tahun - kalau diisi, coba parsing
	tahunStr := bacaInput("  Tahun baru     : ")
	if tahunStr != "" {
		tahunBaru, err := strconv.Atoi(tahunStr)
		if err != nil {
			fmt.Println("  [!] Tahun tidak valid, tahun tidak diubah.")
		} else {
			daftarFilm[idx].tahun = tahunBaru
		}
	}

	// edit deskripsi
	deskripsiBaru := bacaInput("  Deskripsi baru : ")
	if deskripsiBaru != "" {
		daftarFilm[idx].deskripsi = deskripsiBaru
	}

	// edit rating - kalau diisi, coba parsing
	ratingStr := bacaInput("  Rating baru    : ")
	if ratingStr != "" {
		ratingBaru, err := strconv.ParseFloat(ratingStr, 64)
		if err != nil || ratingBaru < 0 || ratingBaru > 10 {
			fmt.Println("  [!] Rating tidak valid (harus 0-10), rating tidak diubah.")
		} else {
			daftarFilm[idx].rating = ratingBaru
		}
	}

	fmt.Printf("\n  [OK] Data film \"%s\" berhasil diperbarui!\n", daftarFilm[idx].judul)
	garis()
}

// urutin dulu data berdasarkan judul sebelum binary search
// pakai selection sort buat judul (ascending)
func urutJudul() {
	for i := 0; i < jumlahFilm-1; i++ {
		idxMin := i
		for j := i + 1; j < jumlahFilm; j++ {
			// bandingkan judul huruf kecil semua
			if strings.ToLower(daftarFilm[j].judul) < strings.ToLower(daftarFilm[idxMin].judul) {
				idxMin = j
			}
		}
		// tukar kalau ketemu yang lebih kecil
		if idxMin != i {
			tmp := daftarFilm[i]
			daftarFilm[i] = daftarFilm[idxMin]
			daftarFilm[idxMin] = tmp
		}
	}
}

// binary search cari film berdasarkan judul (harus sudah urut dulu)
func binarySearch() {
	fmt.Println()
	garis()
	fmt.Println("  >> CARI FILM (Binary Search)")
	garis()

	if jumlahFilm == 0 {
		fmt.Println("  Tidak ada film untuk dicari.")
		garis()
		return
	}

	// urutkan dulu berdasarkan judul sebelum binary search
	fmt.Println("  [*] Mengurutkan data berdasarkan judul...")
	urutJudul()

	keyword := bacaInput("  Masukkan judul film (harus sama persis): ")
	keywordLower := strings.ToLower(keyword)

	// proses binary search - cari di tengah-tengah terus
	kiri := 0
	kanan := jumlahFilm - 1
	ketemu := -1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		judulTengah := strings.ToLower(daftarFilm[tengah].judul)

		if judulTengah == keywordLower {
			// ketemu!
			ketemu = tengah
			break
		} else if judulTengah < keywordLower {
			// cari di sebelah kanan
			kiri = tengah + 1
		} else {
			// cari di sebelah kiri
			kanan = tengah - 1
		}
	}

	fmt.Println()
	if ketemu == -1 {
		fmt.Printf("  Film \"%s\" tidak ditemukan.\n", keyword)
		fmt.Println("  (Pastikan judul ditulis sama persis)")
	} else {
		f := daftarFilm[ketemu]
		fmt.Println("  Film ditemukan!")
		fmt.Println()
		fmt.Printf("  Judul   : %s\n", f.judul)
		fmt.Printf("  Genre   : %s\n", f.genre)
		fmt.Printf("  Tahun   : %d\n", f.tahun)
		fmt.Printf("  Rating  : %.1f / 10\n", f.rating)
		fmt.Printf("  Sinopsis: %s\n", f.deskripsi)
	}
	garis()
}

// insertion sort urutkan film berdasarkan tahun rilis dari yang paling lama
func insertionSort() {
	fmt.Println()
	garis()
	fmt.Println("  >> URUTKAN FILM (Insertion Sort by Tahun)")
	garis()

	if jumlahFilm == 0 {
		fmt.Println("  Tidak ada film untuk diurutkan.")
		garis()
		return
	}

	// proses insertion sort - ambil satu-satu terus sisipkan ke posisi yang bener
	for i := 1; i < jumlahFilm; i++ {
		kunci := daftarFilm[i]
		j := i - 1

		// geser elemen yang lebih besar ke kanan
		for j >= 0 && daftarFilm[j].tahun > kunci.tahun {
			daftarFilm[j+1] = daftarFilm[j]
			j--
		}

		// sisipkan kunci ke posisi yang tepat
		daftarFilm[j+1] = kunci
	}

	fmt.Println("  [OK] Film berhasil diurutkan dari tahun terlama!")
	fmt.Println()

	// tampilkan hasil insertion sort
	for i := 0; i < jumlahFilm; i++ {
		f := daftarFilm[i]
		fmt.Printf("  %d. %-25s %d  (%.1f/10)\n", i+1, f.judul, f.tahun, f.rating)
	}
	garis()
}

// tampilkan statistik semua film
func statistikFilm() {
	fmt.Println()
	garis()
	fmt.Println("  >> STATISTIK FILM")
	garis()

	if jumlahFilm == 0 {
		fmt.Println("  Belum ada data film.")
		garis()
		return
	}

	// hitung rata-rata rating
	totalRating := 0.0
	for i := 0; i < jumlahFilm; i++ {
		totalRating += daftarFilm[i].rating
	}
	rataRata := totalRating / float64(jumlahFilm)

	fmt.Printf("  Total Film      : %d film\n", jumlahFilm)
	fmt.Printf("  Rata-rata Rating: %.2f / 10\n", rataRata)
	fmt.Println()

	// hitung jumlah film tiap genre
	// pakai dua array statis paralel: satu buat nama genre, satu buat jumlahnya
	var namaGenre [MAX_GENRE]string
	var jumlahPerGenre [MAX_GENRE]int
	jumlahGenre := 0

	for i := 0; i < jumlahFilm; i++ {
		genre := strings.ToLower(daftarFilm[i].genre)
		ditemukan := false

		// cek apakah genre ini sudah ada di daftar
		for j := 0; j < jumlahGenre; j++ {
			if namaGenre[j] == genre {
				jumlahPerGenre[j]++
				ditemukan = true
				break
			}
		}

		// kalau genre baru, tambahkan ke daftar array
		if !ditemukan && jumlahGenre < MAX_GENRE {
			namaGenre[jumlahGenre] = genre
			jumlahPerGenre[jumlahGenre] = 1
			jumlahGenre++
		}
	}

	// tampilkan jumlah film per genre
	fmt.Println("  Jumlah Film per Genre:")
	for i := 0; i < jumlahGenre; i++ {
		fmt.Printf("    - %-15s : %d film\n", namaGenre[i], jumlahPerGenre[i])
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
		case "3":
			hapusFilm()
		case "4":
			cariFilm()
		case "5":
			selectionSort()
		case "6":
			editFilm()
		case "7":
			binarySearch()
		case "8":
			insertionSort()
		case "9":
			statistikFilm()
		case "0":
			fmt.Println()
			fmt.Println("  Terima kasih sudah menggunakan CineReview!")
			fmt.Println("  Program selesai.")
			fmt.Println()
			os.Exit(0)
		default:
			fmt.Println("  [!] Menu tidak tersedia, pilih 0-9.")
		}
	}
}
