package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {

	b := []int{1, 2}
	result := Sum(b[0], b[1])

	assert.Equal(t, 0, result)

}
