package field

// ---- auto generated by builder.go, NOT modify this file ----

import (
	fmt "fmt"
	filter "github.com/findyu001/go-db-mongo/mongodb/filter"
	updater "github.com/findyu001/go-db-mongo/mongodb/updater"
	bson "go.mongodb.org/mongo-driver/bson"
)

type Uint161Field struct {
	*Array
}

func NewUint161Field(fName string) *Uint161Field {
	return &Uint161Field{NewArray(fName)}
}

func (i *Uint161Field) EleAt(index int) *Uint160F {
	return NewUint160F(fmt.Sprintf("%s.%d", i.FullName(), index))
}

// 数组的某一个元素，有时也可以理解为 数组的任何一个元素
func (i *Uint161Field) EleOne() *Uint160F {
	return NewUint160F(i.FullName())
}

// update 操作中被filter匹配的那第一个元素
func (i *Uint161Field) EleThat() *Uint160FUpdaterF {
	return NewUint160F(i.FullName() + ".$").Uint160FUpdaterF
}

func (i *Uint161Field) EleAll() *Uint160FUpdaterF {
	return NewUint160F(i.FullName() + ".$[]").Uint160FUpdaterF
}

func (i *Uint161Field) EleByFid(identifier string) *Uint160FUpdaterF {
	return NewUint160F(fmt.Sprintf("%s.$[%s]", i.FullName(), i.FullName()+identifier)).Uint160FUpdaterF
}

func (i *Uint161Field) DeclFid(identifier string) *Uint160FFilterF {
	return NewUint160F(i.FullName() + identifier).Uint160FFilterF
}

func (i *Uint161Field) Include(a []uint16) filter.Filter {
	return filter.New(i, "$all", a)
}

func (i *Uint161Field) Eq(a []uint16) filter.Filter {
	return filter.CompareByValue(i, filter.EQ, a)
}

func (i *Uint161Field) Set(a []uint16) updater.Updater {
	return updater.New(i, "$set", a)
}

func (i *Uint161Field) AddToSet(value uint16) updater.Updater {
	return updater.New(i, "$addToSet", value)
}

func (i *Uint161Field) AddToSetValues(a []uint16) updater.Updater {
	return updater.New(i, "$addToSet", bson.M{"$each": a})
}

func (i *Uint161Field) Pull(value uint16) updater.Updater {
	return updater.New(i, "$pull", value)
}

func (i *Uint161Field) PullAll(a []uint16) updater.Updater {
	return updater.New(i, "$pullAll", a)
}

func (i *Uint161Field) Push(value uint16) updater.Updater {
	return updater.New(i, "$push", value)
}

func (i *Uint161Field) PushByModifier(m updater.PushModifier, each []uint16) updater.Updater {
	return updater.PushByModifier(i, m, each)
}
