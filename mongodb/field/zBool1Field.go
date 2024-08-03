package field

// ---- auto generated by builder.go, NOT modify this file ----

import (
	fmt "fmt"
	filter "github.com/findyu001/go-db-mongo/mongodb/filter"
	updater "github.com/findyu001/go-db-mongo/mongodb/updater"
	bson "go.mongodb.org/mongo-driver/bson"
)

type Bool1Field struct {
	*Array
}

func NewBool1Field(fName string) *Bool1Field {
	return &Bool1Field{NewArray(fName)}
}

func (i *Bool1Field) EleAt(index int) *Bool0F {
	return NewBool0F(fmt.Sprintf("%s.%d", i.FullName(), index))
}

// 数组的某一个元素，有时也可以理解为 数组的任何一个元素
func (i *Bool1Field) EleOne() *Bool0F {
	return NewBool0F(i.FullName())
}

// update 操作中被filter匹配的那第一个元素
func (i *Bool1Field) EleThat() *Bool0FUpdaterF {
	return NewBool0F(i.FullName() + ".$").Bool0FUpdaterF
}

func (i *Bool1Field) EleAll() *Bool0FUpdaterF {
	return NewBool0F(i.FullName() + ".$[]").Bool0FUpdaterF
}

func (i *Bool1Field) EleByFid(identifier string) *Bool0FUpdaterF {
	return NewBool0F(fmt.Sprintf("%s.$[%s]", i.FullName(), i.FullName()+identifier)).Bool0FUpdaterF
}

func (i *Bool1Field) DeclFid(identifier string) *Bool0FFilterF {
	return NewBool0F(i.FullName() + identifier).Bool0FFilterF
}

func (i *Bool1Field) Include(a []bool) filter.Filter {
	return filter.New(i, "$all", a)
}

func (i *Bool1Field) Eq(a []bool) filter.Filter {
	return filter.CompareByValue(i, filter.EQ, a)
}

func (i *Bool1Field) Set(a []bool) updater.Updater {
	return updater.New(i, "$set", a)
}

func (i *Bool1Field) AddToSet(value bool) updater.Updater {
	return updater.New(i, "$addToSet", value)
}

func (i *Bool1Field) AddToSetValues(a []bool) updater.Updater {
	return updater.New(i, "$addToSet", bson.M{"$each": a})
}

func (i *Bool1Field) Pull(value bool) updater.Updater {
	return updater.New(i, "$pull", value)
}

func (i *Bool1Field) PullAll(a []bool) updater.Updater {
	return updater.New(i, "$pullAll", a)
}

func (i *Bool1Field) Push(value bool) updater.Updater {
	return updater.New(i, "$push", value)
}

func (i *Bool1Field) PushByModifier(m updater.PushModifier, each []bool) updater.Updater {
	return updater.PushByModifier(i, m, each)
}
