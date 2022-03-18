package field

import "github.com/xpwu/go-db-mongo/mongodb/geojson"

type prop struct {
  Name string `bson:"name"`
}

type crs struct {
  Type       string `bson:"type"`
  Properties prop   `bson:"properties"`
}
type geom struct {
  Poly *geojson.Polygon `bson:",inline"`
  CRS  crs              `bson:"crs"`
}

func NewCcwCrs(poly *geojson.Polygon) *geom {
  return &geom{
    Poly: poly,
    CRS: crs{
      Type:       "name",
      Properties: prop{Name: "urn:x-mongodb:crs:strictwinding:EPSG:4326"},
    },
  }
}
