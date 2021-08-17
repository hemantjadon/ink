package ink_test

import (
	"context"
	"testing"
	"time"

	"github.com/hemantjadon/ink"
)

func TestNewContextWithFields(t *testing.T) {
	ctx := context.Background()

	t.Run("no existing fields", func(t *testing.T) {
		ctx := ink.ContextWithFields(ctx, ink.String("key1", "value1"), ink.Int("key2", 1))

		fields := ink.ContextFields(ctx)
		if len(fields) != 2 {
			t.Fatalf("got fields count = %v, want fields count = %v", len(fields), 2)
		}

		if fields[0].Key() != "key1" {
			t.Fatalf("fields[0]: got Key() = %v, want Key() = %v", fields[0].Key(), "key1")
		}
		if fields[0].Type() != ink.FieldTypeString {
			t.Fatalf("fields[0]: got Type() = %v, want Type() = %v", fields[0].Type(), ink.FieldTypeString)
		}
		if fields[0].String() != "value1" {
			t.Fatalf("fields[0]: got String() = %v, want String() = %v", fields[0].String(), "value1")
		}

		if fields[1].Key() != "key2" {
			t.Fatalf("fields[1]: got Key() = %v, want Key() = %v", fields[1].Key(), "key2")
		}
		if fields[1].Type() != ink.FieldTypeInt {
			t.Fatalf("fields[1]: got Type() = %v, want Type() = %v", fields[1].Type(), ink.FieldTypeInt)
		}
		if fields[1].Int() != 1 {
			t.Fatalf("fields[1]: got Int() = %v, want Int() = %v", fields[1].Int(), 1)
		}
	})

	t.Run("existing fields in context", func(t *testing.T) {
		ctx := ink.ContextWithFields(ctx, ink.String("key1", "value1"), ink.Int("key2", 1))
		ctx = ink.ContextWithFields(ctx, ink.Bool("key3", true), ink.Duration("key4", 1*time.Second))

		fields := ink.ContextFields(ctx)
		if len(fields) != 4 {
			t.Fatalf("len(fields): got = %v, want = %v", len(fields), 4)
		}

		if fields[0].Key() != "key1" {
			t.Fatalf("fields[0]: got Key() = %v, want Key() = %v", fields[0].Key(), "key1")
		}
		if fields[0].Type() != ink.FieldTypeString {
			t.Fatalf("fields[0]: got Type() = %v, want Type() = %v", fields[0].Type(), ink.FieldTypeString)
		}
		if fields[0].String() != "value1" {
			t.Fatalf("fields[0]: got String() = %v, want String() = %v", fields[0].String(), "value1")
		}

		if fields[1].Key() != "key2" {
			t.Fatalf("fields[1]: got Key() = %v, want Key() = %v", fields[1].Key(), "key2")
		}
		if fields[1].Type() != ink.FieldTypeInt {
			t.Fatalf("fields[1]: got Type() = %v, want Type() = %v", fields[1].Type(), ink.FieldTypeInt)
		}
		if fields[1].Int() != 1 {
			t.Fatalf("fields[1]: got Int() = %v, want Int() = %v", fields[1].Int(), 1)
		}

		if fields[2].Key() != "key3" {
			t.Fatalf("fields[2]: got Key() = %v, want Key() = %v", fields[2].Key(), "key3")
		}
		if fields[2].Type() != ink.FieldTypeBool {
			t.Fatalf("fields[2]: got Type() = %v, want Type() = %v", fields[2].Type(), ink.FieldTypeBool)
		}
		if fields[2].Bool() != true {
			t.Fatalf("fields[2]: got Bool() = %v, want Bool() = %v", fields[2].Bool(), true)
		}

		if fields[3].Key() != "key4" {
			t.Fatalf("fields[3]: got Key() = %v, want Key() = %v", fields[3].Key(), "key4")
		}
		if fields[3].Type() != ink.FieldTypeDuration {
			t.Fatalf("fields[3]: got Type() = %v, want Type() = %v", fields[3].Type(), ink.FieldTypeDuration)
		}
		if fields[3].Duration() != 1*time.Second {
			t.Fatalf("fields[3]: got Int() = %v, want Int() = %v", fields[3].Int(), 1*time.Second)
		}
	})

	t.Run("no fields given but existing fields present", func(t *testing.T) {
		ctx := ink.ContextWithFields(ctx, ink.String("key1", "value1"), ink.Int("key2", 1))
		ctx = ink.ContextWithFields(ctx)

		fields := ink.ContextFields(ctx)
		if len(fields) != 2 {
			t.Fatalf("got fields count = %v, want fields count = %v", len(fields), 2)
		}

		if fields[0].Key() != "key1" {
			t.Fatalf("fields[0]: got Key() = %v, want Key() = %v", fields[0].Key(), "key1")
		}
		if fields[0].Type() != ink.FieldTypeString {
			t.Fatalf("fields[0]: got Type() = %v, want Type() = %v", fields[0].Type(), ink.FieldTypeString)
		}
		if fields[0].String() != "value1" {
			t.Fatalf("fields[0]: got String() = %v, want String() = %v", fields[0].String(), "value1")
		}

		if fields[1].Key() != "key2" {
			t.Fatalf("fields[1]: got Key() = %v, want Key() = %v", fields[1].Key(), "key2")
		}
		if fields[1].Type() != ink.FieldTypeInt {
			t.Fatalf("fields[1]: got Type() = %v, want Type() = %v", fields[1].Type(), ink.FieldTypeInt)
		}
		if fields[1].Int() != 1 {
			t.Fatalf("fields[1]: got Int() = %v, want Int() = %v", fields[1].Int(), 1)
		}
	})

	t.Run("no fields given and no existing fields", func(t *testing.T) {
		ctx := ink.ContextWithFields(ctx)

		fields := ink.ContextFields(ctx)
		if len(fields) != 0 {
			t.Fatalf("got fields count = %v, want fields count = %v", len(fields), 0)
		}
	})
}
