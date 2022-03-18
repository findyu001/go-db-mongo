package geojson

type LineString struct {
  Type string      `bson:"type"`
  C    [][]float64 `bson:"coordinates"`
}

func NewLineString(start, end Coordinate) *LineString {
  return newBySlice([][]float64{start.ToSlice(), end.ToSlice()})
}

func newBySlice(c [][]float64) *LineString {
  return &LineString{
    Type: "LineString",
    C:    c,
  }
}

func (l *LineString) Start() Coordinate {
  return fromSlice(l.C[0])
}

func (l *LineString) End() Coordinate {
  return fromSlice(l.C[1])
}
