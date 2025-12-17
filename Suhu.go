package main

import "fmt"

func main() {
	var suhu float64
	var hasil float64
	var lanjut string
	var pilihan int
	var lanjutProgram bool
	var konversiBerhasil bool
	var valid bool

	lanjutProgram = true
	konversiBerhasil = false
	valid = false

	fmt.Println("=== PROGRAM KONVERSI SUHU ===")

	for lanjutProgram == true {
		fmt.Println("---------------------------------")
		fmt.Println("PILIHAN KONVERSI SUHU:")
		fmt.Println("1. Celcius ke Fahrenheit")
		fmt.Println("2. Fahrenheit ke Celcius")
		fmt.Println("3. Celcius ke Kelvin")
		fmt.Println("4. Kelvin ke Celcius")
		fmt.Println("5. Celcius ke Reamur")
		fmt.Println("6. Reamur ke Celcius")
		fmt.Println("7. Keluar")
		fmt.Print("\nPilihan Anda (1-7): ")

		fmt.Scan(&pilihan)
		if pilihan == 7 {
			lanjutProgram = false
			fmt.Println("Terima kasih! Program selesai.")

		} else if pilihan >= 1 && pilihan <= 6 {
			fmt.Print("Masukkan nilai suhu: ")
			fmt.Scan(&suhu)

			if pilihan == 1 {
				hasil = (suhu * 9.0 / 5.0) + 32.0
				konversiBerhasil = true

				fmt.Print("HASIL KONVERSI: ")
				fmt.Print(suhu)
				fmt.Print("°C = ")
				fmt.Print(hasil)
				fmt.Println("°F")

			} else if pilihan == 2 {
				hasil = (suhu - 32.0) * 5.0 / 9.0
				konversiBerhasil = true

				fmt.Print("HASIL KONVERSI: ")
				fmt.Print(suhu)
				fmt.Print("°F = ")
				fmt.Print(hasil)
				fmt.Println("°C")

			} else if pilihan == 3 {
				hasil = suhu + 273.15
				konversiBerhasil = true

				fmt.Print("HASIL KONVERSI: ")
				fmt.Print(suhu)
				fmt.Print("°C = ")
				fmt.Print(hasil)
				fmt.Println(" K")

			} else if pilihan == 4 {
				hasil = suhu - 273.15
				konversiBerhasil = true

				fmt.Print("HASIL KONVERSI: ")
				fmt.Print(suhu)
				fmt.Print(" K = ")
				fmt.Print(hasil)
				fmt.Println("°C")

			} else if pilihan == 5 {
				hasil = suhu * 4.0 / 5.0
				konversiBerhasil = true

				fmt.Print("HASIL KONVERSI: ")
				fmt.Print(suhu)
				fmt.Print("°C = ")
				fmt.Print(hasil)
				fmt.Println("°R")

			} else if pilihan == 6 {
				hasil = suhu * 5.0 / 4.0
				konversiBerhasil = true

				fmt.Print("HASIL KONVERSI: ")
				fmt.Print(suhu)
				fmt.Print("°R = ")
				fmt.Print(hasil)
				fmt.Println("°C")
			}

			if konversiBerhasil == true {
				fmt.Print("\nLakukan konversi lagi? (y/n): ")
				fmt.Scan(&lanjut)

				if lanjut == "y" || lanjut == "Y" {
					lanjutProgram = true
				} else if lanjut == "n" || lanjut == "N" {
					lanjutProgram = false
				} else {
					fmt.Println("Input tidak jelas, program dilanjutkan")
					lanjutProgram = true
				}
			}
		} else {
			fmt.Println("PILIHAN TIDAK VALID!")
			fmt.Println("Harap masukkan angka 1 sampai 7")
			valid = false
			for valid == false {
				fmt.Print("Masukkan pilihan lagi (1-7): ")
				fmt.Scan(&pilihan)

				if pilihan >= 1 && pilihan <= 7 {
					valid = true
					if pilihan == 7 {
						lanjutProgram = false
					}
				} else {
					fmt.Println("Masih tidak valid!")
					valid = false
				}
			}
		}
	}
	fmt.Println("\n=== PROGRAM SELESAI ===")
}

