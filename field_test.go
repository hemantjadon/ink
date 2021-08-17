package ink_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/hemantjadon/ink"
)

func Check(t *testing.T, field ink.Field, typ ink.FieldType, key string, val interface{}) {
	t.Helper()

	if field.Type() != typ {
		t.Fatalf("Type(): got = %v, want = %v", field.Type(), typ)
	}
	if field.Key() != key {
		t.Fatalf("Key(): got = %v, want = %v", field.Key(), key)
	}

	gots := [22]interface{}{
		field.String(),
		field.Int64(),
		field.Int32(),
		field.Int16(),
		field.Int8(),
		field.Int(),
		field.Uint64(),
		field.Uint32(),
		field.Uint16(),
		field.Uint8(),
		field.Uint(),
		field.Float64(),
		field.Float32(),
		field.Complex128(),
		field.Complex64(),
		field.Bool(),
		field.Time(),
		field.Duration(),
		field.Stringer(),
		field.Slice(),
		field.Reflect(),
		field.Value(),
	}

	wants := [22]interface{}{
		"",
		int64(0),
		int32(0),
		int16(0),
		int8(0),
		int(0),
		uint64(0),
		uint32(0),
		uint16(0),
		uint8(0),
		uint(0),
		float64(0),
		float32(0),
		complex128(0),
		complex64(0),
		false,
		time.Time{},
		time.Duration(0),
		nil,
		nil,
		nil,
		nil,
	}

	switch typ {
	case ink.FieldTypeString:
		wants[0] = val.(string)
	case ink.FieldTypeInt64:
		wants[1] = val.(int64)
	case ink.FieldTypeInt32:
		wants[2] = val.(int32)
	case ink.FieldTypeInt16:
		wants[3] = val.(int16)
	case ink.FieldTypeInt8:
		wants[4] = val.(int8)
	case ink.FieldTypeInt:
		wants[5] = val.(int)
	case ink.FieldTypeUint64:
		wants[6] = val.(uint64)
	case ink.FieldTypeUint32:
		wants[7] = val.(uint32)
	case ink.FieldTypeUint16:
		wants[8] = val.(uint16)
	case ink.FieldTypeUint8:
		wants[9] = val.(uint8)
	case ink.FieldTypeUint:
		wants[10] = val.(uint)
	case ink.FieldTypeFloat64:
		wants[11] = val.(float64)
	case ink.FieldTypeFloat32:
		wants[12] = val.(float32)
	case ink.FieldTypeComplex128:
		wants[13] = val.(complex128)
	case ink.FieldTypeComplex64:
		wants[14] = val.(complex64)
	case ink.FieldTypeBool:
		wants[15] = val.(bool)
	case ink.FieldTypeTime:
		wants[16] = val.(time.Time)
	case ink.FieldTypeDuration:
		wants[17] = val.(time.Duration)
	case ink.FieldTypeStringer:
		wants[18] = val.(fmt.Stringer)
	case ink.FieldTypeSlice:
		wants[19] = val
	case ink.FieldTypeReflect:
		wants[20] = val
	}
	wants[21] = val

	for i := 0; i < len(gots); i++ {
		got := gots[i]
		want := wants[i]
		if !reflect.DeepEqual(got, want) {
			switch i {
			case 0:
				t.Errorf("String(): got = %v, want = %v", got, want)
			case 1:
				t.Errorf("Int64(): got = %v, want = %v", got, want)
			case 2:
				t.Errorf("Int32(): got = %v, want = %v", got, want)
			case 3:
				t.Errorf("Int16(): got = %v, want = %v", got, want)
			case 4:
				t.Errorf("Int8(): got = %v, want = %v", got, want)
			case 5:
				t.Errorf("Int(): got = %v, want = %v", got, want)
			case 6:
				t.Errorf("Uint64(): got = %v, want = %v", got, want)
			case 7:
				t.Errorf("Uint32(): got = %v, want = %v", got, want)
			case 8:
				t.Errorf("Uint16(): got = %v, want = %v", got, want)
			case 9:
				t.Errorf("Uint8(): got = %v, want = %v", got, want)
			case 10:
				t.Errorf("Uint(): got = %v, want = %v", got, want)
			case 11:
				t.Errorf("Float64(): got = %v, want = %v", got, want)
			case 12:
				t.Errorf("Float32(): got = %v, want = %v", got, want)
			case 13:
				t.Errorf("Complex128(): got = %v, want = %v", got, want)
			case 14:
				t.Errorf("Complex64(): got = %v, want = %v", got, want)
			case 15:
				t.Errorf("Bool(): got = %v, want = %v", got, want)
			case 16:
				t.Errorf("Time(): got = %v, want = %v", got, want)
			case 17:
				t.Errorf("Duration(): got = %v, want = %v", got, want)
			case 18:
				t.Errorf("Stringer(): got = %v, want = %v", got, want)
			case 19:
				t.Errorf("Slice(): got = <%T, %v>, want = <%T, %v>", got, got, want, want)
			case 20:
				t.Errorf("Reflect(): got = <%T, %v>, want = <%T, %v>", got, got, want, want)
			case 21:
				t.Errorf("Value(): got = <%T, %v>, want = <%T, %v>", got, got, want, want)
			}
		}
	}
}

