// There are 3 functions for generating JSON:
//
//   - [EncodeString] serializes data into a string.
//   - [EncodeBytes] serializes data into a slice of bytes.
//   - [AppendBytes] is like EncodeBytes but allows you to provide a buffer to use.
//
// Primitive types are constructed by:
//
//   - [Null]
//   - [Bool]
//   - [Int]
//   - [Int8]
//   - [Int16]
//   - [Int32]
//   - [Int64]
//   - [UInt]
//   - [UInt8]
//   - [UInt16]
//   - [UInt32]
//   - [UInt64]
//   - [UIntPtr]
//   - [Float32]
//   - [Float64]
//   - [String]
//   - [SafeString]
//
// Each of the above has a nullable version that serializes to null if the value is nil:
//
//   - [PBool]
//   - [PInt]
//   - [PInt8]
//   - [PInt16]
//   - [PInt32]
//   - [PInt64]
//   - [PUInt]
//   - [PUInt8]
//   - [PUInt16]
//   - [PUInt32]
//   - [PUInt64]
//   - [PUIntPtr]
//   - [PFloat32]
//   - [PFloat64]
//   - [PString]
//
// Collections are constructed by:
//
//   - [Object]
//   - [Map]
//   - [Array]
//   - [MixedArray]
//
// Lastly, there are 2 bridges between jsony and stdlib encoding/json:
//
//   - [Marshaller]
//   - [FromMarshaller]
package jsony
