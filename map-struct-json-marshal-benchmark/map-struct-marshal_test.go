package main

import "testing"

func BenchmarkRunstruct(b *testing.B) {
	Runstruct()
}

func BenchmarkRunmap(b *testing.B) {
	Runmap()
}
