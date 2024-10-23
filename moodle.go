package main

import "fmt"

const NMAX int = 80

type Siswa struct {
	nama                string
	CourseDiambil       [NMAX]DataCourse
	jumlahCourseDiambil int
}

type Guru struct {
	nama         string
	password     string
	CourseGuru   [NMAX]DataCourse
	jumlahCourse int
}

type DataCourse struct {
	namaCourse string
	Tugas      string
	Quiz       string
	Forum      string
	Join       int
	NilaiTugas int
	NilaiQuiz  int
	Dijawab    jawaban
}

type jawaban struct {
	JawabTugas string
	JawabQuiz  string
	JawabForum string
}

type Arrguru struct {
	Tabguru [NMAX]Guru
	Jumguru int
}

type Arrsiswa struct {
	Tabsiswa [NMAX]Siswa
	Jumsiswa int
}

func main() {
	var dataguru Arrguru
	var datasiswa Arrsiswa
	var pilihanUser int
	menuUtama(pilihanUser, &dataguru, &datasiswa)
	fmt.Println("Terimakasih telah memakai program ini :D")
}

func menuUtama(pilihan int, T *Arrguru, T1 *Arrsiswa) {
	//func ini berfungsi sebagai hub utama dari moodleku
	//fungsi diantaranya adalah untuk memilih role dan keluar
	//guru juga bisa menilai dan melihat jumlah dari siswa yang mengikuti suatu course
	fmt.Println("=======================================================================")
	fmt.Print("Selamat datang di program Moodleku :) \n ")
	fmt.Println("Silahkan Pilih Role Sesuai Kebutuhan: 1 Siswa, 2 Guru, 3 Keluar")
	fmt.Println("=======================================================================")
	fmt.Print("Pilihan :")
	fmt.Scan(&pilihan)
	for pilihan != 3 {
		if pilihan == 1 {
			utamaSiswa(pilihan, T1, T)
			fmt.Print("Selamat datang di program Moodleku :) \n")
			fmt.Println("Silahkan Pilih Role Sesuai Kebutuhan: 1 Siswa, 2 Guru, 3 Keluar")
			fmt.Print("Pilihan :")
			fmt.Scan(&pilihan)
		} else if pilihan == 2 {
			utamaguru(pilihan, T, T1)
			fmt.Print("Selamat datang di program Moodleku :) \n")
			fmt.Println("Silahkan Pilih Role Sesuai Kebutuhan: 1 Siswa, 2 Guru, 3 Keluar")
			fmt.Print("Pilihan :")
			fmt.Scan(&pilihan)
		} else {
			fmt.Print("Pilihan:")
			fmt.Scan(&pilihan)
		}
	}
}

func utamaSiswa(pilihan int, T1 *Arrsiswa, T *Arrguru) {
	//Menu utama dari role siswa, siswa bisa melakukan beberapa fungsi diantaranya: mengikuti course yang telah tersedia, dan mengerjakan konten didalam course tersebut
	fmt.Println("=======================================================================")
	fmt.Print("Selamat datang Siswa :D \n")
	fmt.Print("Silahkan masukan nama kakak/adik untuk register \n")
	fmt.Println("=======================================================================")
	fmt.Print("Nama:")
	fmt.Scan(&T1.Tabsiswa[T1.Jumsiswa].nama)
	*&T1.Jumsiswa++
	fmt.Println("Hi kakak/adik", T1.Tabsiswa[T1.Jumsiswa-1].nama, "Silahkan pilih sesuai kebutuhan: 1 Pilih Course, 2 Kerjakan Konten, 3 Kembali Ke Menu Utama")
	fmt.Print("Pilihan:")
	fmt.Scan(&pilihan)
	for pilihan != 3 {
		if pilihan == 1 {
			pilihCourse(T, T1)
			fmt.Println("Silahkan pilih sesuai kebutuhan: 1 Pilih Course, 2 Kerjakan Konten, 3 Kembali Ke Menu Utama")
			fmt.Print("\n Pilihan:")
			fmt.Scan(&pilihan)
		} else if pilihan == 2 {
			kerjakanKonten(T1)
			fmt.Println("Silahkan pilih sesuai kebutuhan:1 Pilih Course, 2 Kerjakan Konten, 3 Kembali Ke Menu Utama")
			fmt.Print("\n Pilihan:")
			fmt.Scan(&pilihan)
		} else {
			fmt.Print("Pilihan:")
			fmt.Scan(&pilihan)
		}
	}

}

