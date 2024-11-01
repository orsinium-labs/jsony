package jsony

func PBool(v *bool) Encoder {
	if v == nil {
		return Null
	}
	return Bool(*v)
}

func PInt(v *int) Encoder {
	if v == nil {
		return Null
	}
	return Int(*v)
}

func PInt8(v *int8) Encoder {
	if v == nil {
		return Null
	}
	return Int8(*v)
}

func PInt16(v *int16) Encoder {
	if v == nil {
		return Null
	}
	return Int16(*v)
}

func PInt32(v *int32) Encoder {
	if v == nil {
		return Null
	}
	return Int32(*v)
}

func PInt64(v *int64) Encoder {
	if v == nil {
		return Null
	}
	return Int64(*v)
}

func PUInt(v *uint) Encoder {
	if v == nil {
		return Null
	}
	return UInt(*v)
}

func PUInt8(v *uint8) Encoder {
	if v == nil {
		return Null
	}
	return UInt8(*v)
}

func PUInt16(v *uint16) Encoder {
	if v == nil {
		return Null
	}
	return UInt16(*v)
}

func PUInt32(v *uint32) Encoder {
	if v == nil {
		return Null
	}
	return UInt32(*v)
}

func PUInt64(v *uint64) Encoder {
	if v == nil {
		return Null
	}
	return UInt64(*v)
}

func PUIntPtr(v *uintptr) Encoder {
	if v == nil {
		return Null
	}
	return UIntPtr(*v)
}

func PFloat32(v *float32) Encoder {
	if v == nil {
		return Null
	}
	return Float32(*v)
}

func PFloat64(v *float64) Encoder {
	if v == nil {
		return Null
	}
	return Float64(*v)
}

func PString(v *string) Encoder {
	if v == nil {
		return Null
	}
	return String(*v)
}
