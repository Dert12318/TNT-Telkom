// Soal No 1
// 2. Anda diminta untuk membuat sebuah function dimana function tersebut berfungsi untuk menentukan apakah dari dua data string yang diberikan membutuhkan sekali proses edit atau lebih. Jika lebih dari sekali proses edit berarti function tersebut akan mengembalikan response False, sedangkan jika hanya sekali proses edit maka function tersebut akan mengembalikan response True. Proses edit di sini dapat berarti melakukan insert sebuah character, remove sebuah character, atau replace sebuah character.
// Contoh
// GIVEN INPUT 1 → telkom
// GIVEN INPUT 2 → telecom
// RESULT → False
//
// GIVEN INPUT 1 → telkom
// GIVEN INPUT 2 → tlkom
// RESULT → True

package main

import "fmt"

func Smallest(input1 int, input2 int) int {
	if input1 > input2 {
		return input2
	} else {
		return input1
	}
}

// Logic edit hanya bisa tambah atau hilangkan kata tidak bisa edit kata atau pindah kata
func isLookLike(input1 string, input2 string) bool {
	var Mybool bool
	for i := 0; i < Smallest(len(input1), len(input2)); i++ {
		if input1[i] != input2[i] {
			temp := input1[:i] + string(input2[i]) + input1[i:]
			temp2 := input2[:i] + string(input1[i]) + input2[i:]
			fmt.Println(temp)
			fmt.Println(temp2)
			if temp == input2 || temp2 == input1 {
				fmt.Println("masuk sini")
				Mybool = true
				break
			} else {
				fmt.Println("masuk sini sini")
				Mybool = false
				break
			}
		}
	}
	return Mybool
}

func main() {
	//Declare variable
	var A string
	var B string
	fmt.Println("Masukan Kalimat Pertama ")
	fmt.Scanln(&A)
	fmt.Println("Masukan Kalimat Kedua ")
	fmt.Scanln(&B)
	fmt.Println(isLookLike(A, B))
}