func utamaguru(pilihan int, T *Arrguru, T1 *Arrsiswa) {
	//menu utama dari guru, guru bisa melakukan beberapa fungsi diantaranya: menambah course dengan 1 course terdiri dari 1 tugas,1 quiz,1 Forum, edit course, dan menghapus course
	var indeksguru, menu int
	var namaGuru string
	var password string
	fmt.Println("=======================================================================")
	fmt.Print("Selamat datang bapak/ibu guru :D \n")
	fmt.Print("Sudah pernah register? 1 Sudah, 2 Belum \n")
	fmt.Print("Pilihan:")
	fmt.Scan(&pilihan)
	fmt.Println("=======================================================================")
	if pilihan == 1 {
		fmt.Println("=======================================================================")
		fmt.Print("Selamat datang kembali bapak/ibu guru :D \n")
		fmt.Print("Silahkan masukan nama dan passwordnya kembali  \n")
		fmt.Print("\n Nama:")
		fmt.Scan(&namaGuru)
		fmt.Print("\n Password:")
		fmt.Scan(&password)
		fmt.Println("=======================================================================")
		for pilihan != 3 && menu != 2 {
			indeksguru = cariNamaguruBinary(*T, T.Jumguru, namaGuru)
			if indeksguru != -1 && password == T.Tabguru[indeksguru].password {
				fmt.Println("Hi bapak/ibu", T.Tabguru[indeksguru].nama, "Silahkan pilih sesuai kebutuhan: 1 Nilai Jawaban Siswa, 2 Lihat Data Course, 3 Kembali Ke Menu Utama")
				fmt.Print("Pilihan:")
				fmt.Scan(&pilihan)
				if pilihan == 1 {
					MenilaiSiswa(T, T1, indeksguru)
					fmt.Println("Silahkan pilih sesuai kebutuhan: 1 Nilai Jawaban Siswa, 2 Lihat Data Course, 3 Kembali Ke Menu Utama")
					fmt.Print("Pilihan:")
					fmt.Scan(&pilihan)
				} else if pilihan == 2 {
					LiatDataCourse(*T, *T1, indeksguru)
					fmt.Println("Silahkan pilih sesuai kebutuhan: 1 Nilai Jawaban Siswa, 2 Lihat Data Course, 3 Kembali Ke Menu Utama")
					fmt.Print("Pilihan:")
					fmt.Scan(&pilihan)
				} else {
					fmt.Print("Pilihan:")
					fmt.Scan(&pilihan)
				}
			} else {
				fmt.Print("\n Maaf bapak/ibu tidak terdaftar atau password dan namanya salah")
				fmt.Print("\n Coba login lagi? 1 Iya 2 Tidak (Kembali Ke Menu): ")
				fmt.Scan(&menu)
				if menu == 1 {
					fmt.Print("Silahkan masukan nama dan passwordnya kembali  \n")
					fmt.Print("\n Nama:")
					fmt.Scan(&namaGuru)
					fmt.Print("\n Password:")
					fmt.Scan(&password)
				}

			}
		}
	} else if pilihan == 2 {
		fmt.Print("Silahkan masukan nama bapak/ibu untuk register \n")
		fmt.Print("\n Nama:")
		fmt.Scan(&T.Tabguru[T.Jumguru].nama)
		fmt.Print("\n Password:")
		fmt.Scan(&T.Tabguru[T.Jumguru].password)
		*&T.Jumguru++
		fmt.Println("Hi bapak/ibu", T.Tabguru[T.Jumguru-1].nama, "Silahkan pilih sesuai kebutuhan: 1 Tambah Data Course, 2 Edit Data Course, 3 Hapus Data Course, 4 Kembali Ke Menu Utama")
		fmt.Print("Pilihan:")
		fmt.Scan(&pilihan)
		for pilihan != 4 {
			if pilihan == 1 {
				tambahCourse(T, &T.Tabguru[T.Jumguru-1].jumlahCourse)
				fmt.Println("Silahkan pilih sesuai kebutuhan: 1 Tambah Data Course, 2 Edit Data Course, 3 Hapus Data Course, 4 Kembali Ke Menu Utama")
				fmt.Print("Pilihan:")
				fmt.Scan(&pilihan)
			} else if pilihan == 2 {
				editCourse(T)
				fmt.Println("Silahkan pilih sesuai kebutuhan: 1 Tambah Data Course, 2 Edit Data Course, 3 Hapus Data Course, 4 Kembali Ke Menu Utama")
				fmt.Print("Pilihan:")
				fmt.Scan(&pilihan)
			} else if pilihan == 3 {
				HapusCourse(T)
				fmt.Println("Silahkan pilih sesuai kebutuhan: 1 Tambah Data Course, 2 Edit Data Course, 3 Hapus Data Course, 4 Kembali Ke Menu Utama")
				fmt.Print("Pilihan:")
				fmt.Scan(&pilihan)
			} else {
				fmt.Print("Pilihan:")
				fmt.Scan(&pilihan)
			}
		}
	}

}

func pilihCourse(T *Arrguru, T1 *Arrsiswa) {
	//Siswa memilih course yang tersedia sesuai keinginannya
	var pilihan string
	var traversal, menu int
	var temp DataCourse
	var jumC, JumS int

	if T.Jumguru == 0 {
		jumC = 0
	} else {
		jumC = T.Jumguru - 1
	}

	if T1.Jumsiswa == 0 {
		JumS = 0
	} else {
		JumS = T1.Jumsiswa - 1
	}

	if T.Tabguru[jumC].jumlahCourse > 0 {
		fmt.Print("\n Berikut adalah course yang ada guru yang mengadakan course tersebut:")
		SortingCourse_Selection(T, T.Tabguru[jumC].jumlahCourse, jumC)
		PrintNamaCourse(*T)
		fmt.Print("\n Masukan nama course yang mau anda ikuti:")
		fmt.Scan(&pilihan)
		for menu != 2 {
			if indekscourseSequencial(*T, pilihan) != -1 {
				temp = T.Tabguru[jumC].CourseGuru[indekscourseSequencial(*T, pilihan)]
				fmt.Print("\n Join course ", T.Tabguru[jumC].CourseGuru[indekscourseSequencial(*T, pilihan)].namaCourse, " 1 Iya 2 Tidak:")
				fmt.Scan(&traversal)
				if traversal == 1 {
					T1.Tabsiswa[JumS].CourseDiambil[T1.Tabsiswa[JumS].jumlahCourseDiambil] = temp
					fmt.Print("\n Course berhasil ditambahkan!")
					T.Tabguru[jumC].CourseGuru[T.Tabguru[jumC].jumlahCourse-1].Join++
					T1.Tabsiswa[JumS].jumlahCourseDiambil++
				} else if traversal == 2 {
					fmt.Print("\n Course tidak jadi ditambahkan")
				} else {
					fmt.Print("\n Join course ini? 1 Iya 2 Tidak:")
					fmt.Scan(&traversal)
				}
			} else {
				fmt.Print("\n Maaf data tidak ditemukan")
			}

			fmt.Print("\n Pilih course lagi? 1 Iya 2 Tidak:")
			fmt.Scan(&menu)
			if menu == 1 {
				fmt.Print("\n Berikut adalah course yang ada guru yang mengadakan course tersebut:")
				PrintNamaCourse(*T)
				fmt.Print("\n Masukan nama course yang mau anda ikuti:")
				fmt.Scan(&pilihan)
			}
		}

	} else {
		fmt.Print("\n Mohon maaf saat ini tidak ada course yang tersedia.")
	}

}

