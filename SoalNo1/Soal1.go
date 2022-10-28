// Soal No 1
// 1. Di Indonesia, ada pecahan mata uang rupiah, yaitu :
// * 100.000,-
// * 50.000,-
// * 20.000,-
// * 10.000,-
// * 5.000,-
// * 2.000,-
// * 1.000,-
// * 500,-
// * 200,-
// * 100,-
// Buatlah sebuah fungsi untuk menghitung berapa lembar pecahan yang harus dikeluarkan dari input harga (dengan pembulatan ke atas jita punya harga pecahan antara 1 sampai 99)
// Input : 145.000
// Output:
//             {
//                  “Rp. 100.000”:1,
//                  “Rp. 20.000”:2,
//                  “Rp. 5.000”:1,
//             }
//             Input: 2050
//             Ouput:
//             {
//                  “Rp. 2.000”:1,
//                  “Rp. 100”:1,
//             }

package main

import (
	"fmt"
	"math"
	"sort"
)

func roundFloat(val float64, precision float64) float64 {
	ratio := math.Round(val / math.Pow(10, precision))
	return ratio * math.Pow(10, precision)
}

func main() {
	//Declare variable
	Pecahan := map[float64]int{100000: 0, 50000: 0, 20000: 0, 10000: 0, 5000: 0, 2000: 0, 1000: 0, 500: 0, 200: 0, 100: 0}
	keys := make([]float64, 0, len(Pecahan))
	var input float64
	//Logic
	fmt.Println("Masukan Nonimal ")
	fmt.Scanln(&input)
	X := roundFloat(input, 2)
	for k := range Pecahan {
		keys = append(keys, k)
	}
	sort.Float64s(keys)
	for i := len(keys) - 1; i >= 0; i-- {
		if X/keys[i] >= 1 {
			// fmt.Println("masuk ", i, "value", X/keys[i])
			Pecahan[keys[i]] = int(X / keys[i]) // jika nilai ingin disimpan ke maps jika tidak bisa dicomment
			//Print Result
			fmt.Println(keys[i], ":", int(X/keys[i]))
			X = math.Mod(X, keys[i])
		}
	}
}
