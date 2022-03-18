package geojson

type GeoCollection struct {
  Type string `bson:"type"`
  Gs []interface{} `bson:"geometries"`
}
