package main

import "fmt"

func Soal1(angka1 int) {
	if angka1%2 == 0 {
		fmt.Println("Bukan Bilangan Prima")
		return
	} else if angka1 < 2 {
		fmt.Println("Angka Harus lebih dari 0 atau 1")
		return
	} else {
		fmt.Println("Bilangan prima")
		return
	}
}

func Soal2(nama string) {
	switch nama {
	case "KPOP":
		fmt.Println("Blackpink, 2NE1, BTS")
	case "INDO":
		fmt.Println("Armada, Kerispatih, Kekeyi")
	case "AMERIKA":
		fmt.Println("Avenged, Avril Lavigne, Dream Theater")
	default:
		fmt.Println("Ga ada datanya :(")
	}
}

func main() {
	// Soal pertama func Soal1 baca function nya ya flow nya gimana

	Soal1(1)    //Ini outputnya apa?
	Soal1(2)    // Ini juga
	Soal1(73)   // Hmmmmm
	Soal1(83)   // Ini prima bukan?
	Soal1(99)   //Jangan googling ya wkwk
	Soal1(1000) // Ini Juga
	/*-------------------------------------------------------------------------------------------------------------------*/

	//Soal kedua func Soal2
	Soal2("Jawa")
	Soal2("KPOP")
	Soal2("SUNDA")
	Soal2("INDO")
}
