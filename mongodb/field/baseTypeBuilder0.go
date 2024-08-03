package field

import (
	"os"
	"text/template"
)

func fieldCode() *template.Template {
	return template.Must(template.New("baseCode").Parse(`
package field

// ---- auto generated by baseTypeBuilder_test.go, NOT modify this file ----

import (
  "github.com/xpwu/go-db-mongo/mongodb/filter"
  "github.com/xpwu/go-db-mongo/mongodb/updater"
  {{if .ModAble}}"go.mongodb.org/mongo-driver/bson"{{end}}
  {{if or .RegexAble .BsonPType}}"go.mongodb.org/mongo-driver/bson/primitive"{{end}}
)

type {{.FType}}0F struct {
  *{{.FType}}0FUpdaterF
  *baseKey
  *{{.FType}}0FFilterF
}

func (a *{{.FType}}0F) FullName() string {
  return a.baseKey.FullName()
}

type {{.FType}}0FUpdaterF struct {
  *baseUpdater
}

type {{.FType}}0FFilterF struct {
  *baseFilter
}

func New{{.FType}}0F(fName string) *{{.FType}}0F {
  uper := &{{.FType}}0FUpdaterF{&baseUpdater{&base{fName}}}
  pri := &baseKey{uper.base}
  flt := &{{.FType}}0FFilterF{&baseFilter{uper.base}}

  return &{{.FType}}0F{uper, pri, flt}
}

{{if .ModAble}}
func (i *{{.FType}}0FFilterF) Mod(divisor, remainder {{.Type}}) filter.Filter {
 return filter.New(i, "$mod", bson.A{divisor, remainder})
}
{{end}}

{{if .ComputeAble}}
func (i *{{.FType}}0FUpdaterF) Inc(num {{.Type}}) updater.Updater {
 return updater.New(i, "$inc", num)
}

func (i *{{.FType}}0FUpdaterF) Mul(num {{.Type}}) updater.Updater {
 return updater.New(i, "$mul", num)
}
{{end}}

{{if .EqAble}}
func (i *{{.FType}}0FFilterF) Eq(value {{.Type}}) filter.Filter {
 return filter.CompareByValue(i, filter.EQ, value)
}

func (i *{{.FType}}0FFilterF) Ne(value {{.Type}}) filter.Filter {
 return filter.CompareByValue(i, filter.NE, value)
}

func (i *{{.FType}}0FFilterF) NeField(f *{{.FType}}0F) filter.Filter {
 return filter.CompareByField(i, filter.NE, f)
}

func (i *{{.FType}}0FFilterF) EqField(f *{{.FType}}0F) filter.Filter {
 return filter.CompareByField(i, filter.EQ, f)
}

func (i *{{.FType}}0FFilterF) Gte(value {{.Type}}) filter.Filter {
 return filter.CompareByValue(i, filter.GTE, value)
}

func (i *{{.FType}}0FFilterF) Lte(value {{.Type}}) filter.Filter {
 return filter.CompareByValue(i, filter.LTE, value)
}

func (i *{{.FType}}0FFilterF) GteField(f *{{.FType}}0F) filter.Filter {
 return filter.CompareByValue(i, filter.GTE, f)
}

func (i *{{.FType}}0FFilterF) LteField(f *{{.FType}}0F) filter.Filter {
 return filter.CompareByValue(i, filter.LTE, f)
}
{{end}}

{{if .RegexAble}}
func (i *{{.FType}}0FUpdaterF) Regex(regex primitive.Regex) filter.Filter {
  return filter.New(i, "$regex", regex)
}
{{end}}
{{if .CompareAble}}
func (i *{{.FType}}0FFilterF) Gt(value {{.Type}}) filter.Filter {
 return filter.CompareByValue(i, filter.GT, value)
}

func (i *{{.FType}}0FFilterF) Lt(value {{.Type}}) filter.Filter {
 return filter.CompareByValue(i, filter.LT, value)
}

func (i *{{.FType}}0FFilterF) GtField(f *{{.FType}}0F) filter.Filter {
 return filter.CompareByValue(i, filter.GT, f)
}

func (i *{{.FType}}0FFilterF) LtField(f *{{.FType}}0F) filter.Filter {
 return filter.CompareByValue(i, filter.LT, f)
}
{{end}}
func (i *{{.FType}}0FFilterF) In(values []{{.Type}}) filter.Filter {
 return filter.New(i, "$in", values)
}

func (i *{{.FType}}0FFilterF) Nin(values []{{.Type}}) filter.Filter {
 return filter.New(i, "$nin", values)
}
{{if .CompareAble}}
func (i *{{.FType}}0FUpdaterF) Min(value {{.Type}}) updater.Updater {
  return updater.New(i, "$min", value)
}

func (i *{{.FType}}0FUpdaterF) Max(value {{.Type}}) updater.Updater {
  return updater.New(i, "$max", value)
}
{{end}}
func (i *{{.FType}}0FUpdaterF) Set(value {{.Type}}) updater.Updater {
  return updater.New(i, "$set", value)
}

func (i *{{.FType}}0FUpdaterF) SetOnIns(value {{.Type}}) updater.Updater {
  return updater.New(i, "$setOnInsert", value)
}

`))

}

