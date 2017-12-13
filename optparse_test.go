package optparse

import (
	"testing"
)

func TestMakeOptions(t *testing.T) {
	opts := NewOptions()
	opts.Add("t").SetDesc("A test option")
	t.Log(opts.Defined("t"))
}
