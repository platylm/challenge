package repeatstring

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func Test_Input_ABC_And_Number_10_Should_Be_7(t *testing.T) {
	expected := 7

	inputString := "aba"
	number := 10

	actual := repeatedString(inputString, number)

	assert.Equal(t, expected, actual)
}

func Test_Input_A_And_Number_1M_Should_Be_1M(t *testing.T) {
	expected := 1000000

	inputString := "a"
	number := 1000000

	actual := repeatedString(inputString, number)

	assert.Equal(t, expected, actual)
}

func Test_Input_A_And_Number_Should_Be_1M(t *testing.T) {
	expected := 588525

	inputString := "aab"
	number := 882787

	actual := repeatedString(inputString, number)

	assert.Equal(t, expected, actual)
}
