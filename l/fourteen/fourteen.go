package fourteen

import "reflect"

func typeElement(v any) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan any:
		return "chan"
	default:
		return "unknown type"
	}
}

func typeElement2(v any) reflect.Type {
	return reflect.TypeOf(v)
}
