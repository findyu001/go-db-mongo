package demo

import (
	"fmt"
	"github.com/findyu001/go-db-mongo/mongodb/field"
	"reflect"
)

func ExampleBuilder() {

	builder := field.New()
	builder.Build(reflect.TypeOf(UserInfo{}))

	fmt.Println(true)
	// Output:
	// true
}
