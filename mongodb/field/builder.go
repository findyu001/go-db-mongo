package field

import (
	"fmt"
	"github.com/findyu001/go-db-mongo/mongodb/filter"
	"github.com/findyu001/go-db-mongo/mongodb/geojson"
	"github.com/findyu001/go-db-mongo/mongodb/tagparser"
	"github.com/findyu001/go-db-mongo/mongodb/updater"
	"os"
	"path"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"text/template"
)

type Typ interface {
	Name() string
	PkgPath() string
}

type rTyp struct {
	name string
	pkg  string
}

func (r *rTyp) Name() string {
	return r.name
}

func (r *rTyp) PkgPath() string {
	return r.pkg
}

func NewTypByFunc(fun interface{}) Typ {
	name := runtime.FuncForPC(reflect.ValueOf(fun).Pointer()).Name()
	f := strings.FieldsFunc(name, func(r rune) bool {
		if r == '.' {
			return true
		}
		return false
	})

	return &rTyp{pkg: strings.Join(f[:len(f)-1], "."), name: f[len(f)-1]}
}

// 名字要求：如果存在UpdaterF与FilterF两个类型，则是在本身类型后面加后缀，变量命名同理
type Type struct {
	F    Typ
	NewF Typ
}

type Builder struct {
	typeMap map[reflect.Type]Type
	kindMap map[reflect.Kind]func(reflect.Type) (Type, bool)
	dir     string
	pkg     string
}

var (
	builderR0 = func(b *Builder) {
	}

	builderR1 = func(b *Builder) {
	}
)

func New() *Builder {
	b := &Builder{
		typeMap: make(map[reflect.Type]Type),
		kindMap: make(map[reflect.Kind]func(reflect.Type) (Type, bool)),
	}

	b.RegisterDefault(reflect.Struct, b.buildStruct).
		RegisterDefault(reflect.Slice, b.buildSlice).
		RegisterDefault(reflect.Ptr, b.buildPtr)

	b.RegisterType(reflect.TypeOf(geojson.Point{}), Type{
		F:    reflect.TypeOf(PointField{}),
		NewF: NewTypByFunc(NewPointField),
	}).
		RegisterType(reflect.TypeOf(geojson.MultiPoint{}), Type{
			F:    reflect.TypeOf(MultiPointField{}),
			NewF: NewTypByFunc(NewMultiPointField),
		}).
		RegisterType(reflect.TypeOf(geojson.Polygon{}), Type{
			F:    reflect.TypeOf(PolygonField{}),
			NewF: NewTypByFunc(NewPointField),
		}).
		RegisterType(reflect.TypeOf(geojson.MultiPolygon{}), Type{
			F:    reflect.TypeOf(MultiPolygonField{}),
			NewF: NewTypByFunc(NewMultiPolygonField),
		})

	builderR0(b)
	builderR1(b)

	return b
}

func (b *Builder) ClearType(rt reflect.Type) *Builder {
	delete(b.typeMap, rt)
	return b
}

func (b *Builder) RegisterType(rt reflect.Type, ft Type) *Builder {
	b.typeMap[rt] = ft
	return b
}

func (b *Builder) RegisterDefault(k reflect.Kind, f func(rt reflect.Type) (Type, bool)) *Builder {
	b.kindMap[k] = f
	return b
}

func (b *Builder) build(rt reflect.Type) (ft Type, ok bool) {
	ft, ok = b.typeMap[rt]
	if ok {
		return
	}

	f := b.kindMap[rt.Kind()]
	ft, ok = f(rt)
	if ok {
		b.typeMap[rt] = ft
	}

	return
}

func (b *Builder) Build(rt reflect.Type) {

	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	b.pkg = rt.PkgPath()
	if b.pkg == "" {
		// 基本类型 就取当前的pkg, 基本类型都是预生成在此pkg中
		// 对于其他pkg为空的类型 暂不支持
		b.pkg = reflect.TypeOf(b).Elem().PkgPath()
	}

	b.build(rt)
}

func (b *Builder) buildPtr(t reflect.Type) (ft Type, ok bool) {
	return b.build(t.Elem())
}

type alias map[string]bool

func (a *alias) get(expect string) string {
	test := expect
	num := 1
	for (*a)[test] {
		num++
		test = fmt.Sprintf("%s%d", expect, num)
	}
	(*a)[test] = true

	return test
}

type imports struct {
	data  map[string]string
	alias alias
	exc   map[string]bool
}

func newImports() *imports {
	return &imports{
		data:  make(map[string]string),
		alias: alias{},
		exc:   make(map[string]bool),
	}
}

func (m *imports) exclude(paths string) {
	m.exc[paths] = true
}

func (m *imports) add(paths string) (alias string) {
	if paths == "" || m.exc[paths] {
		return ""
	}

	if a, ok := m.data[paths]; ok {
		return a + "."
	}

	a := m.alias.get(path.Base(paths))
	m.data[paths] = a

	return a + "."
}

type importT struct {
	Alias  string
	Import string
}

