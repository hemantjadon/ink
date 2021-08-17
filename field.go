package ink

import (
	"fmt"
	"math"
	"time"
)

// FieldType describes the type of the Field.
type FieldType uint8

const (
	// FieldTypeString indicates that the field carries string.
	FieldTypeString FieldType = iota

	// FieldTypeInt64 indicates that the field carries int64.
	FieldTypeInt64

	// FieldTypeInt32 indicates that the field carries int32.
	FieldTypeInt32

	// FieldTypeInt16 indicates that the field carries int16.
	FieldTypeInt16

	// FieldTypeInt8 indicates that the field carries int8.
	FieldTypeInt8

	// FieldTypeInt indicates that the field carries int.
	FieldTypeInt

	// FieldTypeUint64 indicates that the field carries uint64.
	FieldTypeUint64

	// FieldTypeUint32 indicates that the field carries uint32.
	FieldTypeUint32

	// FieldTypeUint16 indicates that the field carries uint16.
	FieldTypeUint16

	// FieldTypeUint8 indicates that the field carries uint8.
	FieldTypeUint8

	// FieldTypeUint indicates that the field carries uint.
	FieldTypeUint

	// FieldTypeFloat64 indicates that the field carries float64.
	FieldTypeFloat64

	// FieldTypeFloat32 indicates that the field carries float32.
	FieldTypeFloat32

	// FieldTypeComplex128 indicates that the field carries complex128.
	FieldTypeComplex128

	// FieldTypeComplex64 indicates that the field carries complex64.
	FieldTypeComplex64

	// FieldTypeBool indicates that the field carries bool.
	FieldTypeBool

	// FieldTypeTime indicates that the field carries time.Time.
	FieldTypeTime

	// FieldTypeDuration indicates that the field carries time.Duration.
	FieldTypeDuration

	// FieldTypeStringer indicates that the field carries a value which
	// satisfies fmt.Stringer interface.
	FieldTypeStringer

	// FieldTypeSlice indicates that the field carries a homogeneous slice of
	// basic go types.
	FieldTypeSlice

	// FieldTypeReflect indicates that the field carries generic(interface{})
	// field.
	FieldTypeReflect
)

// Field describes a field. Do not instantiate Field type directly, use the
// provided creation (constructor) functions instead.
type Field struct {
	key string
	typ FieldType
	str string
	num uint64
	ifc interface{}
}

// Key gives the "key" of the Field.
func (f Field) Key() string {
	return f.key
}

// Type gives the FieldType associated with the Field.
func (f Field) Type() FieldType {
	return f.typ
}

// String gives the string value in the Field if the type is FieldTypeString,
// otherwise it returns empty string ("").
func (f Field) String() string {
	if f.typ != FieldTypeString {
		return ""
	}
	return f.str
}

// Int64 gives the int64 value in the Field if the type is FieldTypeInt64,
// otherwise it returns 0.
func (f Field) Int64() int64 {
	if f.typ != FieldTypeInt64 {
		return 0
	}
	return int64(f.num)
}

// Int32 gives the int32 value in the Field if the type is FieldTypeInt32,
// otherwise it returns 0.
func (f Field) Int32() int32 {
	if f.typ != FieldTypeInt32 {
		return 0
	}
	return int32(f.num)
}

// Int16 gives the int16 value in the Field if the type is FieldTypeInt16,
// otherwise it returns 0.
func (f Field) Int16() int16 {
	if f.typ != FieldTypeInt16 {
		return 0
	}
	return int16(f.num)
}

// Int8 gives the int8 value in the Field if the type is FieldTypeInt8,
// otherwise it returns 0.
func (f Field) Int8() int8 {
	if f.typ != FieldTypeInt8 {
		return 0
	}
	return int8(f.num)
}

// Int gives the int value in the Field if the type is FieldTypeInt, otherwise
// it returns 0.
func (f Field) Int() int {
	if f.typ != FieldTypeInt {
		return 0
	}
	return int(f.num)
}

// Uint64 gives the uint64 value in the Field if the type is FieldTypeUint64,
// otherwise it returns 0.
func (f Field) Uint64() uint64 {
	if f.typ != FieldTypeUint64 {
		return 0
	}
	return f.num
}

// Uint32 gives the uint32 value in the Field if the type is FieldTypeUint32,
// otherwise it returns 0.
func (f Field) Uint32() uint32 {
	if f.typ != FieldTypeUint32 {
		return 0
	}
	return uint32(f.num)
}

