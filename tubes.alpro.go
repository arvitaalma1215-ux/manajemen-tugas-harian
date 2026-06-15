package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const KAPASITAS_MAKSIMAL_TUGAS = 100

type Tugas struct {
	MataKuliah string
	Kategori   string
	Deadline   string
	Jam        string
	Status     bool
}

type KumpulanDataTugas [KAPASITAS_MAKSIMAL_TUGAS]Tugas

var daftarTugasUtama KumpulanDataTugas

func main() {
	var pilihanMenu int
	var masihBerjalan bool
	pembacaInput := bufio.NewReader(os.Stdin)

	masihBerjalan = true
	for masihBerjalan {
		fmt.Println("\n=== APLIKASI MANAJEMEN TUGAS HARIAN (KHUSUS JUNI 2026) ===")
		fmt.Println("1. Tambah Tugas")
		fmt.Println("2. Cari Tugas")
		fmt.Println("3. Ubah Tugas")
		fmt.Println("4. Hapus Tugas")
		fmt.Println("5. Urutkan Tugas")
		fmt.Println("6. Ubah Status Tugas")
		fmt.Println("7. Tampilkan Semua Tugas")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		fmt.Scanf("%d\n", &pilihanMenu)

		if pilihanMenu == 1 {
			tambahTugas(pembacaInput)
		} else if pilihanMenu == 2 {
			cariTugas()
		} else if pilihanMenu == 3 {
			ubahTugas(pembacaInput)
		} else if pilihanMenu == 4 {
			hapusTugas()
		} else if pilihanMenu == 5 {
			urutkanTugas()
		} else if pilihanMenu == 6 {
			ubahStatusTugas()
		} else if pilihanMenu == 7 {
			tampilkanSemuaTugas()
		} else if pilihanMenu == 0 {
			fmt.Println("Terima kasih! Tetap semangat menyelesaikan tugasmu.")
			masihBerjalan = false // variabel lokal, bukan global
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi!")
		}
	}
}

func hitungJumlahTugas() int {
	var jumlahTugasTersimpan int
	var indeks int

	jumlahTugasTersimpan = 0
	indeks = 0
	for indeks < KAPASITAS_MAKSIMAL_TUGAS {
		if daftarTugasUtama[indeks].MataKuliah != "" {
			jumlahTugasTersimpan++
		}
		indeks++
	}
	return jumlahTugasTersimpan
}

func validasiTanggal(tanggalInput string) bool {
	var tahun, bulan, hari int

	if len(tanggalInput) != 10 || tanggalInput[4] != '-' || tanggalInput[7] != '-' {
		return false
	}

	if _, err := fmt.Sscanf(tanggalInput, "%4d-%2d-%2d", &tahun, &bulan, &hari); err != nil {
		return false
	}

	if tahun != 2026 || bulan != 6 {
		fmt.Println("Batasan: Aplikasi ini hanya menerima input untuk bulan Juni 2026!")
		return false
	}

	if hari < 1 || hari > 30 {
		return false
	}

	return true
}

func validasiJam(jamInput string) bool {
	var jam, menit int

	if len(jamInput) != 5 || jamInput[2] != ':' {
		return false
	}

	fmt.Sscanf(jamInput, "%2d:%2d", &jam, &menit)
	if jam < 0 || jam > 23 || menit < 0 || menit > 59 {
		return false
	}
	return true
}

func hitungSisaHari(tanggalDeadline string) int {
	var tahun, bulan, hari int

	fmt.Sscanf(tanggalDeadline, "%4d-%2d-%2d", &tahun, &bulan, &hari)
	waktuSekarang := time.Now()
	tanggalHariIni := waktuSekarang.Day()

	return hari - tanggalHariIni
}

func berikanFeedback(tugasBaru Tugas) {
	var sisaHari int

	fmt.Println("\nData berhasil diproses!")
	sisaHari = hitungSisaHari(tugasBaru.Deadline)

	if sisaHari < 0 {
		fmt.Println("Peringatan: Deadline tugas ini sudah lewat di bulan ini!")
	} else if sisaHari <= 2 {
		fmt.Printf("Bahaya: Tugas %s mendekati deadline (%d hari lagi)! Prioritaskan tugas ini.\n", tugasBaru.MataKuliah, sisaHari)
	} else {
		fmt.Printf("Info: Waktu pengerjaan masih aman (%d hari lagi). Cicil sedikit demi sedikit!\n", sisaHari)
	}

	if strings.EqualFold(tugasBaru.Kategori, "Tubes") || strings.EqualFold(tugasBaru.Kategori, "UAS") {
		fmt.Println("Rekomendasi: Kategori Tubes/UAS membutuhkan fokus tinggi. Jangan ditunda sampai malam terakhir!")
	}
}

