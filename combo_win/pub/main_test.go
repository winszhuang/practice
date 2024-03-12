package main

import "testing"

func Test_genResult(t *testing.T) {
	for i := 0; i < 20; i++ {
		result := genResult()
		println(result)
	}
}