// Uint16 gives the uint16 value in the Field if the type is FieldTypeUint16,
// otherwise it returns 0.
func (f Field) Uint16() uint16 {
	if f.typ != FieldTypeUint16 {
		return 0
	}
	return uint16(f.num)
}

// Uint8 gives the uint8 value in the Field if the type is FieldTypeUint8,
// otherwise it returns 0.
func (f Field) Uint8() uint8 {
	if f.typ != FieldTypeUint8 {
		return 0
	}
	return uint8(f.num)
}

// Uint gives the uint value in the Field if the type is FieldTypeUint,
// otherwise it returns 0.
func (f Field) Uint() uint {
	if f.typ != FieldTypeUint {
		return 0
	}
	return uint(f.num)
}

// Float64 gives the float64 value in the Field if the type is FieldTypeFloat64,
// otherwise it returns 0.
func (f Field) Float64() float64 {
	if f.typ != FieldTypeFloat64 {
		return 0
	}
	return math.Float64frombits(f.num)
}

// Float32 gives the float32 value in the Field if the type is FieldTypeFloat32,
// otherwise it returns 0.
func (f Field) Float32() float32 {
	if f.typ != FieldTypeFloat32 {
		return 0
	}
	return math.Float32frombits(uint32(f.num))
}

// Complex128 gives the complex128 value in the Field if the type is
// FieldTypeComplex128, otherwise it returns 0.
func (f Field) Complex128() complex128 {
	if f.typ != FieldTypeComplex128 {
		return 0
	}
	return f.ifc.(complex128)
}

// Complex64 gives the complex64 value in the Field if the type is
// FieldTypeComplex64, otherwise it returns 0.
func (f Field) Complex64() complex64 {
	if f.typ != FieldTypeComplex64 {
		return 0
	}
	return f.ifc.(complex64)
}

// Bool gives the bool value in the Field if the type is FieldTypeBool,
// otherwise it returns false.
func (f Field) Bool() bool {
	if f.typ != FieldTypeBool {
		return false
	}
	return f.num != 0
}

// Time gives the time.Time in the Field if the type is FieldTypeTime, otherwise
// it returns zero time.Time.
func (f Field) Time() time.Time {
	if f.typ != FieldTypeTime {
		return time.Time{}
	}
	return f.ifc.(time.Time)
}

// Duration gives the time.Duration in the Field if the type is
// FieldTypeDuration, otherwise it returns 0.
func (f Field) Duration() time.Duration {
	if f.typ != FieldTypeDuration {
		return 0
	}
	return time.Duration(f.num)
}

// Stringer gives the fmt.Stringer in the Field if the type is
// FieldTypeStringer, otherwise it returns nil.
func (f Field) Stringer() fmt.Stringer {
	if f.typ != FieldTypeStringer {
		return nil
	}
	return f.ifc.(fmt.Stringer)
}

// Slice gives the slice in the Field if the type is FieldTypeSlice, otherwise
// it returns nil.
func (f Field) Slice() interface{} {
	if f.typ != FieldTypeSlice {
		return nil
	}
	return f.ifc
}

// Reflect gives the interface{} value in the field if the type is
// FieldTypeReflect, otherwise it returns nil.
func (f Field) Reflect() interface{} {
	if f.typ != FieldTypeReflect {
		return nil
	}
	return f.ifc
}

// Value gives the field value irrespective of the FieldType of the field.
//
// As it returns an interface{} it allocates memory (for interface{}) and looses
// type-safety, so it should only be used in cases when performance is not a
// serious concern, and also type safety can be ignored.
func (f Field) Value() interface{} {
	switch f.typ {
	case FieldTypeString:
		return f.String()
	case FieldTypeInt64:
		return f.Int64()
	case FieldTypeInt32:
		return f.Int32()
	case FieldTypeInt16:
		return f.Int16()
	case FieldTypeInt8:
		return f.Int8()
	case FieldTypeInt:
		return f.Int()
	case FieldTypeUint64:
		return f.Uint64()
	case FieldTypeUint32:
		return f.Uint32()
	case FieldTypeUint16:
		return f.Uint16()
	case FieldTypeUint8:
		return f.Uint8()
	case FieldTypeUint:
		return f.Uint()
	case FieldTypeFloat64:
		return f.Float64()
	case FieldTypeFloat32:
		return f.Float32()
	case FieldTypeComplex128:
		return f.Complex128()
	case FieldTypeComplex64:
		return f.Complex64()
	case FieldTypeBool:
		return f.Bool()
	case FieldTypeTime:
		return f.Time()
	case FieldTypeDuration:
		return f.Duration()
	case FieldTypeStringer:
		return f.Stringer()
	case FieldTypeSlice:
		return f.Slice()
	case FieldTypeReflect:
		return f.Reflect()
	default:
		return nil
	}
}

