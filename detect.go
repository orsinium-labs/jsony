package jsony

// Union of built-in signed integer types.
//
// Doesn't support integer subtypes.
type ints interface {
	int | int8 | int16 | int32 | int64
}

// Union of built-in unsigned integer types.
//
// Doesn't support subtypes.
type uints interface {
	uint | uintptr | uint8 | uint16 | uint32 | uint64
}

// Detect the jsony type of a primitive value base on generic union.
//
// If the concrete type is statically known, prefer using
// the constructor for this type instead. For example, [Int] for int,
// [String] for string, etc.
//
// If the type is not statically known, you can use [UnsafeDetect] instead.
func Detect[T ints | uints | string | bool](v T) Encoder {
	return UnsafeDetect(v)
}

// Detect the jsony type of a primitive value in runtime.
//
// Supports only primitive types (int, uint8, string, etc),
// without their subtypes, slices, or structs.
// Returns nil for unknown type.
//
// If the concrete type is statically known, prefer using
// the constructor for this type instead. For example, [Int] for int,
// [String] for string, etc.
//
// If the type is constrained by a generic enum, use [Detect] instead.
func UnsafeDetect(val any) Encoder {
	if val == nil {
		return Null
	}
	switch v := val.(type) {
	case string:
		return String(v)
	case bool:
		return Bool(v)
	case int:
		return Int(v)
	case int8:
		return Int8(v)
	case int16:
		return Int16(v)
	case int32:
		return Int32(v)
	case int64:
		return Int64(v)
	case uint:
		return UInt(v)
	case uintptr:
		return UIntPtr(v)
	case uint8:
		return UInt8(v)
	case uint16:
		return UInt16(v)
	case uint32:
		return UInt32(v)
	case uint64:
		return UInt64(v)
	case float32:
		return Float32(v)
	case float64:
		return Float64(v)
	default:
		return nil
	}
}
