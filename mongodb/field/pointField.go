package field

import (
	"github.com/findyu001/go-db-mongo/mongodb/geojson"
	"github.com/findyu001/go-db-mongo/mongodb/updater"
)

type PointField struct {
	*Geo
}

func NewPointField(name string) *PointField {
	return &PointField{NewGeo(name)}
}

func (p *PointField) Set(value *geojson.Point) updater.Updater {
	return updater.New(p, `$set`, value)
}

func (p *PointField) SetOnInsert(value *geojson.Point) updater.Updater {
	return updater.New(p, `$setOnInsert`, value)
}

func (p *PointField) Sphere() *PointSphere {
	return NewSphere(p.FullName())
}

func (p *PointField) Flat() *PointFlat {
	return NewFlat(p.FullName())
}