// String creates a new field with string type value.
func String(key string, value string) Field {
	return Field{typ: FieldTypeString, key: key, str: value}
}

// StringP creates a new field with *string type value. It safely and explicitly
// represents a "nil" when appropriate.
func StringP(key string, value *string) Field {
	if value == nil {
		return Reflect(key, nil)
	}
	return String(key, *value)
}

// Strings creates a new field with []string type value.
func Strings(key string, value []string) Field {
	if value == nil {
		return Field{typ: FieldTypeSlice, key: key, ifc: value}
	}
	slice := make([]string, 0, len(value))
	slice = append(slice, value...)
	return Field{typ: FieldTypeSlice, key: key, ifc: slice}
}

// Int64 creates a new field with int64 type value.
func Int64(key string, value int64) Field {
	return Field{typ: FieldTypeInt64, key: key, num: uint64(value)}
}

// Int64P creates a new field with *int64 type value. It safely and explicitly
// represents a "nil" when appropriate.
func Int64P(key string, value *int64) Field {
	if value == nil {
		return Reflect(key, nil)
	}
	return Int64(key, *value)
}

// Int64s creates a new field with []int64 type value.
func Int64s(key string, value []int64) Field {
	if value == nil {
		return Field{typ: FieldTypeSlice, key: key, ifc: value}
	}
	slice := make([]int64, 0, len(value))
	slice = append(slice, value...)
	return Field{typ: FieldTypeSlice, key: key, ifc: slice}
}

// Int32 creates a new field with int32 type value.
func Int32(key string, value int32) Field {
	return Field{typ: FieldTypeInt32, key: key, num: uint64(value)}
}

// Int32P creates a new field with *int32 type value. It safely and explicitly
// represents a "nil" when appropriate.
func Int32P(key string, value *int32) Field {
	if value == nil {
		return Reflect(key, nil)
	}
	return Int32(key, *value)
}

// Int32s creates a new field with []int32 type value.
func Int32s(key string, value []int32) Field {
	if value == nil {
		return Field{typ: FieldTypeSlice, key: key, ifc: value}
	}
	slice := make([]int32, 0, len(value))
	slice = append(slice, value...)
	return Field{typ: FieldTypeSlice, key: key, ifc: slice}
}

// Int16 creates a new field with int16 type value.
func Int16(key string, value int16) Field {
	return Field{typ: FieldTypeInt16, key: key, num: uint64(value)}
}

// Int16P creates a new field with *int16 type value. It safely and explicitly
// represents a "nil" when appropriate.
func Int16P(key string, value *int16) Field {
	if value == nil {
		return Reflect(key, nil)
	}
	return Int16(key, *value)
}

// Int16s creates a new field with []int16 type value.
func Int16s(key string, value []int16) Field {
	if value == nil {
		return Field{typ: FieldTypeSlice, key: key, ifc: value}
	}
	slice := make([]int16, 0, len(value))
	slice = append(slice, value...)
	return Field{typ: FieldTypeSlice, key: key, ifc: slice}
}

// Int8 creates a new field with int8 type value.
func Int8(key string, value int8) Field {
	return Field{typ: FieldTypeInt8, key: key, num: uint64(value)}
}

// Int8P creates a new field with *int8 type value. It safely and explicitly
// represents a "nil" when appropriate.
func Int8P(key string, value *int8) Field {
	if value == nil {
		return Reflect(key, nil)
	}
	return Int8(key, *value)
}

// Int8s creates a new field with []int8 type value.
func Int8s(key string, value []int8) Field {
	if value == nil {
		return Field{typ: FieldTypeSlice, key: key, ifc: value}
	}
	slice := make([]int8, 0, len(value))
	slice = append(slice, value...)
	return Field{typ: FieldTypeSlice, key: key, ifc: slice}
}

// Int creates a new field with int type value.
func Int(key string, value int) Field {
	return Field{typ: FieldTypeInt, key: key, num: uint64(value)}
}

// IntP creates a new field with *int type value. It safely and explicitly
// represents a "nil" when appropriate.
func IntP(key string, value *int) Field {
	if value == nil {
		return Reflect(key, nil)
	}
	return Int(key, *value)
}