func tambahTugas(pembacaInput *bufio.Reader) {
	var jumlahTugasSaatIni int
	var mataKuliahInput, kategoriInput string
	var tanggalDeadlineInput, jamPengumpulanInput string
	var apakahTanggalValid, apakahJamValid bool

	jumlahTugasSaatIni = hitungJumlahTugas()
	if jumlahTugasSaatIni >= KAPASITAS_MAKSIMAL_TUGAS {
		fmt.Println("Daftar tugas sudah penuh.")
		return
	}

	fmt.Print("Masukkan Mata Kuliah: ")
	mataKuliahInput, _ = pembacaInput.ReadString('\n')
	mataKuliahInput = strings.TrimSpace(mataKuliahInput)

	fmt.Print("Masukkan Kategori (Tubes/UAS/Kuis/Lainnya): ")
	kategoriInput, _ = pembacaInput.ReadString('\n')
	kategoriInput = strings.TrimSpace(kategoriInput)

	apakahTanggalValid = false
	for !apakahTanggalValid {
		fmt.Print("Masukkan Deadline (YYYY-MM-DD): ")
		fmt.Scanf("%s\n", &tanggalDeadlineInput)
		if validasiTanggal(tanggalDeadlineInput) {
			apakahTanggalValid = true
		} else {
			fmt.Println("Format tanggal salah, tidak ada, atau di luar Juni 2026. Coba lagi.")
		}
	}

	apakahJamValid = false
	for !apakahJamValid {
		fmt.Print("Masukkan Jam Pengumpulan (HH:MM): ")
		fmt.Scanf("%s\n", &jamPengumpulanInput)
		if validasiJam(jamPengumpulanInput) {
			apakahJamValid = true
		} else {
			fmt.Println("Format jam salah. Coba lagi.")
		}
	}

	daftarTugasUtama[jumlahTugasSaatIni] = Tugas{
		MataKuliah: mataKuliahInput,
		Kategori:   kategoriInput,
		Deadline:   tanggalDeadlineInput,
		Jam:        jamPengumpulanInput,
		Status:     false,
	}

	berikanFeedback(daftarTugasUtama[jumlahTugasSaatIni])
}

func cetakSatuTugas(tugasYgDicetak Tugas, nomorUrut int) {
	var statusString string

	statusString = "Belum Selesai"
	if tugasYgDicetak.Status {
		statusString = "Selesai"
	}
	fmt.Printf("%d. [%s] %s | Deadline: %s Pukul %s | Status: %s\n",
		nomorUrut+1, tugasYgDicetak.Kategori, tugasYgDicetak.MataKuliah,
		tugasYgDicetak.Deadline, tugasYgDicetak.Jam, statusString)
}

// tampilkanSemuaTugas mencetak seluruh data tugas yang tersimpan
func tampilkanSemuaTugas() {
	var totalTugas, indeks int

	totalTugas = hitungJumlahTugas()
	if totalTugas == 0 {
		fmt.Println("Belum ada data tugas.")
		return
	}

	fmt.Println("\n=== DAFTAR SEMUA TUGAS ===")
	indeks = 0
	for indeks < totalTugas {
		cetakSatuTugas(daftarTugasUtama[indeks], indeks)
		indeks++
	}
}

