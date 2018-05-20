package sort

import (
	"testing"
	"sort"
	"fmt"
	"math/rand"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	fmt.Println(data)

	Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
	fmt.Println(data)
}
//
//func main() {
//
//}
