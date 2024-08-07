package field

// ---- auto generated by baseTypeBuilder_test.go, NOT modify this file ----

import (
	"github.com/findyu001/go-db-mongo/mongodb/filter"
	"github.com/findyu001/go-db-mongo/mongodb/updater"
)

type Float320F struct {
	*Float320FUpdaterF
	*baseKey
	*Float320FFilterF
}

func (a *Float320F) FullName() string {
	return a.baseKey.FullName()
}

type Float320FUpdaterF struct {
	*baseUpdater
}

type Float320FFilterF struct {
	*baseFilter
}

func NewFloat320F(fName string) *Float320F {
	uper := &Float320FUpdaterF{&baseUpdater{&base{fName}}}
	pri := &baseKey{uper.base}
	flt := &Float320FFilterF{&baseFilter{uper.base}}

	return &Float320F{uper, pri, flt}
}

func (i *Float320FFilterF) Gt(value float32) filter.Filter {
	return filter.CompareByValue(i, filter.GT, value)
}

func (i *Float320FFilterF) Lt(value float32) filter.Filter {
	return filter.CompareByValue(i, filter.LT, value)
}

func (i *Float320FFilterF) GtField(f *Float320F) filter.Filter {
	return filter.CompareByValue(i, filter.GT, f)
}

func (i *Float320FFilterF) LtField(f *Float320F) filter.Filter {
	return filter.CompareByValue(i, filter.LT, f)
}

func (i *Float320FFilterF) In(values []float32) filter.Filter {
	return filter.New(i, "$in", values)
}

func (i *Float320FFilterF) Nin(values []float32) filter.Filter {
	return filter.New(i, "$nin", values)
}

func (i *Float320FUpdaterF) Min(value float32) updater.Updater {
	return updater.New(i, "$min", value)
}

func (i *Float320FUpdaterF) Max(value float32) updater.Updater {
	return updater.New(i, "$max", value)
}

func (i *Float320FUpdaterF) Set(value float32) updater.Updater {
	return updater.New(i, "$set", value)
}

func (i *Float320FUpdaterF) SetOnIns(value float32) updater.Updater {
	return updater.New(i, "$setOnInsert", value)
}