func kerjakanKonten(T1 *Arrsiswa) {
	//Siswa mengerjakan konten yang telah diambil sebelumnya
	var menu, menu2 int
	var jumS int
	var idx int
	var pilihan string

	if T1.Jumsiswa == 0 {
		jumS = 0
	} else {
		jumS = T1.Jumsiswa - 1
	}

	if T1.Tabsiswa[jumS].jumlahCourseDiambil > 0 {
		PrintCourseDiambil(*T1)
		fmt.Print("\n Pilih course yang mau anda kerjakan:")
		fmt.Scan(&pilihan)
		idx = indeksCourseBinary(*T1, T1.Tabsiswa[jumS].jumlahCourseDiambil, jumS, pilihan)
		for menu != 4 && menu2 != 2 {
			if idx != -1 {
				fmt.Print("\n Silahkan pilih tipe konten yang mau dikerjakan di course ", T1.Tabsiswa[jumS].CourseDiambil[idx].namaCourse, ": 1 Tugas, 2 Quiz, 3 Forum, 4 Kembali Ke Menu Siswa")
				fmt.Print("\n Pilihan:")
				fmt.Scan(&menu)
				tipekonten(T1, menu, idx, jumS)
			} else {
				fmt.Print("\n Course tidak ditemukan, silahkan pilih course lagi")
			}

			fmt.Print("\n Mau kerjakan course lain? 1 Iya 2 Tidak:")
			fmt.Scan(&menu2)
			if menu2 == 1 {
				PrintCourseDiambil(*T1)
				fmt.Print("\n Pilih course yang mau anda kerjakan:")
				fmt.Scan(&pilihan)
				idx = indeksCourseBinary(*T1, T1.Tabsiswa[jumS].jumlahCourseDiambil, jumS, pilihan)
			}
		}
	} else {
		fmt.Print("\n Maaf anda belum ada course yang bisa dikerjakan, silahkan ambil/pilih course terlebih dahulu ")
	}
}

func tipekonten(T1 *Arrsiswa, menu, idx, JumS int) {
	var temp jawaban
	var simpan int
	if menu == 1 {
		fmt.Print("\n Berikut adalah soal tugas dari course ini:")
		fmt.Print("\n", T1.Tabsiswa[JumS].CourseDiambil[idx].Tugas)
		fmt.Print("\n Jawaban anda:")
		fmt.Scan(&temp.JawabTugas)
		fmt.Print("\n Simpan jawaban? 1 Iya 2 Tidak")
		fmt.Scan(&simpan)
		if simpan == 1 {
			T1.Tabsiswa[JumS].CourseDiambil[idx].Dijawab.JawabTugas = temp.JawabTugas
			fmt.Print("\n Jawaban tersimpan")
		} else if simpan == 2 {
			fmt.Print("\n Jawaban tidak tersimpan")
		} else {
			fmt.Print("\n Simpan jawaban? 1 Iya 2 Tidak")
			fmt.Scan(&simpan)
		}
	} else if menu == 2 {
		fmt.Print("\n Berikut adalah soal quiz dari course ini:")
		fmt.Print("\n", T1.Tabsiswa[JumS].CourseDiambil[idx].Quiz)
		fmt.Print("\n Jawaban anda:")
		fmt.Scan(&temp.JawabQuiz)
		fmt.Print("\n Simpan jawaban? 1 Iya 2 Tidak")
		fmt.Scan(&simpan)
		if simpan == 1 {
			T1.Tabsiswa[JumS].CourseDiambil[idx].Dijawab.JawabQuiz = temp.JawabQuiz
			fmt.Print("\n Jawaban tersimpan")
		} else if simpan == 2 {
			fmt.Print("\n Jawaban tidak tersimpan")
		} else {
			fmt.Print("\n Simpan jawaban? 1 Iya 2 Tidak")
			fmt.Scan(&simpan)
		}
	} else if menu == 3 {
		fmt.Print("\n Berikut adalah diskusi forum dari course ini:")
		fmt.Print("\n", T1.Tabsiswa[JumS].CourseDiambil[idx].Forum)
		fmt.Print("\n Jawaban anda:")
		fmt.Scan(&temp.JawabForum)
		fmt.Print("\n Simpan jawaban? 1 Iya 2 Tidak")
		fmt.Scan(&simpan)
		if simpan == 1 {
			T1.Tabsiswa[JumS].CourseDiambil[idx].Dijawab.JawabForum = temp.JawabForum
			fmt.Print("\n Jawaban tersimpan")
		} else if simpan == 2 {
			fmt.Print("\n Jawaban tidak tersimpan")
		} else {
			fmt.Print("\n Simpan jawaban? 1 Iya 2 Tidak")
			fmt.Scan(&simpan)
		}
	} else {
		fmt.Print("\n Pilihan:")
		fmt.Scan(&menu)
	}
}