func TestString(t *testing.T) {
	t.Parallel()

	key := "key"
	value := "value"
	field := ink.String(key, value)

	Check(t, field, ink.FieldTypeString, key, value)
}

func TestStringP(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := (*string)(nil)
		field := ink.StringP(key, value)

		Check(t, field, ink.FieldTypeReflect, key, nil) // untyped nil
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := "value"
		field := ink.StringP(key, &value)

		Check(t, field, ink.FieldTypeString, key, value)
	})
}

func TestStrings(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := ([]string)(nil)
		field := ink.Strings(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := []string{"go", "java", "c++"}
		field := ink.Strings(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})
}

func TestInt64(t *testing.T) {
	t.Parallel()

	key := "key"
	value := int64(1)
	field := ink.Int64(key, value)

	Check(t, field, ink.FieldTypeInt64, key, value)
}

func TestInt64P(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := (*int64)(nil)
		field := ink.Int64P(key, value)

		Check(t, field, ink.FieldTypeReflect, key, nil) // untyped nil
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := int64(1)
		field := ink.Int64P(key, &value)

		Check(t, field, ink.FieldTypeInt64, key, value)
	})
}

func TestInt64s(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := ([]int64)(nil)
		field := ink.Int64s(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := []int64{1, 2, 3}
		field := ink.Int64s(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})
}

func TestInt32(t *testing.T) {
	t.Parallel()

	key := "key"
	value := int32(1)
	field := ink.Int32(key, value)

	Check(t, field, ink.FieldTypeInt32, key, value)
}

func TestInt32P(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := (*int32)(nil)
		field := ink.Int32P(key, value)

		Check(t, field, ink.FieldTypeReflect, key, nil) // untyped nil
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := int32(1)
		field := ink.Int32P(key, &value)

		Check(t, field, ink.FieldTypeInt32, key, value)
	})
}

func TestInt32s(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := ([]int32)(nil)
		field := ink.Int32s(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := []int32{1, 2, 3}
		field := ink.Int32s(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})
}

func TestInt16(t *testing.T) {
	t.Parallel()

	key := "key"
	value := int16(1)
	field := ink.Int16(key, value)

	Check(t, field, ink.FieldTypeInt16, key, value)
}

func TestInt16P(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := (*int16)(nil)
		field := ink.Int16P(key, value)

		Check(t, field, ink.FieldTypeReflect, key, nil) // untyped nil
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := int16(1)
		field := ink.Int16P(key, &value)

		Check(t, field, ink.FieldTypeInt16, key, value)
	})
}

func TestInt16s(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := ([]int16)(nil)
		field := ink.Int16s(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := []int16{1, 2, 3}
		field := ink.Int16s(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})
}

func TestInt8(t *testing.T) {
	t.Parallel()

	key := "key"
	value := int8(1)
	field := ink.Int8(key, value)

	Check(t, field, ink.FieldTypeInt8, key, value)
}

