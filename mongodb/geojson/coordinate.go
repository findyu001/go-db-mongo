package geojson

import "fmt"

type Coordinate struct {
  X float64
  Y float64
}

func LngLat(lng float64, lat float64) Coordinate {
  if lng > 180 || lng < -180 || lat > 90 || lat < -90 {
    panic(fmt.Sprintf("lng(%f) lat(%f) error", lng, lat))
  }

  return Coordinate{
    X: lng,
    Y: lat,
  }
}

func C(x, y float64) Coordinate {
  return Coordinate{
    X: x,
    Y: y,
  }
}

func fromSlice(s []float64) Coordinate {
  return Coordinate{
    X: s[0],
    Y: s[1],
  }
}

func (c *Coordinate) ToSlice() []float64 {
  return []float64{c.X, c.Y}
}
