package inktest

import "time"

// stringP gives the pointer to the given string value.
func stringP(val string) *string {
	return &val
}

// int64P gives the pointer to the given int64 value.
func int64P(val int64) *int64 {
	return &val
}

// int32P gives the pointer to the given int32 value.
func int32P(val int32) *int32 {
	return &val
}

// int16P gives the pointer to the given int16 value.
func int16P(val int16) *int16 {
	return &val
}

// int8P gives the pointer to the given int8 value.
func int8P(val int8) *int8 {
	return &val
}

// intP gives the pointer to the given int value.
func intP(val int) *int {
	return &val
}

// uint64P gives the pointer to the given uint64 value.
func uint64P(val uint64) *uint64 {
	return &val
}

// uint32P gives the pointer to the given uint32 value.
func uint32P(val uint32) *uint32 {
	return &val
}

// uint16P gives the pointer to the given uint16 value.
func uint16P(val uint16) *uint16 {
	return &val
}

// uint8P gives the pointer to the given uint8 value.
func uint8P(val uint8) *uint8 {
	return &val
}

// uintP gives the pointer to the given uint value.
func uintP(val uint) *uint {
	return &val
}

// float64P gives the pointer to the given float64 value.
func float64P(val float64) *float64 {
	return &val
}

// float32P gives the pointer to the given float32 value.
func float32P(val float32) *float32 {
	return &val
}

// complex128P gives the pointer to the given complex128 value.
func complex128P(val complex128) *complex128 {
	return &val
}

// complex64P gives the pointer to the given complex64 value.
func complex64P(val complex64) *complex64 {
	return &val
}

// boolP gives the pointer to the given bool value.
func boolP(val bool) *bool {
	return &val
}

// timeP gives the pointer to the given time.Time value.
func timeP(val time.Time) *time.Time {
	return &val
}

// durationP gives the pointer to the given time.Duration value.
func durationP(val time.Duration) *time.Duration {
	return &val
}