func builderR0Code() *template.Template {
	return template.Must(template.New("baseCode").Parse(`
package field

// ---- auto generated by baseTypeBuilder_test.go, NOT modify this file ----

import (
  {{if .BsonPType}}"go.mongodb.org/mongo-driver/bson/primitive"{{end}}
  "reflect"
)

func init() {
  builderR0 = func(b *Builder) {
    {{range .Items }}
    b.RegisterType(reflect.TypeOf((*{{.Type}})(nil)).Elem(), Type{
      F:    reflect.TypeOf((*{{.FType}}0F)(nil)).Elem(),
      NewF: NewTypByFunc(New{{.FType}}0F),
    })
    {{end}}
  }
}

`))
}

type typ struct {
	FType, Type                                                     string
	ComputeAble, EqAble, RegexAble, ModAble, BsonPType, CompareAble bool
}

func createFile(tp *typ) {
	file, err := os.Create("z" + tp.FType + "0Field.go")
	if err != nil {
		panic(err)
	}

	err = fieldCode().Execute(file, tp)
	if err != nil {
		panic(err)
	}
}

type goPrimitiveValue struct {
	ty                                      string
	ComputeAble, EqAble, RegexAble, ModAble bool
}

func goPrimitive() []goPrimitiveValue {
	return []goPrimitiveValue{
		{
			"Fields", true, true, false, true,
		},
		{
			"int", true, true, false, true,
		},
		{
			"int8", true, true, false, true,
		},
		{
			"int16", true, true, false, true,
		},
		{
			"int32", true, true, false, true,
		},
		{
			"int64", true, true, false, true,
		},
		{
			"uint", true, true, false, true,
		},
		{
			"uint8", true, true, false, true,
		},
		{
			"uint16", true, true, false, true,
		},
		{
			"uint32", true, true, false, true,
		},
		{
			"uint64", true, true, false, true,
		},
		{
			"string", false, true, true, false,
		},
		{
			"float32", false, false, false, false,
		},
		{
			"float64", false, false, false, false,
		},
		{
			"bool", false, true, false, false,
		},
	}
}

type bsonPrimitiveValue struct {
	ty                                                   string
	ComputeAble, EqAble, RegexAble, ModAble, CompareAble bool
}

func bsonPrimitive() []bsonPrimitiveValue {
	return []bsonPrimitiveValue{
		{
			"Binary", false, false, false, false, false,
		},
		{
			"Decimal128", true, false, false, false, true,
		},
		{
			"ObjectID", false, true, false, false, true,
		},
	}
}

func buildAll0Field() {

	for _, t := range goPrimitive() {
		tp := &typ{
			FType:       firstUpper(t.ty),
			Type:        t.ty,
			ComputeAble: t.ComputeAble,
			EqAble:      t.EqAble,
			RegexAble:   t.RegexAble,
			ModAble:     t.ModAble,
			CompareAble: true,
		}

		createFile(tp)
	}

	for _, t := range bsonPrimitive() {
		tp := &typ{
			FType:       t.ty,
			Type:        "primitive." + t.ty,
			ComputeAble: t.ComputeAble,
			EqAble:      t.EqAble,
			RegexAble:   t.RegexAble,
			ModAble:     t.ModAble,
			BsonPType:   true,
			CompareAble: t.CompareAble,
		}

		createFile(tp)
	}

}

func build0Builder() {
	d := struct {
		BsonPType bool
		Items     []typ
	}{}

	bsonTs := bsonPrimitive()
	d.BsonPType = len(bsonTs) != 0

	goTs := goPrimitive()

	d.Items = make([]typ, 0, len(goTs)+len(bsonTs))

	for _, t := range goTs {
		d.Items = append(d.Items, typ{
			FType: firstUpper(t.ty),
			Type:  t.ty,
		})
	}

	for _, t := range bsonTs {
		d.Items = append(d.Items, typ{
			FType: t.ty,
			Type:  "primitive." + t.ty,
		})
	}

	file, err := os.Create("zBuilderR0.go")
	if err != nil {
		panic(err)
	}

	err = builderR0Code().Execute(file, d)
	if err != nil {
		panic(err)
	}
}

// 生成全部0Field及相关的其他功能，此函数的执行借助了go提供的 Example 功能，具体见 "baseTypeBuilder_run_test.go"
func buildAll0() {
	buildAll0Field()
	build0Builder()
}
