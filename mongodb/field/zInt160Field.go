package field

// ---- auto generated by baseTypeBuilder_test.go, NOT modify this file ----

import (
	"github.com/findyu001/go-db-mongo/mongodb/filter"
	"github.com/findyu001/go-db-mongo/mongodb/updater"
	"go.mongodb.org/mongo-driver/bson"
)

type Int160F struct {
	*Int160FUpdaterF
	*baseKey
	*Int160FFilterF
}

func (a *Int160F) FullName() string {
	return a.baseKey.FullName()
}

type Int160FUpdaterF struct {
	*baseUpdater
}

type Int160FFilterF struct {
	*baseFilter
}

func NewInt160F(fName string) *Int160F {
	uper := &Int160FUpdaterF{&baseUpdater{&base{fName}}}
	pri := &baseKey{uper.base}
	flt := &Int160FFilterF{&baseFilter{uper.base}}

	return &Int160F{uper, pri, flt}
}

func (i *Int160FFilterF) Mod(divisor, remainder int16) filter.Filter {
	return filter.New(i, "$mod", bson.A{divisor, remainder})
}

func (i *Int160FUpdaterF) Inc(num int16) updater.Updater {
	return updater.New(i, "$inc", num)
}

func (i *Int160FUpdaterF) Mul(num int16) updater.Updater {
	return updater.New(i, "$mul", num)
}

func (i *Int160FFilterF) Eq(value int16) filter.Filter {
	return filter.CompareByValue(i, filter.EQ, value)
}

func (i *Int160FFilterF) Ne(value int16) filter.Filter {
	return filter.CompareByValue(i, filter.NE, value)
}

func (i *Int160FFilterF) NeField(f *Int160F) filter.Filter {
	return filter.CompareByField(i, filter.NE, f)
}

func (i *Int160FFilterF) EqField(f *Int160F) filter.Filter {
	return filter.CompareByField(i, filter.EQ, f)
}

func (i *Int160FFilterF) Gte(value int16) filter.Filter {
	return filter.CompareByValue(i, filter.GTE, value)
}

func (i *Int160FFilterF) Lte(value int16) filter.Filter {
	return filter.CompareByValue(i, filter.LTE, value)
}

func (i *Int160FFilterF) GteField(f *Int160F) filter.Filter {
	return filter.CompareByValue(i, filter.GTE, f)
}

func (i *Int160FFilterF) LteField(f *Int160F) filter.Filter {
	return filter.CompareByValue(i, filter.LTE, f)
}

func (i *Int160FFilterF) Gt(value int16) filter.Filter {
	return filter.CompareByValue(i, filter.GT, value)
}

func (i *Int160FFilterF) Lt(value int16) filter.Filter {
	return filter.CompareByValue(i, filter.LT, value)
}

func (i *Int160FFilterF) GtField(f *Int160F) filter.Filter {
	return filter.CompareByValue(i, filter.GT, f)
}

func (i *Int160FFilterF) LtField(f *Int160F) filter.Filter {
	return filter.CompareByValue(i, filter.LT, f)
}

func (i *Int160FFilterF) In(values []int16) filter.Filter {
	return filter.New(i, "$in", values)
}

func (i *Int160FFilterF) Nin(values []int16) filter.Filter {
	return filter.New(i, "$nin", values)
}

func (i *Int160FUpdaterF) Min(value int16) updater.Updater {
	return updater.New(i, "$min", value)
}

func (i *Int160FUpdaterF) Max(value int16) updater.Updater {
	return updater.New(i, "$max", value)
}

func (i *Int160FUpdaterF) Set(value int16) updater.Updater {
	return updater.New(i, "$set", value)
}

func (i *Int160FUpdaterF) SetOnIns(value int16) updater.Updater {
	return updater.New(i, "$setOnInsert", value)
}
