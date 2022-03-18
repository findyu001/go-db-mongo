package geojson

type MultiPolygon struct {
  Type string          `bson:"type"`
  C    [][][][]float64 `bson:"coordinates"`
}

func NewMPg(pg1 Polygon, pgs ...Polygon) *MultiPolygon {
  ret := &MultiPolygon{
    Type: "MultiPolygon",
    C:    make([][][][]float64, len(pgs)+1),
  }
  ret.C[0] = pg1.C
  for i, pg := range pgs {
    ret.C[i+1] = pg.C
  }

  return ret
}

func (m *MultiPolygon) Polygons() []*Polygon {
  ret := make([]*Polygon, len(m.C))

  for i,c := range m.C {
    ret[i] = newPygBySlice(c)
  }

  return ret
}