func TestInt8P(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := (*int8)(nil)
		field := ink.Int8P(key, value)

		Check(t, field, ink.FieldTypeReflect, key, nil) // untyped nil
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := int8(1)
		field := ink.Int8P(key, &value)

		Check(t, field, ink.FieldTypeInt8, key, value)
	})
}

func TestInt8s(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := ([]int8)(nil)
		field := ink.Int8s(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := []int8{1, 2, 3}
		field := ink.Int8s(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})
}

func TestInt(t *testing.T) {
	t.Parallel()

	key := "key"
	value := 1
	field := ink.Int(key, value)

	Check(t, field, ink.FieldTypeInt, key, value)
}

func TestIntP(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := (*int)(nil)
		field := ink.IntP(key, value)

		Check(t, field, ink.FieldTypeReflect, key, nil) // untyped nil
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := int(1)
		field := ink.IntP(key, &value)

		Check(t, field, ink.FieldTypeInt, key, value)
	})
}

func TestInts(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := ([]int)(nil)
		field := ink.Ints(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := []int{1, 2, 3}
		field := ink.Ints(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})
}

func TestUint64(t *testing.T) {
	t.Parallel()

	key := "key"
	value := uint64(1)
	field := ink.Uint64(key, value)

	Check(t, field, ink.FieldTypeUint64, key, value)
}

func TestUint64P(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := (*uint64)(nil)
		field := ink.Uint64P(key, value)

		Check(t, field, ink.FieldTypeReflect, key, nil) // untyped nil
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := uint64(1)
		field := ink.Uint64P(key, &value)

		Check(t, field, ink.FieldTypeUint64, key, value)
	})
}

func TestUint64s(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := ([]uint64)(nil)
		field := ink.Uint64s(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := []uint64{1, 2, 3}
		field := ink.Uint64s(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})
}

func TestUint32(t *testing.T) {
	t.Parallel()

	key := "key"
	value := uint32(1)
	field := ink.Uint32(key, value)

	Check(t, field, ink.FieldTypeUint32, key, value)
}

func TestUint32P(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := (*uint32)(nil)
		field := ink.Uint32P(key, value)

		Check(t, field, ink.FieldTypeReflect, key, nil) // untyped nil
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := uint32(1)
		field := ink.Uint32P(key, &value)

		Check(t, field, ink.FieldTypeUint32, key, value)
	})
}

func TestUint32s(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := ([]uint32)(nil)
		field := ink.Uint32s(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := []uint32{1, 2, 3}
		field := ink.Uint32s(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})
}

func TestUint16(t *testing.T) {
	t.Parallel()

	key := "key"
	value := uint16(1)
	field := ink.Uint16(key, value)

	Check(t, field, ink.FieldTypeUint16, key, value)
}

func TestUint16P(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := (*uint16)(nil)
		field := ink.Uint16P(key, value)

		Check(t, field, ink.FieldTypeReflect, key, nil) // untyped nil
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := uint16(1)
		field := ink.Uint16P(key, &value)

		Check(t, field, ink.FieldTypeUint16, key, value)
	})
}

func TestUint16s(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := ([]uint16)(nil)
		field := ink.Uint16s(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := []uint16{1, 2, 3}
		field := ink.Uint16s(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})
}

func TestUint8(t *testing.T) {
	t.Parallel()

	key := "key"
	value := uint8(1)
	field := ink.Uint8(key, value)

	Check(t, field, ink.FieldTypeUint8, key, value)
}

func TestUint8P(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := (*uint8)(nil)
		field := ink.Uint8P(key, value)

		Check(t, field, ink.FieldTypeReflect, key, nil) // untyped nil
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := uint8(1)
		field := ink.Uint8P(key, &value)

		Check(t, field, ink.FieldTypeUint8, key, value)
	})
}

func TestUint8s(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := ([]uint8)(nil)
		field := ink.Uint8s(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := []uint8{1, 2, 3}
		field := ink.Uint8s(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})
}

