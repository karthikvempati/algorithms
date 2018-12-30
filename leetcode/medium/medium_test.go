package lcmedium

import (
	"fmt"
	"testing"
)

func TestIslands(t *testing.T) {
	a := [][]byte{
		{49, 49, 48, 48, 48},
		{49, 49, 48, 48, 48},
		{48, 48, 49, 48, 48},
		{49, 48, 48, 49, 49},
		{48, 49, 48, 49, 49},
	}
	//fmt.Println(a)
	count := numIslands(a)
	fmt.Println(count)
}