// Ints creates a new field with []int type value.
func Ints(key string, value []int) Field {
	if value == nil {
		return Field{typ: FieldTypeSlice, key: key, ifc: value}
	}
	slice := make([]int, 0, len(value))
	slice = append(slice, value...)
	return Field{typ: FieldTypeSlice, key: key, ifc: slice}
}

// Uint64 creates a new field with uint64 type value.
func Uint64(key string, value uint64) Field {
	return Field{typ: FieldTypeUint64, key: key, num: value}
}

// Uint64P creates a new field with *uint64 type value. It safely and explicitly
// represents a "nil" when appropriate.
func Uint64P(key string, value *uint64) Field {
	if value == nil {
		return Reflect(key, nil)
	}
	return Uint64(key, *value)
}

// Uint64s creates a new field with []uint64 type value.
func Uint64s(key string, value []uint64) Field {
	if value == nil {
		return Field{typ: FieldTypeSlice, key: key, ifc: value}
	}
	slice := make([]uint64, 0, len(value))
	slice = append(slice, value...)
	return Field{typ: FieldTypeSlice, key: key, ifc: slice}
}

// Uint32 creates a new field with uint32 type value.
func Uint32(key string, value uint32) Field {
	return Field{typ: FieldTypeUint32, key: key, num: uint64(value)}
}

// Uint32P creates a new field with *uint32 type value. It safely and explicitly
// represents a "nil" when appropriate.
func Uint32P(key string, value *uint32) Field {
	if value == nil {
		return Reflect(key, nil)
	}
	return Uint32(key, *value)
}

// Uint32s creates a new field with []uint32 type value.
func Uint32s(key string, value []uint32) Field {
	if value == nil {
		return Field{typ: FieldTypeSlice, key: key, ifc: value}
	}
	slice := make([]uint32, 0, len(value))
	slice = append(slice, value...)
	return Field{typ: FieldTypeSlice, key: key, ifc: slice}
}

// Uint16 creates a new field with uint16 type value.
func Uint16(key string, value uint16) Field {
	return Field{typ: FieldTypeUint16, key: key, num: uint64(value)}
}

// Uint16P creates a new field with *uint16 type value. It safely and explicitly
// represents a "nil" when appropriate.
func Uint16P(key string, value *uint16) Field {
	if value == nil {
		return Reflect(key, nil)
	}
	return Uint16(key, *value)
}

// Uint16s creates a new field with []uint16 type value.
func Uint16s(key string, value []uint16) Field {
	if value == nil {
		return Field{typ: FieldTypeSlice, key: key, ifc: value}
	}
	slice := make([]uint16, 0, len(value))
	slice = append(slice, value...)
	return Field{typ: FieldTypeSlice, key: key, ifc: slice}
}

// Uint8 creates a new field with uint8 type value.
func Uint8(key string, value uint8) Field {
	return Field{typ: FieldTypeUint8, key: key, num: uint64(value)}
}

// Uint8P creates a new field with *uint8 type value. It safely and explicitly
// represents a "nil" when appropriate.
func Uint8P(key string, value *uint8) Field {
	if value == nil {
		return Reflect(key, nil)
	}
	return Uint8(key, *value)
}

// Uint8s creates a new field with []uint8 type value.
func Uint8s(key string, value []uint8) Field {
	if value == nil {
		return Field{typ: FieldTypeSlice, key: key, ifc: value}
	}
	slice := make([]uint8, 0, len(value))
	slice = append(slice, value...)
	return Field{typ: FieldTypeSlice, key: key, ifc: slice}
}

// Uint creates a new field with uint type value.
func Uint(key string, value uint) Field {
	return Field{typ: FieldTypeUint, key: key, num: uint64(value)}
}

// UintP creates a new field with *uint type value. It safely and explicitly
// represents a "nil" when appropriate.
func UintP(key string, value *uint) Field {
	if value == nil {
		return Reflect(key, nil)
	}
	return Uint(key, *value)
}

// Uints creates a new field with []uint type value.
func Uints(key string, value []uint) Field {
	if value == nil {
		return Field{typ: FieldTypeSlice, key: key, ifc: value}
	}
	slice := make([]uint, 0, len(value))
	slice = append(slice, value...)
	return Field{typ: FieldTypeSlice, key: key, ifc: slice}
}

// Float64 creates a new field with float64 type value.
func Float64(key string, value float64) Field {
	return Field{typ: FieldTypeFloat64, key: key, num: math.Float64bits(value)}
}

