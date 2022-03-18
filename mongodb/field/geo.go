package field

type Geo struct {
  *base
  *baseUpdater
  *baseFilter
}

func NewGeo(name string) *Geo {
  ret := &Geo{
    base: &base{name: name},
  }
  ret.baseFilter = &baseFilter{ret.base}
  ret.baseUpdater = &baseUpdater{ret.base}

  return ret
}

func (g *Geo) FullName() string {
  return g.base.FullName()
}
