package geojson

//func TestLineString(t *testing.T) {
//  type js struct {
//    Type string `bson:"type"`
//    C [][]float64 `bson:"coordinates"`
//  }
//
//  cases := []struct {
//    ls *LineString
//    js *js
//  }{
//    {
//      NewLineString(C(0, 0), C(1, 1,)),
//      &js{
//        Type: "",
//        C:    nil,
//      },
//    },
//  }
//
//  for _,cs := range cases {
//    b, err := bson.Marshal(cs.ls)
//    if err != nil {
//      t.Error(err)
//    }
//    js := &js{}
//
//    err = bson.Unmarshal(b, js)
//    if err != nil {
//      t.Error(err)
//    }
//
//    fmt.Println(js)
//  }
//}
