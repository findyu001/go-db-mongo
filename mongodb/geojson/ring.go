package geojson

type Ring struct {
  C []Coordinate
}

// 第一点必须与最后一点一样
func R(c1, c2, c3, c4 Coordinate, cn ...Coordinate) *Ring {
  ret := Ring{}
  ret.C = append(ret.C, c1, c2, c3, c4)
  ret.C = append(ret.C, cn...)
  return &ret
}

func (r *Ring) toSlice() [][]float64 {
  ret := make([][]float64, len(r.C))
  for i, c := range r.C {
    ret[i] = c.ToSlice()
  }

  return ret
}

func rFromSlice(ss [][]float64) *Ring {
  ret := &Ring{C:make([]Coordinate, len(ss))}

  for i,s := range ss {
    ret.C[i] = fromSlice(s)
  }

  return ret
}
