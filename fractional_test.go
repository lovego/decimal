package decimal

import "fmt"

func ExampleDecimal_FractionalBits() {
	fmt.Println(New(1, 0).FractionalBits())
	fmt.Println(New(12, 0).FractionalBits())
	fmt.Println(New(1, 1).FractionalBits())
	fmt.Println(New(1, 1).FractionalBits())
	fmt.Println(New(100, -1).FractionalBits())
	fmt.Println(New(100, -2).FractionalBits())
	// Output:
	// 0
	// 0
	// 0
	// 0
	// 0
	// 0
}

func ExampleDecimal_FractionalBits_2() {
	fmt.Println(New(123, -1).FractionalBits())
	fmt.Println(New(123, -2).FractionalBits())
	fmt.Println(New(123, -3).FractionalBits())
	fmt.Println(New(100, -3).FractionalBits())
	fmt.Println(New(100, -4).FractionalBits())
	fmt.Println(New(100, -5).FractionalBits())
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}