// Float64P creates a new field with *float64 type value. It safely and
// explicitly represents a "nil" when appropriate.
func Float64P(key string, value *float64) Field {
	if value == nil {
		return Reflect(key, nil)
	}
	return Float64(key, *value)
}

// Float64s creates a new field with []float64 type value.
func Float64s(key string, value []float64) Field {
	if value == nil {
		return Field{typ: FieldTypeSlice, key: key, ifc: value}
	}
	slice := make([]float64, 0, len(value))
	slice = append(slice, value...)
	return Field{typ: FieldTypeSlice, key: key, ifc: slice}
}

// Float32 creates a new field with float32 type value.
func Float32(key string, value float32) Field {
	return Field{typ: FieldTypeFloat32, key: key, num: uint64(math.Float32bits(value))}
}

// Float32P creates a new field with *float32 type value. It safely and
// explicitly represents a "nil" when appropriate.
func Float32P(key string, value *float32) Field {
	if value == nil {
		return Reflect(key, nil)
	}
	return Float32(key, *value)
}

// Float32s creates a new field with []float32 type value.
func Float32s(key string, value []float32) Field {
	if value == nil {
		return Field{typ: FieldTypeSlice, key: key, ifc: value}
	}
	slice := make([]float32, 0, len(value))
	slice = append(slice, value...)
	return Field{typ: FieldTypeSlice, key: key, ifc: slice}
}

// Complex128 creates a new field with complex128 type value.
func Complex128(key string, value complex128) Field {
	return Field{typ: FieldTypeComplex128, key: key, ifc: value}
}

// Complex128P creates a new field with *complex128 type value. It safely and
// explicitly represents a "nil" when appropriate.
func Complex128P(key string, value *complex128) Field {
	if value == nil {
		return Reflect(key, nil)
	}
	return Complex128(key, *value)
}

// Complex128s creates a new field with []complex128 type value.
func Complex128s(key string, value []complex128) Field {
	if value == nil {
		return Field{typ: FieldTypeSlice, key: key, ifc: value}
	}
	slice := make([]complex128, 0, len(value))
	slice = append(slice, value...)
	return Field{typ: FieldTypeSlice, key: key, ifc: slice}
}

// Complex64 creates a new field with complex64 type value.
func Complex64(key string, value complex64) Field {
	return Field{typ: FieldTypeComplex64, key: key, ifc: value}
}

// Complex64P creates a new field with *complex64 type value. It safely and
// explicitly represents a "nil" when appropriate.
func Complex64P(key string, value *complex64) Field {
	if value == nil {
		return Reflect(key, nil)
	}
	return Complex64(key, *value)
}

// Complex64s creates a new field with []complex64 type value.
func Complex64s(key string, value []complex64) Field {
	if value == nil {
		return Field{typ: FieldTypeSlice, key: key, ifc: value}
	}
	slice := make([]complex64, 0, len(value))
	slice = append(slice, value...)
	return Field{typ: FieldTypeSlice, key: key, ifc: slice}
}

// Bool creates a new field with bool type value.
func Bool(key string, value bool) Field {
	var v uint64
	if value {
		v = 1
	}
	return Field{typ: FieldTypeBool, key: key, num: v}
}

// BoolP creates a new field with *bool type value. It safely and explicitly
// represents a "nil" when appropriate.
func BoolP(key string, value *bool) Field {
	if value == nil {
		return Reflect(key, nil)
	}
	return Bool(key, *value)
}

// Bools creates a new field with []bool type value.
func Bools(key string, value []bool) Field {
	if value == nil {
		return Field{typ: FieldTypeSlice, key: key, ifc: value}
	}
	slice := make([]bool, 0, len(value))
	slice = append(slice, value...)
	return Field{typ: FieldTypeSlice, key: key, ifc: slice}
}

// Time creates a new field with time.Time type value.
func Time(key string, value time.Time) Field {
	return Field{typ: FieldTypeTime, key: key, ifc: value}
}

// TimeP creates a new field with *time.Time type value. It safely and
// explicitly represents a "nil" when appropriate.
func TimeP(key string, value *time.Time) Field {
	if value == nil {
		return Reflect(key, nil)
	}
	return Time(key, *value)
}

// Times creates a new field with []time.Time type value.
func Times(key string, value []time.Time) Field {
	if value == nil {
		return Field{typ: FieldTypeSlice, key: key, ifc: value}
	}
	slice := make([]time.Time, 0, len(value))
	slice = append(slice, value...)
	return Field{typ: FieldTypeSlice, key: key, ifc: slice}
}

