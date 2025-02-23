package integer

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(6, 7)
	want := 13
	if got != want {
		t.Errorf("Wanted %d got %d", want, got)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
