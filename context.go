package ink

import "context"

type fieldCtxKeyType struct {
}

var fieldCtxKey = fieldCtxKeyType{}

// ContextWithFields returns a new context that carries given fields.
func ContextWithFields(ctx context.Context, fields ...Field) context.Context {
	if len(fields) == 0 {
		return ctx
	}
	existing := ContextFields(ctx)
	fs := make([]Field, 0, len(existing)+len(fields))
	fs = append(fs, existing...)
	fs = append(fs, fields...)
	return context.WithValue(ctx, fieldCtxKey, fs)
}

// ContextFields gives the fields carried by the given context, if any.
func ContextFields(ctx context.Context) []Field {
	val := ctx.Value(fieldCtxKey)
	fs, ok := val.([]Field)
	if !ok {
		return nil
	}
	fields := make([]Field, 0, len(fs))
	fields = append(fields, fs...)
	return fields
}
