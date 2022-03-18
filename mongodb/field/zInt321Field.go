
package field

// ---- auto generated by builder.go, NOT modify this file ----

import (
  fmt "fmt"
  filter "github.com/xpwu/go-db-mongo/mongodb/filter"
  updater "github.com/xpwu/go-db-mongo/mongodb/updater"
  bson "go.mongodb.org/mongo-driver/bson"
)

type Int321Field struct {
  *Array
}

func NewInt321Field(fName string) *Int321Field {
  return &Int321Field { NewArray(fName)}
}

func (i *Int321Field) EleAt(index int) *Int320F {
  return NewInt320F(fmt.Sprintf("%s.%d", i.FullName(), index))
}

// 数组的某一个元素，有时也可以理解为 数组的任何一个元素
func (i *Int321Field) EleOne() *Int320F {
  return NewInt320F(i.FullName())
}

// update 操作中被filter匹配的那第一个元素
func (i *Int321Field) EleThat() *Int320FUpdaterF {
  return NewInt320F(i.FullName() + ".$").Int320FUpdaterF
}

func (i *Int321Field) EleAll() *Int320FUpdaterF {
  return NewInt320F(i.FullName() + ".$[]").Int320FUpdaterF
}

func (i *Int321Field) EleByFid(identifier string) *Int320FUpdaterF {
  return NewInt320F(fmt.Sprintf("%s.$[%s]", i.FullName(), i.FullName()+identifier)).Int320FUpdaterF
}

func (i *Int321Field) DeclFid(identifier string) *Int320FFilterF {
  return NewInt320F(i.FullName()+identifier).Int320FFilterF
}

func (i *Int321Field) Include(a []int32) filter.Filter {
  return filter.New(i, "$all", a)
}

func (i *Int321Field) Eq(a []int32) filter.Filter {
  return filter.CompareByValue(i, filter.EQ, a)
}

func (i *Int321Field) Set(a []int32) updater.Updater {
  return updater.New(i, "$set", a)
}

func (i *Int321Field) AddToSet(value int32) updater.Updater {
  return updater.New(i, "$addToSet", value)
}

func (i *Int321Field) AddToSetValues(a []int32) updater.Updater {
  return updater.New(i, "$addToSet", bson.M{"$each":a})
}

func (i *Int321Field) Pull(value int32) updater.Updater {
  return updater.New(i, "$pull", value)
}

func (i *Int321Field) PullAll(a []int32) updater.Updater {
  return updater.New(i, "$pullAll", a)
}

func (i *Int321Field) Push (value int32) updater.Updater {
  return updater.New(i, "$push", value)
}

func (i *Int321Field) PushByModifier(m updater.PushModifier, each []int32) updater.Updater {
  return updater.PushByModifier(i, m, each)
}