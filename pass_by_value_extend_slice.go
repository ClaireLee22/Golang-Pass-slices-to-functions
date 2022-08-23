package main

import "fmt"

func extendSlice(s []int) {
	fmt.Println("Inside extendSlice function:")
	fmt.Printf("[copied] slice: %v\n", s)
	fmt.Printf("[copied] memory address of slice header: %p, len: %d, cap: %d\n", &s, len(s), cap(s))
	fmt.Printf("[copied] memory address of the 0th element of the backing array: %p\n\n", s)

	s = append(s, len(s))

	fmt.Printf("[copied] memory address of slice header after extending: %p, len: %d, cap: %d\n", &s, len(s), cap(s))
	fmt.Printf("[copied] memory address of the 0th element of the backing array: %p\n\n", s)
	fmt.Printf("[copied] slice after extending: %v\n\n", s)
}

func main() {
	s := []int{0, 1, 2, 3, 4}
	fmt.Printf("[original] slice: %v\n", s)
	fmt.Printf("[original] memory address of slice header: %p, len: %d, cap: %d\n", &s, len(s), cap(s))
	fmt.Printf("[original] memory address of the 0th element of the backing array: %p\n\n", s)

	extendSlice(s)

	fmt.Printf("[original] slice after extending: %v\n", s)
	fmt.Printf("[original] memory address of slice: %p, len: %d, cap: %d\n", &s, len(s), cap(s))
	fmt.Printf("[original] memory address of the 0th element of the backing array: %p\n", s)
}

/* output
[original] slice: [0 1 2 3 4]
[original] memory address of slice header: 0xc0000a4018, len: 5, cap: 5
[original] memory address of the 0th element of the backing array: 0xc0000aa060

Inside extendSlice function:
[copied] slice: [0 1 2 3 4]
[copied] memory address of slice header: 0xc0000a4060, len: 5, cap: 5
[copied] memory address of the 0th element of the backing array: 0xc0000aa060

[copied] memory address of slice header after extending: 0xc0000a4060, len: 6, cap: 10
[copied] memory address of the 0th element of the backing array: 0xc0000b8000

[copied] slice after extending: [0 1 2 3 4 5]

[original] slice after extending: [0 1 2 3 4]
[original] memory address of slice: 0xc0000a4018, len: 5, cap: 5
[original] memory address of the 0th element of the backing array: 0xc0000aa060
*/
