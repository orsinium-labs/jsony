package jsony

// PBool is a pointer version of [Bool].
//
// If the given value is nil, it will be serialized to null.
// Otherwise, behaves exactly as [Bool].
func PBool(v *bool) Encoder {
	if v == nil {
		return Null
	}
	return Bool(*v)
}

// PInt is a pointer version of [Int].
//
// If the given value is nil, it will be serialized to null.
// Otherwise, behaves exactly as [Int].
func PInt(v *int) Encoder {
	if v == nil {
		return Null
	}
	return Int(*v)
}

// PInt8 is a pointer version of [Int8].
//
// If the given value is nil, it will be serialized to null.
// Otherwise, behaves exactly as [Int8].
func PInt8(v *int8) Encoder {
	if v == nil {
		return Null
	}
	return Int8(*v)
}

// PInt16 is a pointer version of [Int16].
//
// If the given value is nil, it will be serialized to null.
// Otherwise, behaves exactly as [Int16].
func PInt16(v *int16) Encoder {
	if v == nil {
		return Null
	}
	return Int16(*v)
}

// PInt32 is a pointer version of [Int32].
//
// If the given value is nil, it will be serialized to null.
// Otherwise, behaves exactly as [Int32].
func PInt32(v *int32) Encoder {
	if v == nil {
		return Null
	}
	return Int32(*v)
}

// PInt64 is a pointer version of [Int64].
//
// If the given value is nil, it will be serialized to null.
// Otherwise, behaves exactly as [Int64].
func PInt64(v *int64) Encoder {
	if v == nil {
		return Null
	}
	return Int64(*v)
}

// PUInt is a pointer version of [UInt].
//
// If the given value is nil, it will be serialized to null.
// Otherwise, behaves exactly as [UInt].
func PUInt(v *uint) Encoder {
	if v == nil {
		return Null
	}
	return UInt(*v)
}

// PUInt8 is a pointer version of [UInt8].
//
// If the given value is nil, it will be serialized to null.
// Otherwise, behaves exactly as [UInt8].
func PUInt8(v *uint8) Encoder {
	if v == nil {
		return Null
	}
	return UInt8(*v)
}

// PUInt16 is a pointer version of [UInt16].
//
// If the given value is nil, it will be serialized to null.
// Otherwise, behaves exactly as [UInt16].
func PUInt16(v *uint16) Encoder {
	if v == nil {
		return Null
	}
	return UInt16(*v)
}

// PUInt32 is a pointer version of [UInt32].
//
// If the given value is nil, it will be serialized to null.
// Otherwise, behaves exactly as [UInt32].
func PUInt32(v *uint32) Encoder {
	if v == nil {
		return Null
	}
	return UInt32(*v)
}

// PUInt64 is a pointer version of [UInt64].
//
// If the given value is nil, it will be serialized to null.
// Otherwise, behaves exactly as [UInt64].
func PUInt64(v *uint64) Encoder {
	if v == nil {
		return Null
	}
	return UInt64(*v)
}

// PUIntPtr is a pointer version of [UIntPtr].
//
// If the given value is nil, it will be serialized to null.
// Otherwise, behaves exactly as [UIntPtr].
func PUIntPtr(v *uintptr) Encoder {
	if v == nil {
		return Null
	}
	return UIntPtr(*v)
}

// PFloat32 is a pointer version of [Float32].
//
// If the given value is nil, it will be serialized to null.
// Otherwise, behaves exactly as [Float32].
func PFloat32(v *float32) Encoder {
	if v == nil {
		return Null
	}
	return Float32(*v)
}

// PFloat64 is a pointer version of [Float64].
//
// If the given value is nil, it will be serialized to null.
// Otherwise, behaves exactly as [Float64].
func PFloat64(v *float64) Encoder {
	if v == nil {
		return Null
	}
	return Float64(*v)
}

// PString is a pointer version of [String].
//
// If the given value is nil, it will be serialized to null.
// Otherwise, behaves exactly as [String].
func PString(v *string) Encoder {
	if v == nil {
		return Null
	}
	return String(*v)
}
