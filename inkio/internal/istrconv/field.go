package istrconv

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hemantjadon/ink"
	"github.com/hemantjadon/ink/internal/pool"
)

// AppendBuilderField appends the field's string to the given strings.Builder.
func AppendBuilderField(sb *strings.Builder, field ink.Field) {
	sb.WriteString(field.Key())
	sb.WriteString(equal)

	switch field.Type() {
	case ink.FieldTypeString:
		sb.WriteString(field.String())
	case ink.FieldTypeInt64:
		sb.WriteString(strconv.FormatInt(field.Int64(), 10))
	case ink.FieldTypeInt32:
		sb.WriteString(strconv.FormatInt(int64(field.Int32()), 10))
	case ink.FieldTypeInt16:
		sb.WriteString(strconv.FormatInt(int64(field.Int16()), 10))
	case ink.FieldTypeInt8:
		sb.WriteString(strconv.FormatInt(int64(field.Int8()), 10))
	case ink.FieldTypeInt:
		sb.WriteString(strconv.FormatInt(int64(field.Int()), 10))
	case ink.FieldTypeUint64:
		sb.WriteString(strconv.FormatUint(field.Uint64(), 10))
	case ink.FieldTypeUint32:
		sb.WriteString(strconv.FormatUint(uint64(field.Uint32()), 10))
	case ink.FieldTypeUint16:
		sb.WriteString(strconv.FormatUint(uint64(field.Uint16()), 10))
	case ink.FieldTypeUint8:
		sb.WriteString(strconv.FormatUint(uint64(field.Uint8()), 10))
	case ink.FieldTypeUint:
		sb.WriteString(strconv.FormatUint(uint64(field.Uint()), 10))
	case ink.FieldTypeFloat64:
		sb.WriteString(strconv.FormatFloat(field.Float64(), 'g', -1, 64))
	case ink.FieldTypeFloat32:
		sb.WriteString(strconv.FormatFloat(float64(field.Float32()), 'g', -1, 32))
	case ink.FieldTypeComplex128:
		sb.WriteString(strconv.FormatComplex(field.Complex128(), 'g', -1, 128))
	case ink.FieldTypeComplex64:
		sb.WriteString(strconv.FormatComplex(complex128(field.Complex64()), 'g', -1, 64))
	case ink.FieldTypeBool:
		sb.WriteString(strconv.FormatBool(field.Bool()))
	case ink.FieldTypeTime:
		sb.WriteString(field.Time().Format(time.RFC3339))
	case ink.FieldTypeDuration:
		sb.WriteString(field.Duration().String())
	case ink.FieldTypeStringer:
		sb.WriteString(field.Stringer().String())
	case ink.FieldTypeSlice:
		sb.WriteString(fmt.Sprintf("%v", field.Slice()))
	case ink.FieldTypeReflect:
		sb.WriteString(fmt.Sprintf("%v", field.Reflect()))
	default:
		sb.WriteString(unknown)
	}
}

// FormatField formats the field to a string.
func FormatField(field ink.Field) string {
	sb := pool.GetStringsBuilder()
	defer pool.PutStringsBuilder(sb)

	AppendBuilderField(sb, field)

	return sb.String()
}

const (
	equal   = "="
	unknown = "<unknown-field-type>"
)