func cariTugas() {
	var totalTugas, pilihanMetodeCari int

	totalTugas = hitungJumlahTugas()
	if totalTugas == 0 {
		fmt.Println("Belum ada data tugas.")
		return
	}

	fmt.Println("\nCari Berdasarkan:")
	fmt.Println("1. Kategori (Sequential Search)")
	fmt.Println("2. Deadline (Binary Search)")
	fmt.Println("3. Tanggal & Jam Spesifik (Sequential Search)")
	fmt.Print("Pilihan: ")
	fmt.Scanf("%d\n", &pilihanMetodeCari)

	if pilihanMetodeCari == 1 {
		cariByKategoriSequential(totalTugas)
	} else if pilihanMetodeCari == 2 {
		cariByDeadlineBinary(totalTugas)
	} else if pilihanMetodeCari == 3 {
		cariByTanggalJamSequential(totalTugas)
	} else {
		fmt.Println("Pilihan pencarian tidak valid.")
	}
}

func cariByKategoriSequential(totalTugas int) {
	var kategoriDicari string
	var apakahDitemukan bool
	var indeks int

	fmt.Print("Masukkan kategori yang dicari: ")
	fmt.Scanf("%s\n", &kategoriDicari)

	apakahDitemukan = false
	indeks = 0
	for indeks < totalTugas {
		if strings.EqualFold(daftarTugasUtama[indeks].Kategori, kategoriDicari) {
			cetakSatuTugas(daftarTugasUtama[indeks], indeks)
			apakahDitemukan = true
		}
		indeks++
	}

	if !apakahDitemukan {
		fmt.Println("Data tugas dengan kategori tersebut tidak ditemukan.")
	}
}

func cariByDeadlineBinary(totalTugas int) {
	var tanggalDicari string
	var posisiKiri, posisiKanan, titikTengah, indeksHasilTemuan int
	var indeks int

	selectionSortTugas(true) 
	fmt.Print("Masukkan tanggal deadline yang dicari (YYYY-MM-DD): ")
	fmt.Scanf("%s\n", &tanggalDicari)

	posisiKiri = 0
	posisiKanan = totalTugas - 1
	indeksHasilTemuan = -1

	for posisiKiri <= posisiKanan && indeksHasilTemuan == -1 {
		titikTengah = (posisiKiri + posisiKanan) / 2
		if daftarTugasUtama[titikTengah].Deadline == tanggalDicari {
			indeksHasilTemuan = titikTengah
		} else if daftarTugasUtama[titikTengah].Deadline < tanggalDicari {
			posisiKiri = titikTengah + 1
		} else {
			posisiKanan = titikTengah - 1
		}
	}

	if indeksHasilTemuan != -1 {
		fmt.Println("\nData ditemukan:")
		indeks = 0
		for indeks < totalTugas {
			if daftarTugasUtama[indeks].Deadline == tanggalDicari {
				cetakSatuTugas(daftarTugasUtama[indeks], indeks)
			}
			indeks++
		}
	} else {
		fmt.Println("Data dengan deadline tersebut tidak ditemukan.")
	}
}

func cariByTanggalJamSequential(totalTugas int) {
	var tanggalDicari, jamDicari string
	var apakahDitemukan bool
	var indeks int

	fmt.Print("Masukkan Tanggal (YYYY-MM-DD): ")
	fmt.Scanf("%s\n", &tanggalDicari)
	fmt.Print("Masukkan Jam (HH:MM): ")
	fmt.Scanf("%s\n", &jamDicari)

	apakahDitemukan = false
	indeks = 0
	for indeks < totalTugas {
		if daftarTugasUtama[indeks].Deadline == tanggalDicari && daftarTugasUtama[indeks].Jam == jamDicari {
			cetakSatuTugas(daftarTugasUtama[indeks], indeks)
			apakahDitemukan = true
		}
		indeks++
	}

	if !apakahDitemukan {
		fmt.Println("Data dengan kombinasi tanggal & jam tersebut tidak ditemukan.")
	}
}

