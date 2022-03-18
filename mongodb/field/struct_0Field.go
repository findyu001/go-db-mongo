package field

import "strings"

func StructNext(self, fName string) string {
  if self == "" {
    return fName
  }
  if fName == "" {
    return self
  }

  b := strings.Builder{}
  b.WriteString(self)
  b.WriteByte('.')
  b.WriteString(fName)

  return b.String()
}

type StructUpdaterF struct {
  *baseUpdater
}

func NewStructUpdaterF(name string) *StructUpdaterF {
  return &StructUpdaterF{&baseUpdater{&base{name:name}}}
}

type StructFilterF struct {
  *baseFilter
}

func NewStructFilterF(name string) *StructFilterF {
  return &StructFilterF{&baseFilter{&base{name:name}}}
}