func LiatDataCourse(T Arrguru, T1 Arrsiswa, indexguru int) {
	//Melihat data seluruh course yang dimiliki oleh guru
	var i, e int
	fmt.Print("\n Berikut adalah data course")
	sortingJumlahIkutCourse(&T, T.Tabguru[indexguru].jumlahCourse, indexguru)
	SortingNilaiQuiz_Insertion(&T1, T1.Tabsiswa[T1.Jumsiswa-1].jumlahCourseDiambil, T1.Jumsiswa-1)
	SortingNilaiTugas_Insertion(&T1, T1.Tabsiswa[T1.Jumsiswa-1].jumlahCourseDiambil, T1.Jumsiswa-1)
	e = 0
	for i < T.Tabguru[indexguru].jumlahCourse && e < T1.Jumsiswa {
		fmt.Println(T1.Tabsiswa[e].CourseDiambil[i].namaCourse, T1.Tabsiswa[e].CourseDiambil[i].NilaiQuiz, T1.Tabsiswa[e].CourseDiambil[i].NilaiTugas, T.Tabguru[indexguru].CourseGuru[i].Join)
		i++
		if i == T.Tabguru[indexguru].jumlahCourse {
			e++
			i--
		}
	}

}

func MenilaiSiswa(T *Arrguru, T1 *Arrsiswa, idx int) {
	var namaSiswa, namacors string
	var menu, pilih int
	var indexsiswa, nilaitemp, indexCourse int
	printSiswa(*T1)
	fmt.Print("\n Masukan nama siswa yang ingin dinilai course nya:")
	fmt.Scan(&namaSiswa)
	PrintCourseDiambil(*T1)
	fmt.Print("\n Masukan nama course yang ingin diseleksi:")
	fmt.Scan(&namacors)
	SortingNamaSiswa_Insertion(T1, T1.Jumsiswa)
	indexsiswa = carinamaSiswaBinary(*T1, T1.Jumsiswa, namaSiswa)
	indexCourse = indekscourseSequencialSiswa(*T1, namacors)
	for menu != 3 && pilih != 2 {
		if indexsiswa != -1 && indexCourse != -1 {
			fmt.Print("\n Pilihlah tipe konten yang mau dikoreksi 1 Tugas, 2 Quiz, 3 Kembali Ke Menu Utama: ")
			fmt.Scan(&menu)
			if menu == 1 {
				fmt.Print("\n Berikut adalah jawaban siswa untuk tugas ini:")
				fmt.Print("\n", T1.Tabsiswa[indexsiswa].CourseDiambil[indexCourse].Tugas)
				fmt.Print("\n Masukan nilai untuk jawaban ini:")
				fmt.Scan(&nilaitemp)
				fmt.Print("\n Simpan Nilai ini? 1 Iya 2 Tidak: ")
				fmt.Scan(&pilih)
				if pilih == 1 {
					T1.Tabsiswa[indexsiswa].CourseDiambil[indexCourse].NilaiTugas = nilaitemp
					fmt.Print("\n Nilai siswa tersimpan!")
				}
			} else if menu == 2 {
				fmt.Print("\n Berikut adalah jawaban siswa untuk tugas ini:")
				fmt.Print("\n", T1.Tabsiswa[indexsiswa].CourseDiambil[indexCourse].Quiz)
				fmt.Print("\n Masukan nilai untuk jawaban ini:")
				fmt.Scan(&nilaitemp)
				fmt.Print("\n Simpan Nilai ini? 1 Iya 2 Tidak: ")
				fmt.Scan(&pilih)
				if pilih == 1 {
					T1.Tabsiswa[indexsiswa].CourseDiambil[indexCourse].NilaiQuiz = nilaitemp
					fmt.Print("\n Nilai siswa tersimpan!")
				}
			}
			fmt.Print("\n Nilai course lagi? 1 Iya 2 Tidak")
			fmt.Scan(&pilih)
			if pilih == 1 {
				printSiswa(*T1)
				fmt.Print("\n Masukan nama siswa yang ingin dinilai course nya:")
				fmt.Scan(&namaSiswa)
				PrintCourseDiambil(*T1)
				fmt.Print("\n Masukan nama course yang ingin diseleksi:")
				fmt.Scan(&namacors)
				indexsiswa = carinamaSiswaBinary(*T1, T1.Jumsiswa, namaSiswa)
				indexCourse = indekscourseSequencialSiswa(*T1, namacors)
			}
		} else {
			fmt.Print("\n Siswa dengan nama tersebut tidak ada atau Siswa tersebut tidak mengambil course tersebut.")
			fmt.Print("\n Cari siswa lain? 1 Iya 2 Tidak")
			fmt.Scan(&pilih)
			if pilih == 1 {
				printSiswa(*T1)
				fmt.Print("\n Masukan nama siswa yang ingin dinilai course nya:")
				fmt.Scan(&namaSiswa)
				fmt.Print("\n Masukan nama course yang ingin diseleksi:")
				fmt.Scan(&namacors)
				indexsiswa = carinamaSiswaBinary(*T1, T1.Jumsiswa, namaSiswa)
				indexCourse = indekscourseSequencialSiswa(*T1, namacors)
			}
		}

	}

}

