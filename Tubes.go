package main

import "fmt"

const maxData = 1000000

type DataPolusi struct {
	Lokasi        string
	Tanggal       string
	AQI           int
	Sumber        string
	TingkatBahaya string
}

type TabPolusi [maxData]DataPolusi

func klasifikasiAQI(aqi int) string {
	if aqi <= 50 {
		return "Baik"
	} else if aqi <= 100 {
		return "Sedang"
	} else if aqi <= 150 {
		return "Tidak Sehat"
	} else {
		return "Berbahaya"
	}
}

func tambahData(data *TabPolusi, dataCount *int) {
	var lokasi, tanggal, sumber string
	var aqi int

	if *dataCount >= maxData {
		fmt.Println("Data penuh, tidak bisa menambah data baru.")
		return
	}

	fmt.Println("Masukkan lokasi: ")
	fmt.Scan(&lokasi)
	fmt.Println("Masukkan tanggal (YYYY-MM-DD): ")
	fmt.Scan(&tanggal)
	fmt.Println("Masukkan AQI: ")
	fmt.Scan(&aqi)
	fmt.Println("Masukkan sumber polusi: ")
	fmt.Scan(&sumber)

	data[*dataCount] = DataPolusi{lokasi, tanggal, aqi, sumber, klasifikasiAQI(aqi)}
	*dataCount++
	fmt.Println("Data berhasil ditambahkan.")
}

func tampilkanData(data *TabPolusi, dataCount int) {
	var i int
	var d DataPolusi
	fmt.Println("===== Daftar Data Polusi =====")
	for i = 0; i < dataCount; i++ {
		d = data[i]
		fmt.Printf("Lokasi = %s\n", d.Lokasi)
		fmt.Printf("Tanggal = %s\n", d.Tanggal)
		fmt.Printf("AQI = %d\n", d.AQI)
		fmt.Printf("Sumber = %s\n", d.Sumber)
		fmt.Printf("Tingkat Bahaya = %s\n", d.TingkatBahaya)
		fmt.Println("==============================")
	}
}

func ubahData(data *TabPolusi, dataCount int, tempat string) {
	var lokasi, tanggal, sumber string
	var aqi, i int
	var dataCocok bool
	dataCocok = false
	for i = 0; i < dataCount; i++ {
		if data[i].Lokasi == tempat {
			fmt.Print("Masukkan tanggal (YYYY-MM-DD): ")
			fmt.Scanln(&tanggal)
			fmt.Print("Masukkan AQI: ")
			fmt.Scanln(&aqi)
			fmt.Print("Masukkan sumber polusi: ")
			fmt.Scanln(&sumber)

			data[i] = DataPolusi{lokasi, tanggal, aqi, sumber, klasifikasiAQI(aqi)}
			dataCocok = true
			fmt.Println("Data berhasil diubah.")
			break
		}
	}

	if !dataCocok {
		fmt.Println("Data tidak ditemukan.")
	}
}

func hapusData(data *TabPolusi, dataCount *int, tempat string) {
	var i, j int
	var dataCocok bool
	dataCocok = false
	for i = 0; i < *dataCount; i++ {
		if data[i].Lokasi == tempat {
			for j = i; j < *dataCount-1; j++ {
				data[j] = data[j+1]
			}
			*dataCount--
			dataCocok = true
			fmt.Println("Data berhasil dihapus.")
			break
		}
	}

	if !dataCocok {
		fmt.Println("Data tidak ditemukan.")
	}
}

func cariLokasi(data *TabPolusi, dataCount int, tempat string) {
	var i int
	var dataCocok bool
	dataCocok = false
	for i = 0; i < dataCount; i++ {
		if data[i].Lokasi == tempat {
			fmt.Printf("Ditemukan: %s | AQI: %d | %s\n", data[i].Lokasi, data[i].AQI, data[i].TingkatBahaya)
			dataCocok = true
		}
	}
	if !dataCocok {
		fmt.Println("Data tidak ditemukan.")
	}
}

func descending(data *TabPolusi, dataCount int) {
	var i, j int
	for i = 0; i < dataCount; i++ {
		for j = i + 1; j < dataCount; j++ {
			if data[i].AQI < data[j].AQI {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	fmt.Println("Data diurutkan berdasarkan AQI tertinggi.")
}

func ascending(data *TabPolusi, dataCount int) {
	var i, j int
	for i = 0; i < dataCount; i++ {
		for j = i + 1; j < dataCount; j++ {
			if data[i].AQI > data[j].AQI {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	fmt.Println("Data diurutkan berdasarkan AQI terendah.")
}

func main() {
	var data TabPolusi
	var dataCount int
	var lokasi string
	var pilihan, edit, urut int

	for {
		fmt.Println("||================================================||")
		fmt.Println("||                                                ||")
		fmt.Println("||      Aplikasi Pemantauaan Polusi Udara         ||")
		fmt.Println("||                   by :                         ||")
		fmt.Println("||           Julian Wimajaya                      ||")
		fmt.Println("||           Raihan Wendra Baswara                ||")
		fmt.Println("||     Tugas Besar Algoritma Pemrograman 2025     ||")
		fmt.Println("||                                                ||")
		fmt.Println("||================================================||")
		fmt.Println("1. Tambah Data Polusi")
		fmt.Println("2. Tampilkan Semua Data")
		fmt.Println("3. Cari Lokasi")
		fmt.Println("4. Urutkan Data")
		fmt.Println("5. Keluar")
		fmt.Println("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahData(&data, &dataCount)
		case 2:
			tampilkanData(&data, dataCount)
			fmt.Println("1. Edit, 2. Hapus, 3. Kembali")
			fmt.Println("Pilih menu: ")
			fmt.Scan(&edit)
			switch edit {
			case 1:
				fmt.Println("Masukkan lokasi: ")
				fmt.Scan(&lokasi)
				ubahData(&data, dataCount, lokasi)
			case 2:
				fmt.Println("Masukkan lokasi: ")
				fmt.Scan(&lokasi)
				hapusData(&data, &dataCount, lokasi)
			case 3:
				fmt.Println("Kembali ke menu utama")
			default:
				fmt.Println("Pilihan tidak valid.")
			}
		case 3:
			fmt.Println("Masukkan lokasi: ")
			fmt.Scan(&lokasi)
			cariLokasi(&data, dataCount, lokasi)
		case 4:
			fmt.Println("1.Dari AQI Terendah , 2.Dari AQI Tertinggi, 3.Kembali")
			fmt.Println("Pilih menu: ")
			fmt.Scan(&urut)
			switch urut {
			case 1:
				ascending(&data, dataCount)
				tampilkanData(&data, dataCount)
			case 2:
				descending(&data, dataCount)
				tampilkanData(&data, dataCount)
			case 3:
				fmt.Println("Kembali ke menu utama")
			default:
				fmt.Println("Pilihan tidak valid.")
			}
		case 5:
			fmt.Println("Keluar dari aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
