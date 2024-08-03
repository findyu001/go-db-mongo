package demo

import "github.com/findyu001/go-db-mongo/mongodb/geojson"

type Wx struct {
	Age  geojson.Point
	Time *int
}

type UserInfo struct {
	Login int
	Pass  []int
	Wx    Wx
	Ws    []Wx
	Pass2 []int16
}
