package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Part1(t *testing.T) {
	testCourse := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	assert.Equal(t, 150, Part1(testCourse))
}

func Test_Part2(t *testing.T) {
	testCourse := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	assert.Equal(t, 900, Part2(testCourse))
}
