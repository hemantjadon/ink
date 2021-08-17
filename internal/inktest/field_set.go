package inktest

import (
	"math"
	"time"

	"github.com/hemantjadon/ink"
)

func stringFieldSet() []ink.Field {
	return []ink.Field{
		ink.String("string_field", "str1"),
		ink.StringP("string_p_field", stringP("str2")),
		ink.StringP("string_p_nil_field", nil),
		ink.Strings("string_slice_field", []string{"str1", "str2"}),
		ink.Strings("string_slice_nil_field", nil),
	}
}

func int64FieldSet() []ink.Field {
	return []ink.Field{
		ink.Int64("int64_field", int64(1)),
		ink.Int64("int64_neg_field", int64(-1)),
		ink.Int64P("int64_p_field", int64P(int64(2))),
		ink.Int64P("int64_p_neg_field", int64P(int64(-2))),
		ink.Int64P("int64_p_nil_field", nil),
		ink.Int64s("int64_slice_field", []int64{-2, -1, 0, 1, 2}),
		ink.Int64s("int64_slice_nil_field", nil),
	}
}

func int32FieldSet() []ink.Field {
	return []ink.Field{
		ink.Int32("int32_field", int32(1)),
		ink.Int32("int32_neg_field", int32(-1)),
		ink.Int32P("int32_p_field", int32P(int32(2))),
		ink.Int32P("int32_p_neg_field", int32P(int32(-2))),
		ink.Int32P("int32_p_nil_field", nil),
		ink.Int32s("int32_slice_field", []int32{-2, -1, 0, 1, 2}),
		ink.Int32s("int32_slice_nil_field", nil),
	}
}

func int16FieldSet() []ink.Field {
	return []ink.Field{
		ink.Int16("int16_field", int16(1)),
		ink.Int16("int16_neg_field", int16(-1)),
		ink.Int16P("int16_p_field", int16P(int16(2))),
		ink.Int16P("int16_p_neg_field", int16P(int16(-2))),
		ink.Int16P("int16_p_nil_field", nil),
		ink.Int16s("int16_slice_field", []int16{-2, -1, 0, 1, 2}),
		ink.Int16s("int16_slice_nil_field", nil),
	}
}

func int8FieldSet() []ink.Field {
	return []ink.Field{
		ink.Int8("int8_field", int8(1)),
		ink.Int8("int8_neg_field", int8(-1)),
		ink.Int8P("int8_p_field", int8P(int8(2))),
		ink.Int8P("int8_p_neg_field", int8P(int8(-2))),
		ink.Int8P("int8_p_nil_field", nil),
		ink.Int8s("int8_slice_field", []int8{-2, -1, 0, 1, 2}),
		ink.Int8s("int8_slice_nil_field", nil),
	}
}
func intFieldSet() []ink.Field {
	return []ink.Field{
		ink.Int("int_field", int(1)),
		ink.Int("int_neg_field", int(-1)),
		ink.IntP("int_p_field", intP(int(2))),
		ink.IntP("int_p_neg_field", intP(int(-2))),
		ink.IntP("int_p_nil_field", nil),
		ink.Ints("int_slice_field", []int{-2, -1, 0, 1, 2}),
		ink.Ints("int_slice_nil_field", nil),
	}
}

func uint64FieldSet() []ink.Field {
	return []ink.Field{
		ink.Uint64("uint64_field", uint64(1)),
		ink.Uint64P("uint64_p_field", uint64P(uint64(2))),
		ink.Uint64P("uint64_p_nil_field", nil),
		ink.Uint64s("uint64_slice_field", []uint64{0, 1, 2, 3, 4}),
		ink.Uint64s("uint64_slice_nil_field", nil),
	}
}

func uint32FieldSet() []ink.Field {
	return []ink.Field{
		ink.Uint32("uint32_field", uint32(1)),
		ink.Uint32P("uint32_p_field", uint32P(uint32(2))),
		ink.Uint32P("uint32_p_nil_field", nil),
		ink.Uint32s("uint32_slice_field", []uint32{0, 1, 2, 3, 4}),
		ink.Uint32s("uint32_slice_nil_field", nil),
	}
}

func uint16FieldSet() []ink.Field {
	return []ink.Field{
		ink.Uint16("uint16_field", uint16(1)),
		ink.Uint16P("uint16_p_field", uint16P(uint16(2))),
		ink.Uint16P("uint16_p_nil_field", nil),
		ink.Uint16s("uint16_slice_field", []uint16{0, 1, 2, 3, 4}),
		ink.Uint16s("uint16_slice_nil_field", nil),
	}
}

func uint8FieldSet() []ink.Field {
	return []ink.Field{
		ink.Uint8("uint8_field", uint8(1)),
		ink.Uint8P("uint8_p_field", uint8P(uint8(2))),
		ink.Uint8P("uint8_p_nil_field", nil),
		ink.Uint8s("uint8_slice_field", []uint8{0, 1, 2, 3, 4}),
		ink.Uint8s("uint8_slice_nil_field", nil),
	}
}

func uintFieldSet() []ink.Field {
	return []ink.Field{
		ink.Uint("uint_field", uint(1)),
		ink.UintP("uint_p_field", uintP(uint(2))),
		ink.UintP("uint_p_nil_field", nil),
		ink.Uints("uint_slice_field", []uint{0, 1, 2, 3, 4}),
		ink.Uints("uint_slice_nil_field", nil),
	}
}