func (m *imports) all() []importT {
	ret := make([]importT, len(m.data))

	var keys []string
	for k := range m.data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for j, k := range keys {
		ret[j] = importT{
			Alias:  m.data[k],
			Import: k,
		}
	}

	return ret
}

var sliceCode = template.Must(template.New("sliceCode").Parse(`
package {{.Pkg}}

// ---- auto generated by builder.go, NOT modify this file ----

import ({{range .Imports}}
  {{.Alias}} "{{.Import}}"{{end}}
)

type {{.Name}} struct {
  *{{.MongoFieldAlias}}Array
}

func New{{.Name}}(fName string) *{{.Name}} {
  return &{{.Name}} { {{.MongoFieldAlias}}NewArray(fName)}
}

func (i *{{.Name}}) EleAt(index int) *{{.EleName}} {
  return {{.NewEleName}}({{.FmtAlias}}Sprintf("%s.%d", i.FullName(), index))
}

// 数组的某一个元素，有时也可以理解为 数组的任何一个元素
func (i *{{.Name}}) EleOne() *{{.EleName}} {
  return {{.NewEleName}}(i.FullName())
}

// update 操作中被filter匹配的那第一个元素
func (i *{{.Name}}) EleThat() *{{.EleName}}UpdaterF {
  return {{.NewEleName}}(i.FullName() + ".$").{{.EleOnlyName}}UpdaterF
}

func (i *{{.Name}}) EleAll() *{{.EleName}}UpdaterF {
  return {{.NewEleName}}(i.FullName() + ".$[]").{{.EleOnlyName}}UpdaterF
}

func (i *{{.Name}}) EleByFid(identifier string) *{{.EleName}}UpdaterF {
  return {{.NewEleName}}({{.FmtAlias}}Sprintf("%s.$[%s]", i.FullName(), i.FullName()+identifier)).{{.EleOnlyName}}UpdaterF
}

func (i *{{.Name}}) DeclFid(identifier string) *{{.EleName}}FilterF {
  return {{.NewEleName}}(i.FullName()+identifier).{{.EleOnlyName}}FilterF
}

func (i *{{.Name}}) Include(a []{{.EleNameType}}) {{.FilterAlias}}Filter {
  return {{.FilterAlias}}New(i, "$all", a)
}

func (i *{{.Name}}) Eq(a []{{.EleNameType}}) {{.FilterAlias}}Filter {
  return {{.FilterAlias}}CompareByValue(i, {{.FilterAlias}}EQ, a)
}

func (i *{{.Name}}) Set(a []{{.EleNameType}}) {{.UpdaterAlias}}Updater {
  return {{.UpdaterAlias}}New(i, "$set", a)
}

func (i *{{.Name}}) AddToSet(value {{.EleNameType}}) {{.UpdaterAlias}}Updater {
  return {{.UpdaterAlias}}New(i, "$addToSet", value)
}

func (i *{{.Name}}) AddToSetValues(a []{{.EleNameType}}) {{.UpdaterAlias}}Updater {
  return {{.UpdaterAlias}}New(i, "$addToSet", {{.BsonAlias}}M{"$each":a})
}

func (i *{{.Name}}) Pull(value {{.EleNameType}}) {{.UpdaterAlias}}Updater {
  return {{.UpdaterAlias}}New(i, "$pull", value)
}

func (i *{{.Name}}) PullAll(a []{{.EleNameType}}) {{.UpdaterAlias}}Updater {
  return {{.UpdaterAlias}}New(i, "$pullAll", a)
}

func (i *{{.Name}}) Push (value {{.EleNameType}}) {{.UpdaterAlias}}Updater {
  return {{.UpdaterAlias}}New(i, "$push", value)
}

func (i *{{.Name}}) PushByModifier(m {{.UpdaterAlias}}PushModifier, each []{{.EleNameType}}) {{.UpdaterAlias}}Updater {
  return {{.UpdaterAlias}}PushByModifier(i, m, each)
}
`))

func firstUpper(i string) string {
	if i == "" {
		return i
	}

	return strings.ToUpper(i[0:1]) + i[1:]
}

func (b *Builder) buildSlice(t reflect.Type) (ft Type, ok bool) {
	ele := t.Elem()
	if ele.Kind() == reflect.Slice {
		panic(fmt.Errorf("not support [][] type"))
	}

	type st struct {
		Pkg             string
		Name            string
		MongoFieldAlias string
		Imports         []importT
		EleName         string
		EleOnlyName     string
		EleNameType     string
		NewEleName      string
		UpdaterAlias    string
		FilterAlias     string
		FmtAlias        string
		BsonAlias       string
	}

	ft, ok = b.build(ele)
	if !ok {
		panic(fmt.Errorf("not support %v", ele))
	}

	imports := newImports()
	imports.exclude(b.pkg)

	s := &st{
		Pkg:             path.Base(b.pkg),
		Name:            firstUpper(ele.Name()) + "1Field",
		MongoFieldAlias: imports.add(reflect.TypeOf(Array{}).PkgPath()),
		EleName:         imports.add(ft.F.PkgPath()) + ft.F.Name(),
		EleOnlyName:     ft.F.Name(),
		EleNameType:     imports.add(ele.PkgPath()) + ele.Name(),
		NewEleName:      imports.add(ft.NewF.PkgPath()) + ft.NewF.Name(),
		UpdaterAlias:    imports.add(reflect.TypeOf((*updater.Updater)(nil)).Elem().PkgPath()),
		FilterAlias:     imports.add(reflect.TypeOf((*filter.Filter)(nil)).Elem().PkgPath()),
		FmtAlias:        imports.add("fmt"),
		BsonAlias:       imports.add("go.mongodb.org/mongo-driver/bson"),
	}
	s.Imports = imports.all()

	file, err := os.Create("z" + s.Name + ".go")
	if err != nil {
		panic(err)
	}

	err = sliceCode.Execute(file, s)
	if err != nil {
		panic(err)
	}

	ft = Type{}
	ft.F = &rTyp{s.Name, b.pkg}
	ft.NewF = &rTyp{"New" + s.Name, b.pkg}

	return ft, true
}

