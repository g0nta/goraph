package goraph

import (
	"testing"
)

func TestNewGraph(t *testing.T) {
	g := NewGraph()
	if g == nil {
		t.Error("NewGraph has faild.")
	}
	t.Log("TestNewGraph has successed.")
}