func float64FieldSet() []ink.Field {
	return []ink.Field{
		ink.Float64("float64_field", 1.234),
		ink.Float64("float64__neg_field", -1.234),
		ink.Float64("float64_nan_field", math.NaN()),
		ink.Float64("float64_pos_inf_field", math.Inf(1)),
		ink.Float64("float64_neg_inf_field", math.Inf(-1)),
		ink.Float64P("float64_p_field", float64P(2.345)),
		ink.Float64P("float64_p_neg_field", float64P(-2.345)),
		ink.Float64P("float64_p_nan_field", float64P(math.NaN())),
		ink.Float64P("float64_p_pos_inf_field", float64P(math.Inf(1))),
		ink.Float64P("float64_p_neg_inf_field", float64P(math.Inf(-1))),
		ink.Float64P("float64_p_nil_field", nil),
		ink.Float64s("float64_slice_field", []float64{3.14, 2.71, 1.61, math.NaN(), math.Inf(1), math.Inf(-1)}),
		ink.Float64s("float64_slice_nil_field", nil),
	}
}

func float32FieldSet() []ink.Field {
	return []ink.Field{
		ink.Float32("float32_field", 1.234),
		ink.Float32("float32__neg_field", -1.234),
		ink.Float32("float32_nan_field", float32(math.NaN())),
		ink.Float32("float32_pos_inf_field", float32(math.Inf(1))),
		ink.Float32("float32_neg_inf_field", float32(math.Inf(-1))),
		ink.Float32P("float32_p_field", float32P(2.345)),
		ink.Float32P("float32_p_neg_field", float32P(-2.345)),
		ink.Float32P("float32_p_nan_field", float32P(float32(math.NaN()))),
		ink.Float32P("float32_p_pos_inf_field", float32P(float32(math.Inf(1)))),
		ink.Float32P("float32_p_neg_inf_field", float32P(float32(math.Inf(-1)))),
		ink.Float32P("float32_p_nil_field", nil),
		ink.Float32s("float32_slice_field", []float32{3.14, 2.71, 1.61, float32(math.NaN()), float32(math.Inf(1)), float32(math.Inf(-1))}),
		ink.Float32s("float32_slice_nil_field", nil),
	}
}

func complex128FieldSet() []ink.Field {
	return []ink.Field{
		ink.Complex128("complex128_field", 3.14+2.71i),
		ink.Complex128P("complex128_p_field", complex128P(2.71-1.61i)),
		ink.Complex128P("complex128_p_nil_field", nil),
		ink.Complex128s("complex128_slice_field", []complex128{3.14 + 2.71i, 2.71 - 1.61i, 1.61 + 3.14i}),
		ink.Complex128s("complex128_slice_nil_field", nil),
	}
}

func complex64FieldSet() []ink.Field {
	return []ink.Field{
		ink.Complex64("complex64_field", 3.14+2.71i),
		ink.Complex64P("complex64_p_field", complex64P(2.71-1.61i)),
		ink.Complex64P("complex64_p_nil_field", nil),
		ink.Complex64s("complex64_slice_field", []complex64{3.14 + 2.71i, 2.71 - 1.61i, 1.61 + 3.14i}),
		ink.Complex64s("complex64_slice_nil_field", nil),
	}
}

func boolFieldSet() []ink.Field {
	return []ink.Field{
		ink.Bool("bool_field", true),
		ink.BoolP("bool_p_field", boolP(false)),
		ink.BoolP("bool_p_nil_field", nil),
		ink.Bools("bool_slice_field", []bool{true, true, false}),
		ink.Bools("bool_slice_nil_field", nil),
	}
}

func timeFieldSet() []ink.Field {
	return []ink.Field{
		ink.Time("time_field", time.Now()),
		ink.TimeP("time_p_field", timeP(time.Now().Add(24*time.Hour))),
		ink.TimeP("time_p_nil_field", nil),
		ink.Times("time_slice_field", []time.Time{time.Now(), time.Now().Add(1 * time.Hour), time.Now().Add(2 * time.Hour)}),
		ink.Times("times_slice_nil_field", nil),
	}
}

func durationFieldSet() []ink.Field {
	return []ink.Field{
		ink.Duration("duration_field", 5*time.Second),
		ink.DurationP("duration_p_field", durationP(10*time.Second)),
		ink.DurationP("duration_p_nil_field", nil),
		ink.Durations("duration_slice_field", []time.Duration{1 * time.Second, 2 * time.Second, 3 * time.Second}),
		ink.Durations("duration_slice_nil_field", nil),
	}
}

func stringerFieldSet() []ink.Field {
	return []ink.Field{
		ink.Stringer("stringer_field", stringer{}),
		ink.Stringer("stringer_nil_field", nil),
	}
}

func objectFieldSet() []ink.Field {
	obj := object{
		Field1: "value1",
		Field2: 1,
		Field3: 2,
		Field4: true,
		Field5: subObject{
			SubField1: nil,
			SubField2: []bool{true, true, false, true},
			SubField3: time.Now(),
			SubField4: 5 * time.Second,
		},
	}

	return []ink.Field{
		ink.Reflect("object_field", obj),
		ink.Reflect("object_nil_field", nil),
		ink.Reflect("object_slice_field", []object{obj, obj}),
		ink.Reflect("object_slice_nil_field", []object(nil)),
	}
}

type stringer struct{}

func (s stringer) String() string {
	return "is_stringer"
}

type object struct {
	Field1 string
	Field2 int64
	Field3 uint32
	Field4 bool
	Field5 subObject
}

type subObject struct {
	SubField1 *string
	SubField2 []bool
	SubField3 time.Time
	SubField4 time.Duration
}