func tambahCourse(T *Arrguru, jumC *int) {
	var pilihanKonten int
	var tambahLagi, simpanSoal int
	var temp DataCourse
	//jumC adalah acuan dari jumlah course yang dimiliki oleh seorang guru dengan nama T.Tabguru.nama
	fmt.Print("\n Masukan nama course:")
	fmt.Scan(&T.Tabguru[*&T.Jumguru-1].CourseGuru[*jumC].namaCourse)
	fmt.Print("\n Untuk course,", T.Tabguru[*&T.Jumguru-1].CourseGuru[*jumC].namaCourse, "\n Pilih tipe konten yang diinginkan: 1 Tugas, 2 Quiz, 3 Forum, 4 Selesai:")
	fmt.Scan(&pilihanKonten)
	*jumC++

	for tambahLagi != 2 && pilihanKonten != 4 {
		if pilihanKonten == 1 {
			fmt.Print("\n Masukan soal untuk tugas yang diinginkan (ganti spasi dengan tanda _):")
			fmt.Scan(&temp.Tugas)
			fmt.Print("\n Apakah anda ingin menyimpan soal tugas tersebut? 1 iya 2 tidak:")
			fmt.Scan(&simpanSoal)
			if simpanSoal == 1 {
				T.Tabguru[*&T.Jumguru-1].CourseGuru[*jumC-1].Tugas = temp.Tugas
				fmt.Print("\n Data tersimpan :D")
			}
			fmt.Print("\n Apakah anda ingin menambahkan data lagi? 1 iya 2 tidak:")
			fmt.Scan(&tambahLagi)

		} else if pilihanKonten == 2 {
			fmt.Print("\n Masukan soal untuk quiz yang diinginkan (ganti spasi dengan tanda _):")
			fmt.Scan(&temp.Quiz)
			fmt.Print("\n Apakah anda ingin menyimpan soal quiz tersebut? 1 iya 2 tidak:")
			fmt.Scan(&simpanSoal)
			if simpanSoal == 1 {
				T.Tabguru[*&T.Jumguru-1].CourseGuru[*jumC-1].Quiz = temp.Quiz
				fmt.Print("\n Data tersimpan :D")
			}
			fmt.Print("\n Apakah anda ingin menambahkan data lagi? 1 iya 2 tidak:")
			fmt.Scan(&tambahLagi)

		} else if pilihanKonten == 3 {
			fmt.Print("\n Masukan soal forum diskusi yang diinginkan (ganti spasi dengan tanda _):")
			fmt.Scan(&temp.Forum)
			fmt.Print("\n Apakah anda ingin menyimpan diskusi forum tersebut? 1 iya 2 tidak:")
			fmt.Scan(&simpanSoal)
			if simpanSoal == 1 {
				T.Tabguru[*&T.Jumguru-1].CourseGuru[*jumC-1].Forum = temp.Forum
				fmt.Print("\n Data tersimpan :D")
			}
			fmt.Print("\n Apakah anda ingin menambahkan data lagi? 1 iya 2 tidak:")
			fmt.Scan(&tambahLagi)
		} else {
			fmt.Print("\n Untuk course,", T.Tabguru[*&T.Jumguru-1].CourseGuru[*jumC].namaCourse, "\n Pilih tipe konten yang diinginkan: 1 Tugas, 2 Quiz, 3 Forum, 4 Selesai:")
			fmt.Scan(&pilihanKonten)
		}

		if tambahLagi == 1 {
			fmt.Print("\n Untuk course,", T.Tabguru[*&T.Jumguru-1].CourseGuru[*jumC-1].namaCourse, "\n Pilih tipe konten yang diinginkan: 1 Tugas, 2 Quiz, 3 Forum, 4 Selesai:")
			fmt.Scan(&pilihanKonten)
		}
		//Untuk saat ini 1 course hanya bisa 1 soal tugas, 1 soal quiz, sama 1 pertanyaan forum
		//Nanti tambahin kalau udah terisi satu tugas kalau mau ngisi yang lain atau edit di bagian editCourse
	}

}

