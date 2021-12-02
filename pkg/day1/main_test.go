package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Part1(t *testing.T) {
	assert.Equal(t, 7, Part1([]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}))
}
func Test_Part2(t *testing.T) {
	assert.Equal(t, 5, Part2([]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}))
}
