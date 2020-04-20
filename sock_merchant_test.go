package merchant

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func Test_Input_Number_9_Array_Number_Should_Be_3(t *testing.T) {
	expected := 3

	number := 9
	arrayNumber := []int{10, 20, 20, 10, 10, 30, 50, 10, 20}

	actual := sockMerchant(number, arrayNumber)

	assert.Equal(t, expected, actual)
}

func Test_Input_Number_4_Array_Number_Should_Be_2(t *testing.T) {
	expected := 2

	number := 4
	arrayNumber := []int{10, 20, 10, 20}

	actual := sockMerchant(number, arrayNumber)

	assert.Equal(t, expected, actual)
}

func Test_Input_Number_10_Array_Number_Should_Be_2(t *testing.T) {
	expected := 2

	number := 10
	arrayNumber := []int{1, 1, 3, 1, 2, 1, 3, 3, 3, 3}

	actual := sockMerchant(number, arrayNumber)

	assert.Equal(t, expected, actual)
}
