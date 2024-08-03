package field

import (
	"github.com/findyu001/go-db-mongo/mongodb/geojson"
	"github.com/findyu001/go-db-mongo/mongodb/updater"
)

type MultiPointField struct {
	*Geo
}

func NewMultiPointField(name string) *MultiPointField {
	return &MultiPointField{NewGeo(name)}
}

func (p *MultiPointField) Set(value *geojson.MultiPoint) updater.Updater {
	return updater.New(p, `$set`, value)
}

func (p *MultiPointField) SetOnInsert(value *geojson.MultiPoint) updater.Updater {
	return updater.New(p, `$setOnInsert`, value)
}

func (p *MultiPointField) Sphere() *PointSphere {
	return NewSphere(p.FullName())
}

func (p *MultiPointField) Flat() *PointFlat {
	return NewFlat(p.FullName())
}