func TestUint(t *testing.T) {
	t.Parallel()

	key := "key"
	value := uint(1)
	field := ink.Uint(key, value)

	Check(t, field, ink.FieldTypeUint, key, value)
}

func TestUintP(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := (*uint)(nil)
		field := ink.UintP(key, value)

		Check(t, field, ink.FieldTypeReflect, key, nil) // untyped nil
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := uint(1)
		field := ink.UintP(key, &value)

		Check(t, field, ink.FieldTypeUint, key, value)
	})
}

func TestUints(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := ([]uint)(nil)
		field := ink.Uints(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := []uint{1, 2, 3}
		field := ink.Uints(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})
}

func TestFloat64(t *testing.T) {
	t.Parallel()

	key := "key"
	value := float64(1.234)
	field := ink.Float64(key, value)

	Check(t, field, ink.FieldTypeFloat64, key, value)
}

func TestFloat64P(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := (*float64)(nil)
		field := ink.Float64P(key, value)

		Check(t, field, ink.FieldTypeReflect, key, nil) // untyped nil
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := float64(1)
		field := ink.Float64P(key, &value)

		Check(t, field, ink.FieldTypeFloat64, key, value)
	})
}

func TestFloat64s(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := ([]float64)(nil)
		field := ink.Float64s(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := []float64{3.14, 2.71, 1.61}
		field := ink.Float64s(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})
}

func TestFloat32(t *testing.T) {
	t.Parallel()

	key := "key"
	value := float32(1.234)
	field := ink.Float32(key, value)

	Check(t, field, ink.FieldTypeFloat32, key, value)
}

func TestFloat32P(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := (*float32)(nil)
		field := ink.Float32P(key, value)

		Check(t, field, ink.FieldTypeReflect, key, nil) // untyped nil
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := float32(1)
		field := ink.Float32P(key, &value)

		Check(t, field, ink.FieldTypeFloat32, key, value)
	})
}

func TestFloat32s(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := ([]float32)(nil)
		field := ink.Float32s(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := []float32{3.14, 2.71, 1.61}
		field := ink.Float32s(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})
}

func TestComplex128(t *testing.T) {
	t.Parallel()

	key := "key"
	value := complex128(3.14 + 2.71i)
	field := ink.Complex128(key, value)

	Check(t, field, ink.FieldTypeComplex128, key, value)
}

func TestComplex128P(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := (*complex128)(nil)
		field := ink.Complex128P(key, value)

		Check(t, field, ink.FieldTypeReflect, key, nil) // untyped nil
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := complex128(3.14 + 2.71i)
		field := ink.Complex128P(key, &value)

		Check(t, field, ink.FieldTypeComplex128, key, value)
	})
}

func TestComplex128s(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := ([]complex128)(nil)
		field := ink.Complex128s(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := []complex128{3.14 + 2.71i, 2.71 + 1.61i, 1.61 + 3.14i}
		field := ink.Complex128s(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})
}

func TestComplex64(t *testing.T) {
	t.Parallel()

	key := "key"
	value := complex64(3.14 + 2.71i)
	field := ink.Complex64(key, value)

	Check(t, field, ink.FieldTypeComplex64, key, value)
}

func TestComplex64P(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := (*complex64)(nil)
		field := ink.Complex64P(key, value)

		Check(t, field, ink.FieldTypeReflect, key, nil) // untyped nil
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := complex64(3.14 + 2.71i)
		field := ink.Complex64P(key, &value)

		Check(t, field, ink.FieldTypeComplex64, key, value)
	})
}

func TestComplex64s(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := ([]complex64)(nil)
		field := ink.Complex64s(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := []complex64{3.14 + 2.71i, 2.71 + 1.61i, 1.61 + 3.14i}
		field := ink.Complex64s(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})
}

func TestBool(t *testing.T) {
	t.Parallel()

	key := "key"
	value := true
	field := ink.Bool(key, value)

	Check(t, field, ink.FieldTypeBool, key, value)
}