func editCourse(T *Arrguru) {
	var namacourse string
	var pilihan, menu int
	var temp DataCourse
	var jumC int

	//Ngedit salah satu aspek dari Course setelah dipilih
	//Yang bisa diedit hanya yang ada di guru tersebut

	if T.Jumguru == 0 {
		jumC = 0
	} else {
		jumC = T.Jumguru - 1
	}
	//Kondisi untuk menunjuk jumlah course di suatu guru
	if T.Tabguru[jumC].jumlahCourse != 0 {
		SortingCourse_Insertion(T, T.Tabguru[jumC].jumlahCourse, jumC)
		PrintIsiCourse(*T)
		fmt.Print("\n Silahkan pilih nama course yang mau diedit:")
		fmt.Scan(&namacourse)
		if cariCourseSequencial(*T, namacourse) == true {
			fmt.Print("\n Untuk Course, ", T.Tabguru[jumC].CourseGuru[indekscourseSequencial(*T, namacourse)].namaCourse, " pilih jenis data yang ingin diedit: 1 Soal Tugas, 2 Soal Quiz, 3 Forum Diskusi, 4 Kembali ke menu guru:")
			fmt.Scan(&pilihan)
			for menu != 4 {
				if pilihan == 1 {
					fmt.Print("\n Masukan soal Tugas yang baru:")
					fmt.Scan(&temp.Tugas)
					fmt.Print("\n Apakah soal ingin disimpan? 1 Iya 2 Tidak:")
					fmt.Scan(&pilihan)
					if pilihan == 1 {
						T.Tabguru[jumC].CourseGuru[indekscourseSequencial(*T, namacourse)].Tugas = temp.Tugas
						//T.Tabguru[T.Jumguru-1] udah bener tapi .CourseGuru[T.Jumguru-1] indeksnya harus diganti dengan indeks dari Course yang berhasil dicari dengan sequencial search
						fmt.Print("\n Data Telah Tersimpan! :D")
					}
				} else if pilihan == 2 {
					fmt.Print("\n Masukan soal Quiz yang baru:")
					fmt.Scan(&temp.Quiz)
					fmt.Print("\n Apakah soal ingin disimpan? 1 Iya 2 Tidak:")
					fmt.Scan(&pilihan)
					if pilihan == 1 {
						T.Tabguru[jumC].CourseGuru[indekscourseSequencial(*T, namacourse)].Quiz = temp.Quiz
						fmt.Print("\n Data Telah Tersimpan! :D")
					}
				} else if pilihan == 3 {
					fmt.Print("\n Masukan diskusi Forum yang baru:")
					fmt.Scan(&temp.Forum)
					fmt.Print("\n Apakah soal ingin disimpan? 1 Iya 2 Tidak:")
					fmt.Scan(&pilihan)
					if pilihan == 1 {
						T.Tabguru[jumC].CourseGuru[indekscourseSequencial(*T, namacourse)].Forum = temp.Forum
						fmt.Print("\n Data Telah Tersimpan! :D")
					}
				} else {
					fmt.Print("\n Untuk Course, ", T.Tabguru[jumC].CourseGuru[indekscourseSequencial(*T, namacourse)].namaCourse, " Pilih jenis data yang ingin diedit: 1 Soal Tugas, 2 Soal Quiz, 3 Forum Diskusi:")
					fmt.Scan(&pilihan)
				}
				fmt.Print("\n Data terbaru:")
				PrintIsiCourse(*T)
				fmt.Print("\n Untuk Course, ", T.Tabguru[jumC].CourseGuru[indekscourseSequencial(*T, namacourse)].namaCourse, "Pilih jenis data yang ingin diedit: 1 Soal Tugas, 2 Soal Quiz, 3 Forum Diskusi, 4 Kembali ke menu guru:")
				fmt.Scan(&menu)
			}
		} else {
			fmt.Print("\n Course tidak ditemukan")
			fmt.Print("\n Silahkan pilih course yang mau diedit:")
			fmt.Scan(&namacourse)
		}
	} else {
		fmt.Println("\n Maaf belum ada data yang bisa diolah, silahkan tambahkan data terlebih dahulu :D")
	}

}

func HapusCourse(T *Arrguru) {
	//Procedure yang berguna untuk menghapus sebuah course secara keseluruhan (Dari soal tugas,quiz, maupun forum diskusi)
	var temp DataCourse
	var jumC int
	var traversal int
	var pilihan string
	var hasil_cari_indeks int
	temp.namaCourse = " "
	temp.Tugas = " "
	temp.Forum = " "
	temp.Quiz = " "
	if T.Jumguru == 0 {
		jumC = 0
	} else {
		jumC = T.Jumguru - 1
	}
	if T.Tabguru[jumC].jumlahCourse != 0 {
		SortingCourse_Insertion(T, T.Tabguru[jumC].jumlahCourse, jumC)
		PrintIsiCourse(*T)
		fmt.Print("\nSilahkan masukan nama course yang mau dihapus:")
		fmt.Scan(&pilihan)
		for traversal != 2 {
			if cariCourseSequencial(*T, pilihan) == true {
				fmt.Print("\n Anda akan menghapus Course:", T.Tabguru[jumC].CourseGuru[indekscourseSequencial(*T, pilihan)].namaCourse)
				fmt.Print("\n 1 Iya 2 Tidak :")
				fmt.Scan(&traversal)
				if traversal == 1 {
					hasil_cari_indeks = indekscourseSequencial(*T, pilihan)
					T.Tabguru[jumC].CourseGuru[hasil_cari_indeks].namaCourse = temp.namaCourse
					T.Tabguru[jumC].CourseGuru[hasil_cari_indeks].Tugas = temp.namaCourse
					T.Tabguru[jumC].CourseGuru[hasil_cari_indeks].Quiz = temp.Quiz
					T.Tabguru[jumC].CourseGuru[hasil_cari_indeks].Forum = temp.Forum
					if T.Tabguru[jumC].jumlahCourse-1 != 0 {
						for i := hasil_cari_indeks; i < T.Tabguru[jumC].jumlahCourse; i++ {
							T.Tabguru[jumC].CourseGuru[hasil_cari_indeks] = T.Tabguru[jumC].CourseGuru[(hasil_cari_indeks)+1]
							hasil_cari_indeks++
						}
					}
					T.Tabguru[jumC].jumlahCourse--
					fmt.Print("\n Data berhasil dihapus :D")
				} else {
					fmt.Print("\n Anda akan menghapus Course:", T.Tabguru[jumC].CourseGuru[indekscourseSequencial(*T, pilihan)].namaCourse)
					fmt.Print("\n 1 Iya 2 Tidak :")
					fmt.Scan(&traversal)
				}
				fmt.Print("\n Data terbaru:")
				PrintIsiCourse(*T)
				fmt.Print("Apakah anda ingin menghapus course lainnya? 1 Iya 2 Tidak (kembali ke menu guru):")
				fmt.Scan(&traversal)
			} else {
				fmt.Print("\nCourse tidak ditemukan")
				fmt.Print("\nCari lagi course yang mau dihapus? 1 Iya 2 Tidak:")
				fmt.Scan(&traversal)
			}
			if traversal == 1 {
				fmt.Print("\nSilahkan masukan nama course yang mau dihapus:")
				fmt.Scan(&pilihan)
			}
		}

	} else {
		fmt.Println("\n Maaf belum ada data yang bisa diolah, silahkan tambahkan data terlebih dahulu :D")
	}
}

