package main

import "fmt"

/*
	TYPES CH

		Boolean: bool represents a boolean value, which can be either true or false.

		Numeric types: Go provides a variety of numeric types, including:

		int8, int16, int32, int64, int: signed integers of various sizes
		uint8, uint16, uint32, uint64, uint: unsigned integers of various sizes
		float32, float64: floating-point numbers with single and double precision
		complex64, complex128: complex numbers with single and double precision
		String: string represents a string of characters.

		Array: An array is a fixed-size sequence of elements of the same type. The size of an array is specified when it is declared and cannot be changed later.

		Slice: A slice is a dynamic array that can grow or shrink as needed. It is represented by a pointer to an underlying array, a length, and a capacity.

		Map: A map is an unordered collection of key-value pairs. The keys are unique and the values can be of any type.

		Struct: A struct is a composite data type that groups together zero or more values with different types under a single name.

		Pointer: A pointer is a variable that stores the memory address of another variable.
*/

func main() {
	// boolean types
	fmt.Println("- Booleans -")
	var isTrue bool = true
	var isFalse bool = false
	fmt.Println("true -> ", isTrue)
	fmt.Println("false -> ", isFalse)
	fmt.Println("true and false -> ", isTrue && isFalse)
	fmt.Println("true or fakse -> ", isTrue || isFalse)
	fmt.Println("not true -> ", !isTrue)

	// Integer types
	fmt.Println("\n- Numerics -")
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	var i int = 2147483647 // platform dependent

	// Unsigned integer types
	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615
	var u uint = 4294967295 // platform dependent

	// Floating-point types
	var f32 float32 = 3.14159
	var f64 float64 = 3.141592653589793

	// Complex types
	var c64 complex64 = 1 + 2i
	var c128 complex128 = 3 + 4i

	// Arithmetic operations
	sum := i32 + int32(u32)
	diff := f64 - float64(i)
	product := f32 * float32(u16)
	quotient := f64 / float64(f32)

	// Type conversions
	i = int(u)
	f32 = float32(i32)
	c128 = complex128(c64)

	// Output
	fmt.Println("Integer types -> ", i8, i16, i32, i64, i, u8, u16, u32, u64, u)
	fmt.Println("Floating-point types -> ", f32, f64)
	fmt.Println("Complex types -> ", c64, c128)
	fmt.Println("Arithmetic operations -> ", sum, diff, product, quotient)

	// Strings
	fmt.Println("\n- Strings -")

	// -- Declaring #1
	var name1 string = "Alex"

	// -- Declaring #2
	name2 := "Samuel"
	println("name2 -> ", name2, '\n')

	// -- Concatenating
	greeting := "Hi, " + name1 + "!\n"
	println("greetings -> ", greeting)

	// -- Length
	length := len(name1)
	fmt.Printf("Length of 'name' -> %d\n", length)

	// -- iterating over a string
	for i, ch := range name1 {
		fmt.Printf("iterating %d %s", i, string(ch))
	}
	fmt.Printf("\n")

	// Arrays
	fmt.Println("\n- Arrays -")
	// -- Declaring
	var arr1 [5]int
	var arr2 = [5]int{1, 2, 3, 4, 5}
	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("Declaring arrays -> ", arr1, arr2, arr3)

	// -- Accessing elements of an array
	var thirdElArr1 = arr1[2]
	fmt.Println("arr1[2] -> ", thirdElArr1)
	

}
