package optparse

import (
//	"errors"
)

type Opt struct {
	Short    string
	Long     string
	Desc     string
	Sample   string
	Required bool
	Value    string
}

type Options map[string]*Opt

func NewOptions() *Options {
	o := make(Options)
	return &o
}

//func (opt *Opt) SetShort(v string) *Opt {
//	opt.Short = v
//	return opt
//}

func (opt *Opt) SetLong(v string) *Opt {
	opt.Long = v
	return opt
}

func (opt *Opt) SetDesc(v string) *Opt {
	opt.Desc = v
	return opt
}

func (opt *Opt) SetSample(v string) *Opt {
	opt.Sample = v
	return opt
}

func (opt *Opt) SetRequired(v bool) *Opt {
	opt.Required = v
	return opt
}

func (opt *Opt) SetValue(v bool) *Opt {
	opt.Value = v
	return opt
}

func (opts Options) Add(k string) *Opt {
	o := new(Opt)
	//	o.SetShort(k)
	o.Short = k
	opts[k] = o

	return o
}

func (opts Options) Defined(k string) bool {
	_, ok := opts[k]
	return ok
}

func (opts Options) ParseArgs(input string) error {
	return nil
}
