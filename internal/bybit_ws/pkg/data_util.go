package pkg

import (
	"go-clean-arch/internal/bybit_ws/params"
	"sort"
)

func UpdateAsks(data *params.BidAsk, newData params.BidAsk) {
	for i := 0; i < len(newData.A); i++ {
		if item := findItemToAsks(data, newData.A[i][0]); item == false {
			if newData.A[i][1] != "0" {
				data.A = append(data.A, [][]string{{newData.A[i][0], newData.A[i][1]}}...)
			}
		}
	}

	if len(data.A) > len(newData.A) || len(data.A) == len(newData.A) {
		for i := 0; i < len(newData.A); i++ {
			for j := 0; j < len(newData.A[i]); j++ {
				if len(data.A[i]) > 0 {
					if data.A[i][0] == newData.A[i][0] {
						data.A[i][1] = newData.A[i][1]
					}
				}
			}
		}
	} else if len(newData.A) > len(data.A) || len(data.A) == len(newData.A) {
		for i := 0; i < len(data.A); i++ {
			for j := 0; j < len(data.A[i]); j++ {
				if len(data.A[i]) > 0 {
					if data.A[i][0] == newData.A[i][0] {
						data.A[i][1] = newData.A[i][1]
					}
				}
			}
		}
	}
}
func UpdateBids(data *params.BidAsk, newData params.BidAsk) {

	for i := 0; i < len(newData.B); i++ {
		if item := findItemToBids(data, newData.B[i][0]); item == false {
			if newData.B[i][1] != "0" {
				data.B = append(data.B, [][]string{{newData.B[i][0], newData.B[i][1]}}...)
			}
		}
	}

	if len(data.B) > len(newData.B) || len(data.B) == len(newData.B) {
		for i := 0; i < len(newData.B); i++ {
			for j := 0; j < len(newData.B[i]); j++ {
				//fmt.Printf("Old Element at1 [%d][%d]: %d\n", i, j, ba.Data.B[i])
				//fmt.Printf("Old Element at2 [%d][%d]: %d\n", i, j, ba.Data.B[i][1])
				if len(data.B[i]) > 0 {
					if data.B[i][0] == newData.B[i][0] {
						data.B[i][1] = newData.B[i][1]
					}
				}
				//fmt.Printf("New Element at [%d][%d]: %d\n", i, j, ba.Data.B[i][1])
			}
		}
	} else if len(newData.B) > len(data.B) || len(data.B) == len(newData.B) {

		for i := 0; i < len(data.B); i++ {
			for j := 0; j < len(data.B[i]); j++ {
				if len(data.B[i]) > 0 {
					if data.B[i][0] == newData.B[i][0] {
						data.B[i][1] = newData.B[i][1]
					}
				}
			}
		}
	}
}
func findItemToBids(data *params.BidAsk, B string) bool {
	for i, _ := range data.B {
		//fmt.Println("findItemToBids1:", len(ba.Data.B[i]))
		//fmt.Println("findItemToBids2:", len(ba.Data.B[i][0]))
		if len(data.B[i]) > 0 {
			if data.B[i][0] == B {
				return true
			}
		}
	}
	//for i := 0; i < len(ba.Data.B); i++ {
	//	fmt.Println("findItemToBids1:", len(ba.Data.B), i)
	//	fmt.Println("findItemToBids2:", ba.Data.B[i][0])
	//	if ba.Data.B[i][0] == B {
	//		return true
	//	}
	//}
	return false
}
func findItemToAsks(data *params.BidAsk, A string) bool {
	for i, _ := range data.A {
		if data.A[i][0] == A {
			return true
		}
	}
	return false
}

func RemoveDuplicatesAsks(data *params.BidAsk) {
	sort.Slice(data.A, func(i, j int) bool {
		if i >= len(data.A) || j >= len(data.A) {
			return false
		}
		return data.A[i][1] > data.A[j][1]
	})
	for i, _ := range data.A {
		if data.A[i][1] == "0" {
			data.A = data.A[:i]
			//fmt.Println(i, "  ", ba.Data.B)
			return
		}
	}
}
func RemoveDuplicatesBids(data *params.BidAsk) {

	sort.Slice(data.B, func(i, j int) bool {
		if i >= len(data.B) || j >= len(data.B) {
			return false
		}
		return data.B[i][1] > data.B[j][1]
	})

	for i, _ := range data.B {
		if data.B[i][1] == "0" {
			data.B = data.B[:i]
			//fmt.Println(i, "  ", ba.Data.B)
			return
		}
	}
}
