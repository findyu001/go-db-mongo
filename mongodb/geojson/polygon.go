package geojson

type Polygon struct {
  Type string        `bson:"type"`
  C    [][][]float64 `bson:"coordinates"`
}

func NewPolygon(r1 Ring, rs ...Ring) *Polygon {
  c := make([][][]float64, len(rs)+1)
  c[0] = r1.toSlice()
  for i, r := range rs {
    c[i+1] = r.toSlice()
  }

  return newPygBySlice(c)
}

func newPygBySlice(ss [][][]float64) *Polygon {
  return &Polygon{
    Type: "Polygon",
    C:    ss,
  }
}

func (p *Polygon) Rings() []*Ring {
  ret := make([]*Ring, len(p.C))

  for i, c := range p.C {
    ret[i] = rFromSlice(c)
  }

  return ret
}