func ubahTugas(pembacaInput *bufio.Reader) {
	var totalTugas, nomorTugasPilihan, indeksPilihan int
	var mataKuliahBaru, kategoriBaru string
	var tanggalDeadlineBaru, jamPengumpulanBaru string

	totalTugas = hitungJumlahTugas()
	if totalTugas == 0 {
		fmt.Println("Tidak ada tugas yang bisa diubah.")
		return
	}

	tampilkanSemuaTugas()

	fmt.Print("Masukkan nomor tugas yang ingin diubah: ")
	fmt.Scanf("%d\n", &nomorTugasPilihan)
	indeksPilihan = nomorTugasPilihan - 1

	if indeksPilihan < 0 || indeksPilihan >= totalTugas {
		fmt.Println("Nomor tidak valid.")
		return
	}

	fmt.Print("Masukkan Mata Kuliah baru: ")
	mataKuliahBaru, _ = pembacaInput.ReadString('\n')
	mataKuliahBaru = strings.TrimSpace(mataKuliahBaru)

	fmt.Print("Masukkan Kategori baru: ")
	kategoriBaru, _ = pembacaInput.ReadString('\n')
	kategoriBaru = strings.TrimSpace(kategoriBaru)

	fmt.Print("Masukkan Deadline baru (YYYY-MM-DD): ")
	fmt.Scanf("%s\n", &tanggalDeadlineBaru)
	fmt.Print("Masukkan Jam Pengumpulan baru (HH:MM): ")
	fmt.Scanf("%s\n", &jamPengumpulanBaru)

	if validasiTanggal(tanggalDeadlineBaru) && validasiJam(jamPengumpulanBaru) {
		daftarTugasUtama[indeksPilihan].MataKuliah = mataKuliahBaru
		daftarTugasUtama[indeksPilihan].Kategori = kategoriBaru
		daftarTugasUtama[indeksPilihan].Deadline = tanggalDeadlineBaru
		daftarTugasUtama[indeksPilihan].Jam = jamPengumpulanBaru
		fmt.Println("Data tugas berhasil diperbarui!")
	} else {
		fmt.Println("Perubahan gagal. Pastikan tanggal berada di rentang Juni 2026.")
	}
}

func hapusTugas() {
	var totalTugas, nomorTugasPilihan, indeksPilihan, indeks int
	var teksKonfirmasi string

	totalTugas = hitungJumlahTugas()
	if totalTugas == 0 {
		fmt.Println("Tidak ada tugas untuk dihapus.")
		return
	}

	tampilkanSemuaTugas()

	fmt.Print("Masukkan nomor tugas yang ingin dihapus: ")
	fmt.Scanf("%d\n", &nomorTugasPilihan)
	indeksPilihan = nomorTugasPilihan - 1

	if indeksPilihan < 0 || indeksPilihan >= totalTugas {
		fmt.Println("Nomor tidak valid.")
		return
	}

	fmt.Printf("Apakah Anda yakin menghapus tugas %s? (y/n): ", daftarTugasUtama[indeksPilihan].MataKuliah)
	fmt.Scanf("%s\n", &teksKonfirmasi)

	if teksKonfirmasi == "y" || teksKonfirmasi == "Y" {
		indeks = indeksPilihan
		for indeks < totalTugas-1 {
			daftarTugasUtama[indeks] = daftarTugasUtama[indeks+1]
			indeks++
		}
		daftarTugasUtama[totalTugas-1] = Tugas{}
		fmt.Println("Tugas berhasil dihapus.")
	} else {
		fmt.Println("Penghapusan dibatalkan.")
	}
}

func urutkanTugas() {
	var totalTugas, pilihanMetode, pilihanKriteria int

	totalTugas = hitungJumlahTugas()
	if totalTugas == 0 {
		fmt.Println("Tidak ada data untuk diurutkan.")
		return
	}

	fmt.Println("\nPilih Metode Pengurutan:")
	fmt.Println("1. Selection Sort (Berdasarkan Deadline)")
	fmt.Println("2. Insertion Sort (Berdasarkan Status)")
	fmt.Print("Pilihan: ")
	fmt.Scanf("%d\n", &pilihanMetode)

	if pilihanMetode == 1 {
		fmt.Println("Urutkan Deadline secara:")
		fmt.Println("1. Terdekat (Ascending)")
		fmt.Println("2. Terlama (Descending)")
		fmt.Print("Pilihan: ")
		fmt.Scanf("%d\n", &pilihanKriteria)
		selectionSortTugas(pilihanKriteria == 1)
		fmt.Println("Data berhasil diurutkan berdasarkan Deadline Juni 2026!")
	} else if pilihanMetode == 2 {
		fmt.Println("Urutkan Status secara:")
		fmt.Println("1. Belum Selesai -> Selesai (Ascending)")
		fmt.Println("2. Selesai -> Belum Selesai (Descending)")
		fmt.Print("Pilihan: ")
		fmt.Scanf("%d\n", &pilihanKriteria)
		insertionSortTugas(pilihanKriteria == 1)
		fmt.Println("Data berhasil diurutkan berdasarkan Status!")
	} else {
		fmt.Println("Metode tidak tersedia.")
		return
	}

	tampilkanSemuaTugas()
}

