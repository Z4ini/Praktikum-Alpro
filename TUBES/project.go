package main

import (
	"fmt"
	"strings"
	"time"
)

const NMAX int = 1000 // batas jumlah array

type bahan struct {
	nama          string 
	status        string
	jumlah        int
	inDate        time.Time
	kadaluarsa    int       // dalam jumlah hari sejak inDate (dalam bentuk hari)
	tglKadaluarsa time.Time // tanggal kadaluarsa dalam bentuk tanggal
}

// mendeklarasikan array untuk bahan
type tabBahan [NMAX]bahan

func main() {
	var data tabBahan
	var p, n int

	dumyData(&data, &n) // untuk bahan test

	for {
		menu() // memanggil fungsi menu
		fmt.Scan(&p)

		switch p {
		case 1:
			show(data, n)
		case 2:
			input(&data, &n)
		case 3:
			update(&data, n)
		case 4:
			delete(&data, &n)
		case 5:
			fmt.Println("Keluar dari program")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

// untuk menampilkan menu
func menu() {
	fmt.Println("\n===== MENU =====")
	fmt.Println("1. INFO")
	fmt.Println("2. Input")
	fmt.Println("3. Update")
	fmt.Println("4. Delete")
	fmt.Println("5. Keluar")
	fmt.Print("Masukkan Pilihan: ")
}

// untuk membuat data baru
func input(A *tabBahan, n *int) {
	var jumlah, kadaluarsa int
	var nama string

	fmt.Println("Masukkan data bahan (nama jumlah kadaluarsa), ketik 'none' untuk berhenti:")

	fmt.Scan(&nama)

	for nama != "none" && *n <= NMAX {
		fmt.Scan(&jumlah, &kadaluarsa)
		A[*n].nama = nama
		A[*n].jumlah = jumlah
		A[*n].kadaluarsa = kadaluarsa
		A[*n].inDate = time.Now()
		A[*n].tglKadaluarsa = A[*n].inDate.AddDate(0, 0, A[*n].kadaluarsa)
		A[*n].status = cekStatus(A[*n])

		*n++
		fmt.Scan(&nama)
	}
}

// menampilkan data
func show(A tabBahan, n int) {
	fmt.Println("\nTanggal Sekarang: ", time.Now().Format("02-01-2006"))
	fmt.Printf("\n%-4s %-15s %-8s %-15s %-20s\n", "NO", "Nama", "Jumlah", "Kadaluarsa", "Status")
	fmt.Println(strings.Repeat("-", 65)) // garis pembatas

	for i := 0; i < n; i++ {
		if A[i].nama != "" {
			fmt.Printf(
				"%-4d %-15s %-8d %-15s %-20s\n",
				i+1,
				A[i].nama,
				A[i].jumlah,
				A[i].tglKadaluarsa.Format("02-01-2006"),
				A[i].status,
			)
		}
	}
}

// untuk mengedit data yang sudah ada
func update(A *tabBahan, n int) {
	var p, jumlahBaru, kadaluarsaBaru int
	var namaBaru string

	show(*A, n)

	fmt.Print("Edit Data ke? ")
	fmt.Scan(&p)

	if p < 1 || p > n {
		fmt.Println("Data tidak ditemukan")
		return
	}

	fmt.Scan(&namaBaru, &jumlahBaru, &kadaluarsaBaru)
	A[p-1].nama = namaBaru
	A[p-1].jumlah = jumlahBaru
	A[p-1].kadaluarsa = kadaluarsaBaru
	A[p-1].inDate = time.Now()
	A[p-1].tglKadaluarsa = A[p-1].inDate.AddDate(0, 0, A[p-1].kadaluarsa)
	A[p-1].status = cekStatus(A[p-1])
}

// untuk menghapus data
func delete(A *tabBahan, n *int) {
	var i, p int

	show(*A, *n)

	fmt.Print("Hapus Data ke? ")
	fmt.Scan(&p)

	if p < 1 || p > *n {
		fmt.Println("Data tidak ditemukan")
		return
	}

	for i = p - 1; i < *n-1; i++ {
		A[i] = A[i+1]
	}

	*n--
	fmt.Println("Data berhasil di hapus")
}

// untuk mengecek apakah akan kadaluarsa atau tidak
func cekStatus(B bahan) string {
	var sisa int

	sisa = int(B.tglKadaluarsa.Sub(time.Now()).Hours() / 24)

	if sisa < 0 {
		return "Sudah Kadaluarsa"
	} else if sisa <= 1 {
		return "Segera Kadaluarsa"
	} else if sisa <= 3 {
		return "Akan Kadaluarsa"
	} else {
		return "Aman"
	}
}

// Dumy data
func dumyData(A *tabBahan, n *int) {
	now := time.Now()

	A[0].nama = "Wortel"
	A[0].jumlah = 2
	A[0].kadaluarsa = 20
	A[0].inDate = now
	A[0].tglKadaluarsa = A[0].inDate.AddDate(0, 0, A[0].kadaluarsa)
	A[0].status = cekStatus(A[0])

	A[1].nama = "Kentang"
	A[1].jumlah = 5
	A[1].kadaluarsa = 25
	A[1].inDate = now
	A[1].tglKadaluarsa = A[1].inDate.AddDate(0, 0, A[1].kadaluarsa)
	A[1].status = cekStatus(A[1])

	A[2].nama = "Tomat"
	A[2].jumlah = 3
	A[2].kadaluarsa = 2 // akan segera kadaluarsa
	A[2].inDate = now
	A[2].tglKadaluarsa = A[2].inDate.AddDate(0, 0, A[2].kadaluarsa)
	A[2].status = cekStatus(A[2])

	A[3].nama = "Daging Ayam"
	A[3].jumlah = 1
	A[3].kadaluarsa = 1
	A[3].inDate = now.AddDate(0, 0, -2) // sudah lewat
	A[3].tglKadaluarsa = A[3].inDate.AddDate(0, 0, A[3].kadaluarsa)
	A[3].status = cekStatus(A[3])

	A[4].nama = "Susu"
	A[4].jumlah = 2
	A[4].kadaluarsa = 3 // akan kadaluarsa
	A[4].inDate = now
	A[4].tglKadaluarsa = A[4].inDate.AddDate(0, 0, A[4].kadaluarsa)
	A[4].status = cekStatus(A[4])

	*n = 5
}
