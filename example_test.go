package jsony_test

import (
	"fmt"

	"github.com/orsinium-labs/jsony"
)

func Example() {
	user := jsony.Object{
		jsony.Field{"name", jsony.String("aragorn")},
		jsony.Field{"age", jsony.Int(87)},
		jsony.Field{"admin", jsony.Bool(true)},
	}
	res := jsony.EncodeString(user)
	fmt.Println(res)
	//Output: {"name":"aragorn","age":87,"admin":true}
}

func ExampleBool() {
	user := jsony.Object{
		jsony.Field{"admin", jsony.Bool(true)},
	}
	res := jsony.EncodeString(user)
	fmt.Println(res)
	//Output: {"admin":true}
}

func ExampleInt() {
	user := jsony.Object{
		jsony.Field{"age", jsony.Int(123)},
	}
	res := jsony.EncodeString(user)
	fmt.Println(res)
	//Output: {"age":123}
}

func ExampleInt8() {
	user := jsony.Object{
		jsony.Field{"age", jsony.Int8(123)},
	}
	res := jsony.EncodeString(user)
	fmt.Println(res)
	//Output: {"age":123}
}

func ExampleInt16() {
	user := jsony.Object{
		jsony.Field{"age", jsony.Int16(123)},
	}
	res := jsony.EncodeString(user)
	fmt.Println(res)
	//Output: {"age":123}
}

func ExampleInt32() {
	user := jsony.Object{
		jsony.Field{"age", jsony.Int32(123)},
	}
	res := jsony.EncodeString(user)
	fmt.Println(res)
	//Output: {"age":123}
}

func ExampleInt64() {
	user := jsony.Object{
		jsony.Field{"age", jsony.Int64(123)},
	}
	res := jsony.EncodeString(user)
	fmt.Println(res)
	//Output: {"age":123}
}

func ExampleUInt() {
	user := jsony.Object{
		jsony.Field{"age", jsony.UInt(123)},
	}
	res := jsony.EncodeString(user)
	fmt.Println(res)
	//Output: {"age":123}
}

func ExampleUInt8() {
	user := jsony.Object{
		jsony.Field{"age", jsony.UInt8(123)},
	}
	res := jsony.EncodeString(user)
	fmt.Println(res)
	//Output: {"age":123}
}

func ExampleUInt16() {
	user := jsony.Object{
		jsony.Field{"age", jsony.UInt16(123)},
	}
	res := jsony.EncodeString(user)
	fmt.Println(res)
	//Output: {"age":123}
}

func ExampleUInt32() {
	user := jsony.Object{
		jsony.Field{"age", jsony.UInt32(123)},
	}
	res := jsony.EncodeString(user)
	fmt.Println(res)
	//Output: {"age":123}
}

func ExampleUInt64() {
	user := jsony.Object{
		jsony.Field{"age", jsony.UInt64(123)},
	}
	res := jsony.EncodeString(user)
	fmt.Println(res)
	//Output: {"age":123}
}

func ExampleUIntPtr() {
	user := jsony.Object{
		jsony.Field{"age", jsony.UIntPtr(123)},
	}
	res := jsony.EncodeString(user)
	fmt.Println(res)
	//Output: {"age":123}
}

func ExampleFloat32() {
	user := jsony.Object{
		jsony.Field{"weight", jsony.Float32(62.3)},
	}
	res := jsony.EncodeString(user)
	fmt.Println(res)
	//Output: {"weight":62.3}
}

func ExampleFloat64() {
	user := jsony.Object{
		jsony.Field{"weight", jsony.Float64(62.3)},
	}
	res := jsony.EncodeString(user)
	fmt.Println(res)
	//Output: {"weight":62.3}
}

func ExampleString() {
	user := jsony.Object{
		jsony.Field{"name", jsony.String("johny")},
	}
	res := jsony.EncodeString(user)
	fmt.Println(res)
	//Output: {"name":"johny"}
}

func ExampleSafeString() {
	user := jsony.Object{
		jsony.Field{"name", jsony.SafeString("johny")},
	}
	res := jsony.EncodeString(user)
	fmt.Println(res)
	//Output: {"name":"johny"}
}

func ExampleEncodeBytes() {
	res := jsony.EncodeBytes(jsony.Object{})
	fmt.Printf("%v\n", res)
	//Output: [123 125]
}

func ExampleAppendBytes() {
	buf := make([]byte, 0, 1024)
	res := jsony.AppendBytes(buf, jsony.Object{})
	fmt.Printf("%v\n", res)
	//Output: [123 125]
}

func ExampleEncodeString() {
	res := jsony.EncodeString(jsony.Object{})
	fmt.Printf("%v\n", res)
	//Output: {}
}

func ExampleObject() {
	user := jsony.Object{
		jsony.Field{"name", jsony.String("aragorn")},
		jsony.Field{"age", jsony.Int(87)},
		jsony.Field{"admin", jsony.Bool(true)},
	}
	res := jsony.EncodeString(user)
	fmt.Println(res)
	//Output: {"name":"aragorn","age":87,"admin":true}
}

func ExampleArray() {
	user := jsony.Array[jsony.Int]{
		jsony.Int(13), jsony.Int(14), jsony.Int(15),
	}
	res := jsony.EncodeString(user)
	fmt.Println(res)
	//Output: [13,14,15]
}

func ExampleMixedArray() {
	user := jsony.MixedArray{
		jsony.String("johny"), jsony.Int(14), jsony.Null,
	}
	res := jsony.EncodeString(user)
	fmt.Println(res)
	//Output: ["johny",14,null]
}
