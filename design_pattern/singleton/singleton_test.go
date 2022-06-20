package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	c1 := Singleton()
	c2 := Singleton()
	if c1 != c2 {
		t.Fatal("c1 != c2")
	}
}