// Duration creates a new field with time.Duration type value.
func Duration(key string, value time.Duration) Field {
	return Field{typ: FieldTypeDuration, key: key, num: uint64(value)}
}

// DurationP creates a new field with *time.Duration type value. It safely and
// explicitly represents a "nil" when appropriate.
func DurationP(key string, value *time.Duration) Field {
	if value == nil {
		return Reflect(key, nil)
	}
	return Duration(key, *value)
}

// Durations creates a new field with []time.Duration type value.
func Durations(key string, value []time.Duration) Field {
	if value == nil {
		return Field{typ: FieldTypeSlice, key: key, ifc: value}
	}
	slice := make([]time.Duration, 0, len(value))
	slice = append(slice, value...)
	return Field{typ: FieldTypeSlice, key: key, ifc: slice}
}

// Stringer creates a new field with fmt.String type value.
func Stringer(key string, value fmt.Stringer) Field {
	if value == nil {
		return Reflect(key, nil)
	}
	return Field{typ: FieldTypeStringer, key: key, ifc: value}
}

// Reflect creates a new field with value of any arbitrary type. Handling of
// FieldTypeReflect values is implementation specific, and implementations are
// free to choose how they encode such values. Usually it is reflection based
// thus may lead to heavy allocation.
//
// Any is always a better choice, as it automatically selects the correct type.
// thus results in lesser number of allocations from the implementations.
func Reflect(key string, value interface{}) Field {
	return Field{typ: FieldTypeReflect, key: key, ifc: value}
}

// Any takes a key and an arbitrary value and chooses the best way to represent
// them as a field, falling back to a reflection-based approach only if
// necessary.
func Any(key string, value interface{}) Field {
	switch val := value.(type) {
	case string:
		return String(key, val)
	case *string:
		return StringP(key, val)
	case []string:
		return Strings(key, val)
	case int64:
		return Int64(key, val)
	case *int64:
		return Int64P(key, val)
	case []int64:
		return Int64s(key, val)
	case int32:
		return Int32(key, val)
	case *int32:
		return Int32P(key, val)
	case []int32:
		return Int32s(key, val)
	case int16:
		return Int16(key, val)
	case *int16:
		return Int16P(key, val)
	case []int16:
		return Int16s(key, val)
	case int8:
		return Int8(key, val)
	case *int8:
		return Int8P(key, val)
	case []int8:
		return Int8s(key, val)
	case int:
		return Int(key, val)
	case *int:
		return IntP(key, val)
	case []int:
		return Ints(key, val)
	case uint64:
		return Uint64(key, val)
	case *uint64:
		return Uint64P(key, val)
	case []uint64:
		return Uint64s(key, val)
	case uint32:
		return Uint32(key, val)
	case *uint32:
		return Uint32P(key, val)
	case []uint32:
		return Uint32s(key, val)
	case uint16:
		return Uint16(key, val)
	case *uint16:
		return Uint16P(key, val)
	case []uint16:
		return Uint16s(key, val)
	case uint8:
		return Uint8(key, val)
	case *uint8:
		return Uint8P(key, val)
	case []uint8:
		return Uint8s(key, val)
	case uint:
		return Uint(key, val)
	case *uint:
		return UintP(key, val)
	case []uint:
		return Uints(key, val)
	case float64:
		return Float64(key, val)
	case *float64:
		return Float64P(key, val)
	case []float64:
		return Float64s(key, val)
	case float32:
		return Float32(key, val)
	case *float32:
		return Float32P(key, val)
	case []float32:
		return Float32s(key, val)
	case complex128:
		return Complex128(key, val)
	case *complex128:
		return Complex128P(key, val)
	case []complex128:
		return Complex128s(key, val)
	case complex64:
		return Complex64(key, val)
	case *complex64:
		return Complex64P(key, val)
	case []complex64:
		return Complex64s(key, val)
	case bool:
		return Bool(key, val)
	case *bool:
		return BoolP(key, val)
	case []bool:
		return Bools(key, val)
	case time.Time:
		return Time(key, val)
	case *time.Time:
		return TimeP(key, val)
	case []time.Time:
		return Times(key, val)
	case time.Duration:
		return Duration(key, val)
	case *time.Duration:
		return DurationP(key, val)
	case []time.Duration:
		return Durations(key, val)
	case fmt.Stringer:
		return Stringer(key, val)
	default:
		return Reflect(key, val)
	}
}
