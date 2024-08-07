package field

// ---- auto generated by builder.go, NOT modify this file ----

import (
	fmt "fmt"
	filter "github.com/findyu001/go-db-mongo/mongodb/filter"
	updater "github.com/findyu001/go-db-mongo/mongodb/updater"
	bson "go.mongodb.org/mongo-driver/bson"
)

type Float321Field struct {
	*Array
}

func NewFloat321Field(fName string) *Float321Field {
	return &Float321Field{NewArray(fName)}
}

func (i *Float321Field) EleAt(index int) *Float320F {
	return NewFloat320F(fmt.Sprintf("%s.%d", i.FullName(), index))
}

// 数组的某一个元素，有时也可以理解为 数组的任何一个元素
func (i *Float321Field) EleOne() *Float320F {
	return NewFloat320F(i.FullName())
}

// update 操作中被filter匹配的那第一个元素
func (i *Float321Field) EleThat() *Float320FUpdaterF {
	return NewFloat320F(i.FullName() + ".$").Float320FUpdaterF
}

func (i *Float321Field) EleAll() *Float320FUpdaterF {
	return NewFloat320F(i.FullName() + ".$[]").Float320FUpdaterF
}

func (i *Float321Field) EleByFid(identifier string) *Float320FUpdaterF {
	return NewFloat320F(fmt.Sprintf("%s.$[%s]", i.FullName(), i.FullName()+identifier)).Float320FUpdaterF
}

func (i *Float321Field) DeclFid(identifier string) *Float320FFilterF {
	return NewFloat320F(i.FullName() + identifier).Float320FFilterF
}

func (i *Float321Field) Include(a []float32) filter.Filter {
	return filter.New(i, "$all", a)
}

func (i *Float321Field) Eq(a []float32) filter.Filter {
	return filter.CompareByValue(i, filter.EQ, a)
}

func (i *Float321Field) Set(a []float32) updater.Updater {
	return updater.New(i, "$set", a)
}

func (i *Float321Field) AddToSet(value float32) updater.Updater {
	return updater.New(i, "$addToSet", value)
}

func (i *Float321Field) AddToSetValues(a []float32) updater.Updater {
	return updater.New(i, "$addToSet", bson.M{"$each": a})
}

func (i *Float321Field) Pull(value float32) updater.Updater {
	return updater.New(i, "$pull", value)
}

func (i *Float321Field) PullAll(a []float32) updater.Updater {
	return updater.New(i, "$pullAll", a)
}

func (i *Float321Field) Push(value float32) updater.Updater {
	return updater.New(i, "$push", value)
}

func (i *Float321Field) PushByModifier(m updater.PushModifier, each []float32) updater.Updater {
	return updater.PushByModifier(i, m, each)
}