func TestBoolP(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := (*bool)(nil)
		field := ink.BoolP(key, value)

		Check(t, field, ink.FieldTypeReflect, key, nil) // untyped nil
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := true
		field := ink.BoolP(key, &value)

		Check(t, field, ink.FieldTypeBool, key, value)
	})
}

func TestBools(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := ([]bool)(nil)
		field := ink.Bools(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := []bool{true, true, false}
		field := ink.Bools(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})
}

func TestTime(t *testing.T) {
	t.Parallel()

	key := "key"
	value := time.Date(2021, 01, 02, 03, 04, 05, 0, time.UTC)
	field := ink.Time(key, value)

	Check(t, field, ink.FieldTypeTime, key, value)
}

func TestTimeP(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := (*time.Time)(nil)
		field := ink.TimeP(key, value)

		Check(t, field, ink.FieldTypeReflect, key, nil) // untyped nil
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := time.Date(2021, 01, 02, 03, 04, 05, 0, time.UTC)
		field := ink.TimeP(key, &value)

		Check(t, field, ink.FieldTypeTime, key, value)
	})
}

func TestTimes(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := ([]time.Time)(nil)
		field := ink.Times(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		ti := time.Date(2021, 01, 02, 03, 04, 05, 0, time.UTC)
		value := []time.Time{ti, ti.Add(1 * time.Hour), ti.Add(2 * time.Hour)}
		field := ink.Times(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})
}

func TestDuration(t *testing.T) {
	t.Parallel()

	key := "key"
	value := 5 * time.Second
	field := ink.Duration(key, value)

	Check(t, field, ink.FieldTypeDuration, key, value)
}

func TestDurationP(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := (*time.Duration)(nil)
		field := ink.DurationP(key, value)

		Check(t, field, ink.FieldTypeReflect, key, nil) // untyped nil
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := time.Duration(1)
		field := ink.DurationP(key, &value)

		Check(t, field, ink.FieldTypeDuration, key, value)
	})
}

func TestDurations(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := ([]time.Duration)(nil)
		field := ink.Durations(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := []time.Duration{1 * time.Second, 2 * time.Second, 3 * time.Second}
		field := ink.Durations(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})
}