var structCode = template.Must(template.New("structCode").Parse(`
package {{.Pkg}}

// ---- auto generated by builder.go, NOT modify this file ----

import ({{range .Imports}}
  {{.Alias}} "{{.Import}}"
{{end}})

type {{.Name}}UpdaterF struct {
  *base{{.Name}}
  *{{.MongoFieldAlias}}StructUpdaterF
}

func (s *{{.Name}}UpdaterF) FullName() string {
  return s.name
}

type {{.Name}}FilterF struct {
  *base{{.Name}}
  *{{.MongoFieldAlias}}StructFilterF
}

func (s *{{.Name}}FilterF) FullName() string {
  return s.name
}

type {{.Name}} struct {
  *base{{.Name}}
  {{.Name}}UpdaterF  *{{.Name}}UpdaterF
  {{.Name}}FilterF   *{{.Name}}FilterF
}

func New{{.Name}}(fName string) *{{.Name}} {
  base := &base{{.Name}}{fName}
  // 没有name时，不能做updater与filter操作，比如最顶层的Struct
  if fName == "" {
    return &{{.Name}} {base{{.Name}}:base}
  }
  up := &{{.Name}}UpdaterF{base, {{.MongoFieldAlias}}NewStructUpdaterF(fName)}
  fl := &{{.Name}}FilterF{base, {{.MongoFieldAlias}}NewStructFilterF(fName)}

  return &{{.Name}} {base, up, fl}
}

// 对应于 bson struct 中的 inline 修饰符
func New{{.Name}}Inline(fName string) *{{.Name}} {
  return &{{.Name}} {base{{.Name}}: &base{{.Name}}{fName}}
}

func (s *{{.Name}}) FullName() string {
  return s.name
}

type base{{.Name}} struct {
  name      string
}
{{range .Fields}}
func (s *base{{$.Name}}) {{.MethodName}}() *{{.FieldName}} {
  n := {{$.NextNameM}}(s.name, "{{.TagName}}")
  return {{.New}}(n)
}
{{end}}
`))

func (b *Builder) buildStruct(t reflect.Type) (ft Type, ok bool) {
	type Field struct {
		MethodName string
		FieldName  string
		TagName    string
		New        string
	}

	type st struct {
		Pkg             string
		Name            string
		NextNameM       string
		MongoFieldAlias string
		Imports         []importT
		Fields          []Field
	}

	imports := newImports()
	imports.exclude(b.pkg)

	nextMt := NewTypByFunc(StructNext)

	s := &st{
		Pkg:             path.Base(b.pkg),
		Name:            t.Name() + "0Field",
		NextNameM:       imports.add(nextMt.PkgPath()) + nextMt.Name(),
		MongoFieldAlias: imports.add(reflect.TypeOf(StructUpdaterF{}).PkgPath()),
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		// unexported
		if f.PkgPath != "" {
			continue
		}
		tag, _ := tagparser.StructTagParser(f)
		if tag.Skip {
			continue
		}
		fd := Field{}
		fd.MethodName = f.Name
		fd.TagName = tag.Name
		ft, ok := b.build(f.Type)
		if !ok {
			panic(fmt.Errorf("not support %v", f.Type))
		}
		fd.FieldName = imports.add(ft.F.PkgPath()) + ft.F.Name()
		fd.New = imports.add(ft.NewF.PkgPath()) + ft.NewF.Name()

		// inline 域能进行普通子域访问，但不能有任何操作，也不能在访问路径中添加新的嵌套字段
		if tag.Inline && f.Type.Kind() == reflect.Struct {
			fd.TagName = ""
			fd.New = imports.add(ft.NewF.PkgPath()) + ft.NewF.Name() + "Inline"
		}
		s.Fields = append(s.Fields, fd)
	}

	s.Imports = imports.all()

	file, err := os.Create("z" + s.Name + ".go")
	if err != nil {
		panic(err)
	}

	err = structCode.Execute(file, s)
	if err != nil {
		panic(err)
	}

	ft = Type{}
	ft.F = &rTyp{s.Name, b.pkg}
	ft.NewF = &rTyp{"New" + s.Name, b.pkg}

	return ft, true
}
