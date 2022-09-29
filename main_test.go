package main

import "testing"

func BenchmarkJsonMarshalInt32s(b *testing.B) {
	marshalInt32s(true)
}

func BenchmarkSelfMarshalInt32s(b *testing.B) {
	marshalInt32s(true)
}
