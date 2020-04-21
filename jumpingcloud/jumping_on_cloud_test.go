package jumpingcloud

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func Test_Input_Cloud_7_Should_Be_4(t *testing.T) {
	expected := 4

	cloud := []int{0, 0, 1, 0, 0, 1, 0}

	actual := jumpingOnClouds(cloud)

	assert.Equal(t, expected, actual)
}

func Test_Input_Cloud_6_Should_Be_3(t *testing.T) {
	expected := 3

	cloud := []int{0, 0, 0, 0, 1, 0}

	actual := jumpingOnClouds(cloud)

	assert.Equal(t, expected, actual)
}

func Test_Input_Cloud_6_Should_Be_2(t *testing.T) {
	expected := 2

	cloud := []int{0, 0, 0, 0}

	actual := jumpingOnClouds(cloud)

	assert.Equal(t, expected, actual)
}