func SortingCourse_Insertion(T *Arrguru, n, jumC int) {
	//Mengurutkan course guru dari A-Z secara insersi
	var i, j int
	var temp DataCourse
	i = 1
	for i <= n-1 {
		j = i
		temp = T.Tabguru[jumC].CourseGuru[j]
		for j > 0 && temp.namaCourse < T.Tabguru[jumC].CourseGuru[j-1].namaCourse {
			T.Tabguru[jumC].CourseGuru[j] = T.Tabguru[jumC].CourseGuru[j-1]
			j--
		}
		T.Tabguru[jumC].CourseGuru[j] = temp
		i++
	}
}

func sortingJumlahIkutCourse(T *Arrguru, n, idx int) {
	//Mengurutkan course dari jumlah siswa yang mengikuti course
	var i, j int
	var temp DataCourse
	i = 1
	for i <= n-1 {
		j = i
		temp = T.Tabguru[idx].CourseGuru[j]
		for j > 0 && temp.Join > T.Tabguru[idx].CourseGuru[j-1].Join {
			T.Tabguru[idx].CourseGuru[j] = T.Tabguru[idx].CourseGuru[j-1]
			j--
		}
		T.Tabguru[idx].CourseGuru[j] = temp
		i++
	}
}

func SortingNamaSiswa_Insertion(T1 *Arrsiswa, n int) {
	//Mengurutkan course guru dari A-Z secara insersi
	var i, j int
	var temp Siswa
	i = 1
	for i <= n-1 {
		j = i
		temp = T1.Tabsiswa[j]
		for j > 0 && temp.nama < T1.Tabsiswa[j-1].nama {
			T1.Tabsiswa[j] = T1.Tabsiswa[j-1]
			j--
		}
		T1.Tabsiswa[j] = temp
		i++
	}
}

func SortingNilaiTugas_Insertion(T *Arrsiswa, n, jumS int) {
	var i, j int
	var temp DataCourse
	i = 1
	for i <= n-1 {
		j = i
		temp = T.Tabsiswa[T.Jumsiswa-1].CourseDiambil[j]
		for j > 0 && temp.NilaiTugas > T.Tabsiswa[T.Jumsiswa-1].CourseDiambil[j-1].NilaiTugas {
			T.Tabsiswa[T.Jumsiswa-1].CourseDiambil[j] = T.Tabsiswa[T.Jumsiswa-1].CourseDiambil[j-1]
			j--
		}
		T.Tabsiswa[T.Jumsiswa-1].CourseDiambil[j] = temp
		i++
	}
}

func SortingNilaiQuiz_Insertion(T *Arrsiswa, n, jumS int) {
	var i, j int
	var temp DataCourse
	i = 1
	for i <= n-1 {
		j = i
		temp = T.Tabsiswa[T.Jumsiswa-1].CourseDiambil[j]
		for j > 0 && temp.NilaiQuiz > T.Tabsiswa[T.Jumsiswa-1].CourseDiambil[j-1].NilaiQuiz {
			T.Tabsiswa[T.Jumsiswa-1].CourseDiambil[j] = T.Tabsiswa[T.Jumsiswa-1].CourseDiambil[j-1]
			j--
		}
		T.Tabsiswa[T.Jumsiswa-1].CourseDiambil[j] = temp
		i++
	}
}

func SortingCourse_Selection(T *Arrguru, n, jumC int) {
	//Mengurutkan course guru dari A-Z secara seleksi
	var i, j, idx int
	var temp DataCourse
	i = 1
	for i <= n-1 {
		idx = i - 1
		j = i
		for j < n {
			if T.Tabguru[jumC].CourseGuru[idx].namaCourse < T.Tabguru[jumC].CourseGuru[j].namaCourse {
				idx = j
			}
		}
		temp = T.Tabguru[jumC].CourseGuru[idx]
		T.Tabguru[jumC].CourseGuru[idx] = T.Tabguru[jumC].CourseGuru[i-1]
		T.Tabguru[jumC].CourseGuru[i-1] = temp
		i++
	}
}

func printSiswa(T1 Arrsiswa) {
	var i int

	fmt.Print("\n Berikut adalah list nama siswa")
	for i = 0; i < T1.Jumsiswa; i++ {
		fmt.Print("\n", T1.Tabsiswa[i].nama)
	}

}

