package geojson

type MultiPoint struct {
  Type string      `bson:"type"`
  C    [][]float64 `bson:"coordinates"`
}

func (m *MultiPoint) Coordinates() []Coordinate {
  ret := make([]Coordinate, len(m.C))
  for i, c := range m.C {
    ret[i] = fromSlice(c)
  }

  return ret
}

func NewMPoint(c1 Coordinate, cs ...Coordinate) *MultiPoint {
  ret := &MultiPoint{
    Type: "MultiPoint",
    C:    make([][]float64, len(cs)+1),
  }

  ret.C[0] = c1.ToSlice()
  for i, c := range cs {
    ret.C[i+1] = c.ToSlice()
  }

  return ret
}
