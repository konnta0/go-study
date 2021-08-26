package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Empty(t, randomString(0))
	assert.Len(t, randomString(1), 1)
	assert.Len(t, randomString(2), 2)
	assert.Len(t, randomString(3), 3)
	assert.Len(t, randomString(4), 4)
	assert.Len(t, randomString(5), 5)
	assert.Len(t, randomString(6), 5)
}

func Benchmark0(b *testing.B) {

	b.ResetTimer()
	for i := 0; i < 10000; i++ {
		fmt.Println(randomString(3))
	}
}