func selectionSortTugas(apakahAscending bool) {
	var totalTugas, indeksLuar, indeksDalam, indeksTerpilih int

	totalTugas = hitungJumlahTugas()
	indeksLuar = 0
	for indeksLuar < totalTugas-1 {
		indeksTerpilih = indeksLuar
		indeksDalam = indeksLuar + 1
		for indeksDalam < totalTugas {
			if apakahAscending {
				if daftarTugasUtama[indeksDalam].Deadline < daftarTugasUtama[indeksTerpilih].Deadline {
					indeksTerpilih = indeksDalam
				}
			} else {
				if daftarTugasUtama[indeksDalam].Deadline > daftarTugasUtama[indeksTerpilih].Deadline {
					indeksTerpilih = indeksDalam
				}
			}
			indeksDalam++
		}
		daftarTugasUtama[indeksLuar], daftarTugasUtama[indeksTerpilih] = daftarTugasUtama[indeksTerpilih], daftarTugasUtama[indeksLuar]
		indeksLuar++
	}
}

func insertionSortTugas(apakahAscending bool) {
	var totalTugas, indeksLuar, indeksDalam int
	var dataPenyimpanSementara Tugas
	var nilaiBobotIndeksDalam, nilaiBobotSementara int
	var apakahMemenuhiKondisiTukar, loopTerus bool

	totalTugas = hitungJumlahTugas()
	indeksLuar = 1
	for indeksLuar < totalTugas {
		dataPenyimpanSementara = daftarTugasUtama[indeksLuar]
		indeksDalam = indeksLuar - 1

		loopTerus = indeksDalam >= 0
		for loopTerus {
			if daftarTugasUtama[indeksDalam].Status {
				nilaiBobotIndeksDalam = 1
			} else {
				nilaiBobotIndeksDalam = 0
			}
			if dataPenyimpanSementara.Status {
				nilaiBobotSementara = 1
			} else {
				nilaiBobotSementara = 0
			}

			apakahMemenuhiKondisiTukar = false
			if apakahAscending {
				apakahMemenuhiKondisiTukar = nilaiBobotIndeksDalam > nilaiBobotSementara
			} else {
				apakahMemenuhiKondisiTukar = nilaiBobotIndeksDalam < nilaiBobotSementara
			}

			if apakahMemenuhiKondisiTukar {
				daftarTugasUtama[indeksDalam+1] = daftarTugasUtama[indeksDalam]
				indeksDalam--
				loopTerus = indeksDalam >= 0
			} else {
				loopTerus = false
			}
		}
		daftarTugasUtama[indeksDalam+1] = dataPenyimpanSementara
		indeksLuar++
	}
}

func ubahStatusTugas() {
	var totalTugas, nomorTugasPilihan, indeksPilihan, pilihanStatus int

	totalTugas = hitungJumlahTugas()
	if totalTugas == 0 {
		fmt.Println("Belum ada data tugas.")
		return
	}

	tampilkanSemuaTugas()

	fmt.Print("Masukkan nomor tugas yang ingin diubah statusnya: ")
	fmt.Scanf("%d\n", &nomorTugasPilihan)
	indeksPilihan = nomorTugasPilihan - 1

	if indeksPilihan < 0 || indeksPilihan >= totalTugas {
		fmt.Println("Nomor tidak valid.")
		return
	}

	fmt.Println("Ubah status menjadi:")
	fmt.Println("1. Selesai")
	fmt.Println("2. Belum Selesai")
	fmt.Print("Pilihan: ")
	fmt.Scanf("%d\n", &pilihanStatus)

	if pilihanStatus == 1 {
		daftarTugasUtama[indeksPilihan].Status = true
		fmt.Println("Mantap! Tugas telah ditandai SELESAI.")
	} else if pilihanStatus == 2 {
		daftarTugasUtama[indeksPilihan].Status = false
		fmt.Println("Status diatur kembali ke BELUM SELESAI.")
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}
