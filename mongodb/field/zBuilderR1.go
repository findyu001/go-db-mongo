
package field

// ---- auto generated by baseTypeBuilder_test.go, NOT modify this file ----

import (
  "reflect"
)

func init() {
  builderR1 = func(b *Builder) {
    all1 := all1Types()

    fieldTypes := []interface{} {
      (*Int1Field)(nil),
      (*Int81Field)(nil),
      (*Int161Field)(nil),
      (*Int321Field)(nil),
      (*Int641Field)(nil),
      (*Uint1Field)(nil),
      (*Uint81Field)(nil),
      (*Uint161Field)(nil),
      (*Uint321Field)(nil),
      (*Uint641Field)(nil),
      (*Float321Field)(nil),
      (*Float641Field)(nil),
      (*Bool1Field)(nil),
      (*String1Field)(nil),
      (*Binary1Field)(nil),
      (*Decimal1281Field)(nil),
      (*Fields1Field)(nil),
    }
  
    fieldNews := []interface{}{
      NewInt1Field,
      NewInt81Field,
      NewInt161Field,
      NewInt321Field,
      NewInt641Field,
      NewUint1Field,
      NewUint81Field,
      NewUint161Field,
      NewUint321Field,
      NewUint641Field,
      NewFloat321Field,
      NewFloat641Field,
      NewBool1Field,
      NewString1Field,
      NewBinary1Field,
      NewDecimal1281Field,
      NewFields1Field,
    }
  
    for i,t := range all1 {
      b.RegisterType(t, Type{
        F:    reflect.TypeOf(fieldTypes[i]).Elem(),
        NewF: NewTypByFunc(fieldNews[i]),
      })
    }
  }
}

