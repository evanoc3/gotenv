package slices

import (
	"testing"

	"github.com/stretchr/testify/require"
)


func TestTransform(t *testing.T) {
	require := require.New(t)

	require.Equal( Transform([]int{1, 2, 3}, func(x int) int { return x * 2; }), []int{2, 4, 6} )
	require.Equal( Transform([]int{1, 2, 3}, func(x int) rune { return 'a'; }), []rune{'a', 'a', 'a'} )
}


func TestFilter(t *testing.T) {
	require := require.New(t)

	require.Equal( Filter([]int{1, 2, 3, 4}, func(x int) bool { return x % 2 == 0; }), []int{2, 4} )

	require.Equal( Filter([]int{1, 2, 3, 4}, func(x int) bool { return true; }), []int{1, 2, 3, 4} )
	require.Equal( Filter([]int{1, 2, 3, 4}, func(x int) bool { return false; }), []int{} )
}
