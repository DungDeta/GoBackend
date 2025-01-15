package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddOne(t *testing.T) {
	// var (
	// 	input  = 1
	// 	output = 3
	// )
	// actual := AddOne(1)
	// if actual != output {
	// 	t.Errorf("AddOne(%d) = %d; want %d", input, actual, output)
	// }
	assert.Equal(t, 3, AddOne(1))
}

func TestRequire(t *testing.T) {
	require.Equal(t, 3, 2)
	fmt.Println("This line will not be printed")
}

func TestAssert(t *testing.T) {
	assert.Equal(t, 3, 2)
	fmt.Println("This line will be printed")
}
