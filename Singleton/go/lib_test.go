package lib

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	s1 := NewSingleton()
	s2 := NewSingleton()
	if s1 != s2 {
		t.Error("Singleton pattern failed")
	}
}
