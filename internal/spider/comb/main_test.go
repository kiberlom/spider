package comb

import (
	"fmt"
	"testing"
)

func BenchmarkNext(b *testing.B) {
	sn := "0"

	for i := 0; i < b.N; i++ {
		s, err := Next(sn)
		if err != nil {
			fmt.Println(err)
			continue
		}
		sn = s

		fmt.Println(s)
	}
}