func PrintCourseDiambil(T1 Arrsiswa) {
	var i int
	var e int
	e = T1.Tabsiswa[T1.Jumsiswa-1].jumlahCourseDiambil
	fmt.Print("\n Berikut adalah list nama course yang sudah diambil:")
	for i < e || i < T1.Jumsiswa {
		fmt.Print("\n ", T1.Tabsiswa[T1.Jumsiswa-1].CourseDiambil[i].namaCourse)
		i++
		if i == e {
			T1.Jumsiswa++
		}
	}
}

func PrintNamaCourse(T Arrguru) {
	//Mengoutput isi course sebuah guru ke T.jumguru-1 yang ditandai dengan variable e agar statis
	var i int
	var e int
	e = T.Tabguru[T.Jumguru-1].jumlahCourse
	fmt.Print("\n Berikut adalah data course yang tersedia dengan format Guru, Nama Course")
	for i < e || i < T.Jumguru {
		fmt.Print("\n", T.Tabguru[T.Jumguru-1].nama, " ", T.Tabguru[T.Jumguru-1].CourseGuru[i].namaCourse)
		i++
		if i == e {
			T.Jumguru++
		}
	}
}

func PrintIsiCourse(T Arrguru) {
	//Mengoutput isi course sebuah guru ke T.jumguru-1 yang ditandai dengan variable e agar statis
	var i int
	var e int
	e = T.Tabguru[T.Jumguru-1].jumlahCourse
	fmt.Print("\n Berikut adalah data course yang tersedia dengan format Nama Course, Soal Tugas, Soal Quiz, Pertanyaan Forum:")
	for i < e {
		fmt.Print("\n", T.Tabguru[T.Jumguru-1].CourseGuru[i].namaCourse, " ", T.Tabguru[T.Jumguru-1].CourseGuru[i].Tugas, " ", T.Tabguru[T.Jumguru-1].CourseGuru[i].Quiz, " ", T.Tabguru[T.Jumguru-1].CourseGuru[i].Forum)
		i++
	}
}

func indeksCourseBinary(T1 Arrsiswa, n, JumS int, NamaCourse string) int {
	//Mencari course yang dimiliki oleh siswa dan mengembalikan nilai indeks dari course yang telah diinput
	var ketemu, tengah int
	var kr, kn int
	ketemu = -1
	kr = 0
	kn = n - 1
	for kr <= kn && ketemu == -1 {
		tengah = (kr + kn) / 2
		if NamaCourse < T1.Tabsiswa[JumS].CourseDiambil[tengah].namaCourse {
			kn = tengah - 1
		} else if NamaCourse > T1.Tabsiswa[JumS].CourseDiambil[tengah].namaCourse {
			kr = tengah + 1
		} else {
			ketemu = tengah
		}
	}
	return ketemu
}

func carinamaSiswaBinary(T1 Arrsiswa, n int, namaSiswa string) int {
	//Mencari nama siswa yang telah terinput sebelumnya di array dan mengembalikan indeks dari nama guru yang sesuai (jika ada)
	var ketemu, tengah int
	var kr, kn int
	ketemu = -1
	kr = 0
	kn = n - 1
	for kr <= kn && ketemu == -1 {
		tengah = (kr + kn) / 2
		if namaSiswa < T1.Tabsiswa[tengah].nama {
			kn = tengah - 1
		} else if namaSiswa > T1.Tabsiswa[tengah].nama {
			kr = tengah + 1
		} else {
			ketemu = tengah
		}
	}
	return ketemu
}

func cariNamaguruBinary(T Arrguru, n int, Namaguru string) int {
	//Mencari nama guru yang telah terinput sebelumnya di array dan mengembalikan indeks dari nama guru yang sesuai (jika ada)
	var ketemu, tengah int
	var kr, kn int
	ketemu = -1
	kr = 0
	kn = n - 1
	for kr <= kn && ketemu == -1 {
		tengah = (kr + kn) / 2
		if Namaguru > T.Tabguru[tengah].nama {
			kn = tengah - 1
		} else if Namaguru < T.Tabguru[tengah].nama {
			kr = tengah + 1
		} else {
			ketemu = tengah
		}
	}
	return ketemu
}

func cariCourseSequencial(T Arrguru, namaC string) bool {
	//Nyari nama course di array T.tabguru.courseguru.namacourse
	var i int
	var ketemu bool
	ketemu = false
	i = 0
	for i < T.Tabguru[T.Jumguru-1].jumlahCourse && !ketemu {
		ketemu = T.Tabguru[T.Jumguru-1].CourseGuru[i].namaCourse == namaC
		i++
	}
	return ketemu
}

func indekscourseSequencialSiswa(T Arrsiswa, namaC string) int {
	//Mengembalikan indeks dari suatu course yang dimiliki oleh guru
	var i int
	var ketemu int
	ketemu = -1
	i = 0
	for i < T.Tabsiswa[T.Jumsiswa-1].jumlahCourseDiambil && ketemu == -1 {
		if T.Tabsiswa[T.Jumsiswa-1].CourseDiambil[i].namaCourse == namaC {
			ketemu = i
		}
		i++
	}
	return ketemu

}

func indekscourseSequencial(T Arrguru, namaC string) int {
	//Mengembalikan indeks dari suatu course yang dimiliki oleh guru
	var i int
	var ketemu int
	ketemu = -1
	i = 0
	for i < T.Tabguru[T.Jumguru-1].jumlahCourse && ketemu == -1 {
		if T.Tabguru[T.Jumguru-1].CourseGuru[i].namaCourse == namaC {
			ketemu = i
		}
		i++
	}
	return ketemu

}
