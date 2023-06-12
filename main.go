package main

import "fmt"

// ini function main
/*
	ini function main sample komen berbeda
*/
func main() {
	fmt.Println("hello world")
	fmt.Println("hello", "world")

	// contoh variabel panjang
	var nama = "ibam"
	var nama2 string = "budi"
	fmt.Println(nama, nama2)

	// contoh variabel golang short declaration
	orangPertama := "saya" // camelcase
	orang_kedua := "dia"   // snake case
	var tiga int = 3
	tigax := 3

	malamHari := true

	fmt.Println(orangPertama, orang_kedua)
	fmt.Println(tiga, tigax)
	fmt.Printf("Orang pertama adalah %s dan Orang kedua adalah %s\n", orangPertama, orang_kedua)
	fmt.Printf("%T \n", tiga)

	fmt.Println(malamHari)

	// multiple declaration
	orangKetiga, orangKeempat := "aku", "kamu"
	fmt.Println(orangKetiga, orangKeempat)

	// underscore variable
	orangKelima, _ := daftarOrang()
	fmt.Println(orangKelima)

	fizzbuzz()
}

// function sample
func daftarOrang() (string, string) {
	return "aku", "kamu"
}

func fizzbuzz() {
	i := 1

	for {
		if i == 3 || i == 5 || i == 9 || i == 12 {
			fmt.Println("Fizz")
		} else if i == 5 || i == 10 {
			fmt.Println("Buzz")
		} else if i == 15 {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println(i)
		}

		if i == 15 {
			break
		}
		i++
	}
}
