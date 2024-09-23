package main

import (
	"fmt"
	"os"
	"os/exec" // Mengimpor paket kustom
	"sort"
	"strconv"
)

// clearScreen adalah fungsi untuk membersihkan layar terminal
func clearScreen() {
	// Membuat perintah sistem untuk membersihkan layar pada Windows
	// 'cmd' adalah command prompt di Windows, '/c' mengeksekusi perintah berikutnya, yaitu 'cls' untuk membersihkan layar.
	cmd := exec.Command("cmd", "/c", "cls")

	// Mengarahkan output dari perintah 'cmd' ke stdout (layar terminal)
	cmd.Stdout = os.Stdout

	// Menjalankan perintah yang telah dibuat di atas
	cmd.Run()
}

// sortStrings adalah fungsi variadic yang menerima sejumlah string dan mengembalikan slice yang terurut.
func sortStrings(order string, strings ...string) []string {
	// Menggunakan sort.Slice untuk mengurutkan berdasarkan order yang diinginkan
	sort.Slice(strings, func(i, j int) bool {
		if order == "asc" {
			return strings[i] < strings[j] // a-z
		}
		return strings[i] > strings[j] // z-a
	})
	return strings
}

// Fungsi untuk menyimpan data ke dalam map
func storeMap(maps map[string]int, str string, number int) {
	maps[str] = number
}

// Tekan enter adalah fungsi untuk menahan user pada tampilan tertentu sementara
func tekanEnter() {
	// Menampilkan pesan untuk menekan Enter
	fmt.Println("Tekan Enter untuk melanjutkan...")
	fmt.Scanln() // Menunggu hingga pengguna menekan Enter
}

// Variadic function
func operasiMatematika(bil1, bil2 float64) []float64 {
	hasil := make([]float64, 4) // Slice untuk menyimpan hasil operasi

	// Penjumlahan
	hasil[0] = bil1 + bil2

	// Pengurangan
	hasil[1] = bil1 - bil2

	// Perkalian
	hasil[2] = bil1 * bil2

	// Pembagian
	if bil2 != 0 {
		hasil[3] = bil1 / bil2
	} else {
		hasil[3] = 0 // Not a Number untuk hasil yang tidak valid
	}

	return hasil
}
func filterUmurUser(minUmur int, users map[string]int) {
	fmt.Printf("Pengguna dengan usia lebih dari %d:\n", minUmur)
	fmt.Println("=========================================================")
	// Mendefinisikan anonymous function untuk memfilter usia
	filter := func(umur int) bool {
		return umur >= minUmur
	}

	// Mengiterasi map 'users' dan menerapkan fungsi filter
	for nama, umur := range users {
		if filter(umur) { // Memeriksa apakah usia lebih besar dari 'minUmur'
			fmt.Printf("Nama: %s, Umur: %d\n", nama, umur)
		}
	}
}

// Faktorial menghitung faktorial dari n secara rekursif
func faktorial(n int) int {
	if n == 0 {
		return 1 // Basis: faktorial dari 0 adalah 1
	}
	return n * faktorial(n-1) // Rekursi
}

// Fibonacci menghitung bilangan Fibonacci ke-n secara rekursif
func fibonacci(n int) int {
	if n <= 0 {
		return 0 // Basis: Fibonacci(0) adalah 0
	} else if n == 1 {
		return 1 // Basis: Fibonacci(1) adalah 1
	}
	return fibonacci(n-1) + fibonacci(n-2) // Rekursi
}

