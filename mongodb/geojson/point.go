package geojson

type Point struct {
  Type string `bson:"type"`
  C []float64 `bson:"coordinates"`
}

func (p *Point) Coordinate() Coordinate {
  return Coordinate{
    X: p.C[0],
    Y: p.C[1],
  }
}

func NewPoint(c *Coordinate) *Point {
  return &Point{
    Type: `Point`,
    C:    []float64{c.X, c.Y},
  }
}
