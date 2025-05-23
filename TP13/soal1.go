package main
import "fmt"

const NMAX int = 99
type tabInt [NMAX]int

func main(){
	var data tabInt
	var nData int

	fmt.Scan(&nData)
	bacaData(&data, nData)
	cetakDataa(data, nData)
	SelectionSort(&data, nData)
	cetakDataa(data, nData)
}

func bacaData(A *tabInt, N int){
	var i int

	for i = 0; i < N; i++ {
		fmt.Scan(&A[i])
	}
}

func cetakDataa(A tabInt, N int){
	var i int

	for i = 0; i < N; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println("")
}

func SelectionSort(A *tabInt, N int) {
	var i, idx, pass int
	var temp int

	pass = 1

	for pass < N {
		idx = pass - 1
		i = pass

		for i < N {
			if A[i] > A[idx] {
				idx = i 
			}
			i = i + 1
		}
		temp = A[pass - 1]
		A[pass - 1] = A[idx]
		A[idx] = temp
		pass = pass + 1
	}
}