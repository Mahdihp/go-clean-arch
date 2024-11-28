package params

import (
	"sort"
)

type SpotOrderBook struct {
	Ts    int64  `json:"ts"`    // The timestamp (ms) that the system generates the data
	Cts   int64  `json:"cts"`   //The timestamp from the match engine when this orderbook data is produced.
	Topic string `json:"topic"` //Topic name
	Type  string `json:"type"`  //Data type. snapshot,delta
	Data  BidAsk `json:"data"`
}
type BidAsk struct {
	U   int        `json:"u"`   //Update ID
	Seq int64      `json:"seq"` //Cross sequence
	S   string     `json:"s"`   //Symbol name
	B   [][]string `json:"b"`   //Bids. For snapshot stream, the element is sorted by price in descending order
	A   [][]string `json:"a"`   //Asks. For snapshot stream, the element is sorted by price in ascending order
}

func (ba *SpotOrderBook) UpdateSnapShot(data BidAsk, limitOrderBook int) {
	ba.updateBids(data)
	ba.updateAsks(data)
	ba.removeDuplicatesBids()
	ba.removeDuplicatesAsks()
	if len(ba.Data.A) > limitOrderBook {
		ba.Data.A = ba.Data.A[:limitOrderBook]
	}
	if len(ba.Data.B) > limitOrderBook {
		ba.Data.B = ba.Data.B[:limitOrderBook]
	}
}

func (ba *SpotOrderBook) removeDuplicatesAsks() {
	sort.Slice(ba.Data.A, func(i, j int) bool {
		if i >= len(ba.Data.A) || j >= len(ba.Data.A) {
			return false
		}
		return ba.Data.A[i][1] > ba.Data.A[j][1]
	})
	for i, _ := range ba.Data.A {
		if ba.Data.A[i][1] == "0" {
			ba.Data.A = ba.Data.A[:i]
			//fmt.Println(i, "  ", ba.Data.B)
			return
		}
	}
}
func (ba *SpotOrderBook) removeDuplicatesBids() {

	sort.Slice(ba.Data.B, func(i, j int) bool {
		if i >= len(ba.Data.B) || j >= len(ba.Data.B) {
			return false
		}
		return ba.Data.B[i][1] > ba.Data.B[j][1]
	})

	for i, _ := range ba.Data.B {
		if ba.Data.B[i][1] == "0" {
			ba.Data.B = ba.Data.B[:i]
			//fmt.Println(i, "  ", ba.Data.B)
			return
		}
	}
}
func (ba *SpotOrderBook) updateAsks(data BidAsk) {
	for i := 0; i < len(data.A); i++ {
		if item := ba.findItemToAsks(data.A[i][0]); item == false {
			if data.A[i][1] != "0" {
				ba.Data.A = append(ba.Data.A, [][]string{{data.A[i][0], data.A[i][1]}}...)
			}
		}
	}

	if len(ba.Data.A) > len(data.A) || len(ba.Data.A) == len(data.A) {
		for i := 0; i < len(data.A); i++ {
			for j := 0; j < len(data.A[i]); j++ {
				if len(ba.Data.A[i]) > 0 {
					if ba.Data.A[i][0] == data.A[i][0] {
						ba.Data.A[i][1] = data.A[i][1]
					}
				}
			}
		}
	} else if len(data.A) > len(ba.Data.A) || len(ba.Data.A) == len(data.A) {
		for i := 0; i < len(ba.Data.A); i++ {
			for j := 0; j < len(ba.Data.A[i]); j++ {
				if len(ba.Data.A[i]) > 0 {
					if ba.Data.A[i][0] == data.A[i][0] {
						ba.Data.A[i][1] = data.A[i][1]
					}
				}
			}
		}
	}
}
func (ba *SpotOrderBook) updateBids(data BidAsk) {

	for i := 0; i < len(data.B); i++ {
		if item := ba.findItemToBids(data.B[i][0]); item == false {
			if data.B[i][1] != "0" {
				ba.Data.B = append(ba.Data.B, [][]string{{data.B[i][0], data.B[i][1]}}...)
			}
		}
	}

	if len(ba.Data.B) > len(data.B) || len(ba.Data.B) == len(data.B) {
		for i := 0; i < len(data.B); i++ {
			for j := 0; j < len(data.B[i]); j++ {
				//fmt.Printf("Old Element at1 [%d][%d]: %d\n", i, j, ba.Data.B[i])
				//fmt.Printf("Old Element at2 [%d][%d]: %d\n", i, j, ba.Data.B[i][1])
				if len(ba.Data.B[i]) > 0 {
					if ba.Data.B[i][0] == data.B[i][0] {
						ba.Data.B[i][1] = data.B[i][1]
					}
				}
				//fmt.Printf("New Element at [%d][%d]: %d\n", i, j, ba.Data.B[i][1])
			}
		}
	} else if len(data.B) > len(ba.Data.B) || len(ba.Data.B) == len(data.B) {

		for i := 0; i < len(ba.Data.B); i++ {
			for j := 0; j < len(ba.Data.B[i]); j++ {
				if len(ba.Data.B[i]) > 0 {
					if ba.Data.B[i][0] == data.B[i][0] {
						ba.Data.B[i][1] = data.B[i][1]
					}
				}
			}
		}
	}
}
func (ba *SpotOrderBook) findItemToBids(B string) bool {

	for i, _ := range ba.Data.B {
		//fmt.Println("findItemToBids1:", len(ba.Data.B[i]))
		//fmt.Println("findItemToBids2:", len(ba.Data.B[i][0]))
		if len(ba.Data.B[i]) > 0 {
			if ba.Data.B[i][0] == B {
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
func (ba *SpotOrderBook) findItemToAsks(A string) bool {
	for i, _ := range ba.Data.A {
		if ba.Data.A[i][0] == A {
			return true
		}
	}
	return false
}
