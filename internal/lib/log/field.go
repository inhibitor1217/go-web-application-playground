package log

import "go.uber.org/zap"

type Field struct {
	Key   string
	Type  string
	Value interface{}
}

func String(key string, value string) Field {
	return Field{Key: key, Type: "string", Value: value}
}

func Error(err error) Field {
	return Field{Key: "error", Type: "error", Value: err}
}

func zapField(field Field) zap.Field {
	switch field.Type {
	case "string":
		return zap.String(field.Key, field.Value.(string))
	case "error":
		return zap.Error(field.Value.(error))
	default:
		return zap.Any(field.Key, field.Value)
	}
}

func zapFields(fields []Field) []zap.Field {
	zapFields := make([]zap.Field, len(fields))
	for i, field := range fields {
		zapFields[i] = zapField(field)
	}
	return zapFields
}
