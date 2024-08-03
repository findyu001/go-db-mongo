
package field

// ---- auto generated by builder.go, NOT modify this file ----

import (
  fmt "fmt"
  bin "github.com/gotd/td/bin"
  filter "github.com/xpwu/go-db-mongo/mongodb/filter"
  updater "github.com/xpwu/go-db-mongo/mongodb/updater"
  bson "go.mongodb.org/mongo-driver/bson"
)

type Fields1Field struct {
  *Array
}

func NewFields1Field(fName string) *Fields1Field {
  return &Fields1Field { NewArray(fName)}
}

func (i *Fields1Field) EleAt(index int) *Fields0F {
  return NewFields0F(fmt.Sprintf("%s.%d", i.FullName(), index))
}

// 数组的某一个元素，有时也可以理解为 数组的任何一个元素
func (i *Fields1Field) EleOne() *Fields0F {
  return NewFields0F(i.FullName())
}

// update 操作中被filter匹配的那第一个元素
func (i *Fields1Field) EleThat() *Fields0FUpdaterF {
  return NewFields0F(i.FullName() + ".$").Fields0FUpdaterF
}

func (i *Fields1Field) EleAll() *Fields0FUpdaterF {
  return NewFields0F(i.FullName() + ".$[]").Fields0FUpdaterF
}

func (i *Fields1Field) EleByFid(identifier string) *Fields0FUpdaterF {
  return NewFields0F(fmt.Sprintf("%s.$[%s]", i.FullName(), i.FullName()+identifier)).Fields0FUpdaterF
}

func (i *Fields1Field) DeclFid(identifier string) *Fields0FFilterF {
  return NewFields0F(i.FullName()+identifier).Fields0FFilterF
}

func (i *Fields1Field) Include(a []bin.Fields) filter.Filter {
  return filter.New(i, "$all", a)
}

func (i *Fields1Field) Eq(a []bin.Fields) filter.Filter {
  return filter.CompareByValue(i, filter.EQ, a)
}

func (i *Fields1Field) Set(a []bin.Fields) updater.Updater {
  return updater.New(i, "$set", a)
}

func (i *Fields1Field) AddToSet(value bin.Fields) updater.Updater {
  return updater.New(i, "$addToSet", value)
}

func (i *Fields1Field) AddToSetValues(a []bin.Fields) updater.Updater {
  return updater.New(i, "$addToSet", bson.M{"$each":a})
}

func (i *Fields1Field) Pull(value bin.Fields) updater.Updater {
  return updater.New(i, "$pull", value)
}

func (i *Fields1Field) PullAll(a []bin.Fields) updater.Updater {
  return updater.New(i, "$pullAll", a)
}

func (i *Fields1Field) Push (value bin.Fields) updater.Updater {
  return updater.New(i, "$push", value)
}

func (i *Fields1Field) PushByModifier(m updater.PushModifier, each []bin.Fields) updater.Updater {
  return updater.PushByModifier(i, m, each)
}