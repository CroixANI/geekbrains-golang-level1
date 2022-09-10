package numbers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitNumber(t *testing.T) {
	units, tens, hundreds, err := SplitNumber(923)
	assert.Equal(t, 3, units, "units should be equal")
	assert.Equal(t, 2, tens, "units should be equal")
	assert.Equal(t, 9, hundreds, "units should be equal")
	assert.Nil(t, err, "there should be no error")
}

func ExampleSplitNumber() {
	in := 932

	units, tens, hundreds, err := SplitNumber(in)
	if err != nil {
		panic(err)
	}

	fmt.Println(units)
	fmt.Println(tens)
	fmt.Println(hundreds)
	// output:
	// 2
	// 3
	// 9
}