func TestStringer(t *testing.T) {
	t.Parallel()

	t.Run(`nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := (fmt.Stringer)(nil)
		field := ink.Stringer(key, value)

		Check(t, field, ink.FieldTypeReflect, key, nil) // untyped nil
	})

	t.Run(`non nil`, func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := stringer{Str: "is a stringer"}
		field := ink.Stringer(key, value)

		Check(t, field, ink.FieldTypeStringer, key, value)
	})
}

func TestAny(t *testing.T) {
	t.Parallel()

	key := "key"

	t.Run(`string`, func(t *testing.T) {
		t.Parallel()

		value := "value"
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeString, key, value)
	})

	t.Run(`*string`, func(t *testing.T) {
		t.Parallel()

		value := "value"
		field := ink.Any(key, &value)

		Check(t, field, ink.FieldTypeString, key, value)
	})

	t.Run(`[]string`, func(t *testing.T) {
		t.Parallel()

		value := []string{"go", "java", "c++"}
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`int64`, func(t *testing.T) {
		t.Parallel()

		value := int64(1)
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeInt64, key, value)
	})

	t.Run(`*int64`, func(t *testing.T) {
		t.Parallel()

		value := int64(1)
		field := ink.Any(key, &value)

		Check(t, field, ink.FieldTypeInt64, key, value)
	})

	t.Run(`[]int64`, func(t *testing.T) {
		t.Parallel()

		value := []int64{1, 2, 3}
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`int32`, func(t *testing.T) {
		t.Parallel()

		value := int32(1)
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeInt32, key, value)
	})

	t.Run(`*int32`, func(t *testing.T) {
		t.Parallel()

		value := int32(1)
		field := ink.Any(key, &value)

		Check(t, field, ink.FieldTypeInt32, key, value)
	})

	t.Run(`[]int32`, func(t *testing.T) {
		t.Parallel()

		value := []int32{1, 2, 3}
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`int16`, func(t *testing.T) {
		t.Parallel()

		value := int16(1)
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeInt16, key, value)
	})

	t.Run(`*int16`, func(t *testing.T) {
		t.Parallel()

		value := int16(1)
		field := ink.Any(key, &value)

		Check(t, field, ink.FieldTypeInt16, key, value)
	})

	t.Run(`[]int16`, func(t *testing.T) {
		t.Parallel()

		value := []int16{1, 2, 3}
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`int8`, func(t *testing.T) {
		t.Parallel()

		value := int8(1)
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeInt8, key, value)
	})

	t.Run(`*int8`, func(t *testing.T) {
		t.Parallel()

		value := int8(1)
		field := ink.Any(key, &value)

		Check(t, field, ink.FieldTypeInt8, key, value)
	})

	t.Run(`[]int8`, func(t *testing.T) {
		t.Parallel()

		value := []int8{1, 2, 3}
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`int`, func(t *testing.T) {
		t.Parallel()

		value := 1
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeInt, key, value)
	})

	t.Run(`*int`, func(t *testing.T) {
		t.Parallel()

		value := int(1)
		field := ink.Any(key, &value)

		Check(t, field, ink.FieldTypeInt, key, value)
	})

	t.Run(`[]int`, func(t *testing.T) {
		t.Parallel()

		value := []int{1, 2, 3}
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`uint64`, func(t *testing.T) {
		t.Parallel()

		value := uint64(1)
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeUint64, key, value)
	})

	t.Run(`*uint64`, func(t *testing.T) {
		t.Parallel()

		value := uint64(1)
		field := ink.Any(key, &value)

		Check(t, field, ink.FieldTypeUint64, key, value)
	})

	t.Run(`[]uint64`, func(t *testing.T) {
		t.Parallel()

		value := []uint64{1, 2, 3}
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`uint32`, func(t *testing.T) {
		t.Parallel()

		value := uint32(1)
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeUint32, key, value)
	})

	t.Run(`*uint32`, func(t *testing.T) {
		t.Parallel()

		value := uint32(1)
		field := ink.Any(key, &value)

		Check(t, field, ink.FieldTypeUint32, key, value)
	})

	t.Run(`[]uint32`, func(t *testing.T) {
		t.Parallel()

		value := []uint32{1, 2, 3}
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`uint16`, func(t *testing.T) {
		t.Parallel()

		value := uint16(1)
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeUint16, key, value)
	})

	t.Run(`*uint16`, func(t *testing.T) {
		t.Parallel()

		value := uint16(1)
		field := ink.Any(key, &value)

		Check(t, field, ink.FieldTypeUint16, key, value)
	})

	t.Run(`[]uint16`, func(t *testing.T) {
		t.Parallel()

		value := []uint16{1, 2, 3}
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`uint8`, func(t *testing.T) {
		t.Parallel()

		value := uint8(1)
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeUint8, key, value)
	})

	t.Run(`*uint8`, func(t *testing.T) {
		t.Parallel()

		value := uint8(1)
		field := ink.Any(key, &value)

		Check(t, field, ink.FieldTypeUint8, key, value)
	})

	t.Run(`[]uint8`, func(t *testing.T) {
		t.Parallel()

		value := []uint8{1, 2, 3}
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`uint`, func(t *testing.T) {
		t.Parallel()

		value := uint(1)
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeUint, key, value)
	})

	t.Run(`*uint`, func(t *testing.T) {
		t.Parallel()

		value := uint(1)
		field := ink.Any(key, &value)

		Check(t, field, ink.FieldTypeUint, key, value)
	})

	t.Run(`[]uint`, func(t *testing.T) {
		t.Parallel()

		value := []uint{1, 2, 3}
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`float64`, func(t *testing.T) {
		t.Parallel()

		value := float64(1.234)
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeFloat64, key, value)
	})

	t.Run(`*float64`, func(t *testing.T) {
		t.Parallel()

		value := float64(1.234)
		field := ink.Any(key, &value)

		Check(t, field, ink.FieldTypeFloat64, key, value)
	})

	t.Run(`[]float64`, func(t *testing.T) {
		t.Parallel()

		value := []float64{3.14, 2.71, 1.61}
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`float32`, func(t *testing.T) {
		t.Parallel()

		value := float32(1.234)
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeFloat32, key, value)
	})

	t.Run(`*float32`, func(t *testing.T) {
		t.Parallel()

		value := float32(1.234)
		field := ink.Any(key, &value)

		Check(t, field, ink.FieldTypeFloat32, key, value)
	})

	t.Run(`[]float32`, func(t *testing.T) {
		t.Parallel()

		value := []float32{3.14, 2.71, 1.61}
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`complex128`, func(t *testing.T) {
		t.Parallel()

		value := complex128(3.14 + 2.71i)
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeComplex128, key, value)
	})

	t.Run(`*complex128`, func(t *testing.T) {
		t.Parallel()

		value := complex128(3.14 + 2.71i)
		field := ink.Any(key, &value)

		Check(t, field, ink.FieldTypeComplex128, key, value)
	})

	t.Run(`[]complex128`, func(t *testing.T) {
		t.Parallel()

		value := []complex128{3.14 + 2.71i, 2.71 + 1.61i, 1.61 + 3.14i}
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`complex64`, func(t *testing.T) {
		t.Parallel()

		value := complex64(3.14 + 2.71i)
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeComplex64, key, value)
	})

	t.Run(`*complex64`, func(t *testing.T) {
		t.Parallel()

		value := complex64(3.14 + 2.71i)
		field := ink.Any(key, &value)

		Check(t, field, ink.FieldTypeComplex64, key, value)
	})

	t.Run(`[]complex64`, func(t *testing.T) {
		t.Parallel()

		value := []complex64{3.14 + 2.71i, 2.71 + 1.61i, 1.61 + 3.14i}
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`bool`, func(t *testing.T) {
		t.Parallel()

		value := true
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeBool, key, value)
	})

	t.Run(`*bool`, func(t *testing.T) {
		t.Parallel()

		value := false
		field := ink.Any(key, &value)

		Check(t, field, ink.FieldTypeBool, key, value)
	})

	t.Run(`[]bool`, func(t *testing.T) {
		t.Parallel()

		value := []bool{true, true, false}
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`time.Time`, func(t *testing.T) {
		t.Parallel()

		value := time.Date(2021, 01, 02, 03, 04, 05, 0, time.UTC)
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeTime, key, value)
	})

	t.Run(`*time.Time`, func(t *testing.T) {
		t.Parallel()

		value := time.Date(2021, 01, 02, 03, 04, 05, 0, time.UTC)
		field := ink.Any(key, &value)

		Check(t, field, ink.FieldTypeTime, key, value)
	})

	t.Run(`[]time.Time`, func(t *testing.T) {
		t.Parallel()

		ti := time.Date(2021, 01, 02, 03, 04, 05, 0, time.UTC)
		value := []time.Time{ti, ti.Add(1 * time.Hour), ti.Add(2 * time.Hour)}
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`time.Duration`, func(t *testing.T) {
		t.Parallel()

		value := 1 * time.Second
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeDuration, key, value)
	})

	t.Run(`*time.Duration`, func(t *testing.T) {
		t.Parallel()

		value := 1 * time.Second
		field := ink.Any(key, &value)

		Check(t, field, ink.FieldTypeDuration, key, value)
	})

	t.Run(`[]time.Duration`, func(t *testing.T) {
		t.Parallel()

		value := []time.Duration{1 * time.Second, 2 * time.Second, 3 * time.Second}
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeSlice, key, value)
	})

	t.Run(`fmt.Stringer`, func(t *testing.T) {
		t.Parallel()

		value := stringer{Str: "is a stringer"}
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeStringer, key, value)
	})

	t.Run(`others`, func(t *testing.T) {
		t.Parallel()

		type test struct {
			f1 string
			f2 int
		}

		key := "key"
		value := test{f1: "hemant", f2: 25}
		field := ink.Any(key, value)

		Check(t, field, ink.FieldTypeReflect, key, value)
	})
}

type stringer struct {
	Str string
}

func (s stringer) String() string {
	return s.Str
}
