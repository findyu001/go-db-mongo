package geojson

type MultiLineString struct {
  Type string        `bson:"type"`
  C    [][][]float64 `bson:"coordinates"`
}

func NewMLineString(lines []*LineString) *MultiLineString {
  ret := &MultiLineString{
    Type: "MultiLineString",
    C:    make([][][]float64, len(lines)),
  }

  for i, l := range lines {
    ret.C[i] = l.C
  }

  return ret
}

func (m *MultiLineString) LineStrings() []*LineString {
  ret := make([]*LineString, len(m.C))

  for i, ml := range m.C {
    ret[i] = newBySlice(ml)
  }

  return ret
}