// main adalah fungsi utama yang akan dijalankan
func main() {
	// Deklarasi dari setiap variable
	var num1, num2 string
	var pilihan int
	var n int     // Deklarasi variable yang akan digunakan untuk fibonacci dan faktorial
	var err error // Variable untuk menanggani error
	// Menampilkan menu utama beserta pilihannya
mainMenu: // Label untuk savepoint
	clearScreen() // Membersihkan layar
	fmt.Println("======================================")
	fmt.Println("\t      Menu Utama")
	fmt.Println("======================================")
	fmt.Println("1. Menampilkan Hello World!")
	fmt.Println("2. Kalkulator Sederhana")
	fmt.Println("3. Simpan dan Tampilkan Data Pengguna")
	fmt.Println("4. Hitung Faktorial")
	fmt.Println("5. Deret Fibonacci")
	fmt.Println("6. Sorting")
	fmt.Println("7. Exit")
	fmt.Println("======================================")
	fmt.Print("Pilih Menu (1/2/3/4/5/6/7): ")
	fmt.Scanln(&pilihan) // Menginput data ke dalam variable
	fmt.Println("======================================")
	// Membuat switch case untuk memilih menu
	switch pilihan {
	case 1:
		fmt.Println("Hello World!") //Menampilkan Hello World!
		tekanEnter()
		goto mainMenu // Kembali ke savepoint (Menu Utama)
	case 2:
		for {
			var bil1, bil2 float64 // Deklarasi variabel ke-1 dan ke-2
			// Tampilan Menu kalkulator
			fmt.Print("Masukkan bilangan ke-1: ")
			fmt.Scanln(&num1) // Menginput data ke dalam Array
			//Mengecek inputan apakah berupa angka
			bil1, err = strconv.ParseFloat(num1, 64)
			if err != nil {
				fmt.Println("Data tidak valid, inputan harus angka.")
				tekanEnter()
				clearScreen() // Membersihkan layar
				continue
			}
			fmt.Print("Masukkan bilangan ke-2: ")
			fmt.Scanln(&num2) // Menginput data ke dalam Array
			fmt.Println("======================================")
			//Mengecek inputan apakah berupa angka
			bil2, err = strconv.ParseFloat(num2, 64)
			if err != nil {
				fmt.Println("Data tidak valid, inputan harus angka.")
				tekanEnter()
				clearScreen() // Membersihkan layar
				continue
			}
			hasil := operasiMatematika(bil1, bil2)
			fmt.Printf("Penjumlahan (+): %.2f + %.2f = %.2f\n", bil1, bil2, hasil[0])
			fmt.Printf("Pengurangan (-): %.2f - %.2f = %.2f\n", bil1, bil2, hasil[1])
			fmt.Printf("Perkalian   (x): %.2f x %.2f = %.2f\n", bil1, bil2, hasil[2])
			if hasil[3] != 0 {
				fmt.Printf("Pembagian   (/): %.2f / %.2f = %.2f\n", bil1, bil2, hasil[3])
			} else {
				fmt.Println("Pembagian   (/): tidak dapat dilakukan karena pembagi adalah 0")
			}

			tekanEnter()
			goto mainMenu // Kembali ke savepoint (Menu Utama)
		}
	case 3:
		users := make(map[string]int)
		var nama, umurStr string
		var minUmur int

		// Tampilan Menu Simpan dan Tampilkan Data Pengguna
		for {
			fmt.Print("Masukkan nama pengguna (atau ketik 'selesai' untuk berhenti): ")
			fmt.Scanln(&nama)
			if nama == "selesai" {
				fmt.Println("=========================================================")
				fmt.Print("Masukkan minimum umur pengguna yang ingin ditampilkan: ")
				fmt.Scanln(&minUmur)
				filterUmurUser(minUmur, users)
				tekanEnter()
				goto mainMenu // Kembali ke savepoint (Menu Utama)
			}
			fmt.Print("Masukkan umur pengguna: ")
			fmt.Scanln(&umurStr)
			umur, err := strconv.Atoi(umurStr)
			if err != nil {
				fmt.Println("Umur tidak valid, harus berupa angka.")
				continue
			}
			storeMap(users, nama, umur)
			fmt.Println("=========================================================")
			fmt.Println("Data berhasil disimpan: ")
			fmt.Printf("Nama: %s Umur: %d\n", nama, users[nama])
			fmt.Println("=========================================================")

		}
	case 4:
		for {
			fmt.Print("Masukkan bilangan untuk menghitung faktorial: ")
			fmt.Scanln(&num1)
			//Mengecek inputan apakah berupa angka
			n, err = strconv.Atoi(num1)
			if err != nil {
				fmt.Println("Data tidak valid, inputan harus angka.")
				tekanEnter()
				clearScreen() // Membersihkan layar
				continue
			}
			fmt.Printf("Faktorial dari %d adalah %d\n", n, faktorial(n))
			tekanEnter()
			goto mainMenu // Kembali ke savepoint (Menu Utama)
		}
	case 5:
		for {
			fmt.Print("Masukkan bilangan untuk Fibonacci: ")
			fmt.Scanln(&num1)
			//Mengecek inputan apakah berupa angka
			n, err = strconv.Atoi(num1)
			if err != nil {
				fmt.Println("Data tidak valid, inputan harus angka.")
				tekanEnter()
				clearScreen() // Membersihkan layar
				continue
			}
			fmt.Printf("Bilangan Fibonacci ke-%d adalah %d\n", n, fibonacci(n))
			tekanEnter()
			goto mainMenu // Kembali ke savepoint (Menu Utama)
		}
	case 6:
		fmt.Println("Kalimat sebelum disorting:")
		fmt.Println("banana", "apple", "orange", "grape")
		// Mengurutkan secara ascending
		stringsAscending := sortStrings("asc", "banana", "apple", "orange", "grape")
		fmt.Println("Bilangan terurut (a-z):", stringsAscending)

		stringsDescending := sortStrings("desc", "banana", "apple", "orange", "grape")
		fmt.Println("Bilangan terurut (z-a):", stringsDescending)
		tekanEnter()
		goto mainMenu // Kembali ke savepoint (Menu Utama)
	case 7:
		return // Keluar dari program
	default:
		fmt.Println("Input tidak valid, silakan masukkan pilihan yang benar.")
		tekanEnter()
		goto mainMenu // Kembali ke savepoint (Menu Utama)

	}
}
