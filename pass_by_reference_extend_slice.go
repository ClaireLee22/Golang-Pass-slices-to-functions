package main

import "fmt"

func extendSlice(s *[]int) {
	fmt.Println("Inside extendSlice function:")
	fmt.Printf("[pass by reference] slice: %v\n", *s)
	fmt.Printf("[pass by reference] memory address of slice header: %p, len: %d, cap: %d\n", s, len(*s), cap(*s))
	fmt.Printf("[pass by reference] memory address of the 0th element of the backing array: %p\n\n", *s)

	*s = append(*s, len(*s))

	fmt.Printf("[pass by reference] memory address of slice header after extending: %p, len: %d, cap: %d\n", s, len(*s), cap(*s))
	fmt.Printf("[pass by reference] memory address of the 0th element of the backing array: %p\n\n", *s)
	fmt.Printf("[pass by reference] slice after extending: %v\n\n", *s)
}

func main() {
	s := []int{0, 1, 2, 3, 4}
	fmt.Printf("[original] slice: %v\n", s)
	fmt.Printf("[original] memory address of slice header: %p, len: %d, cap: %d\n", &s, len(s), cap(s))
	fmt.Printf("[original] memory address of the 0th element of the backing array: %p\n\n", s)

	extendSlice(&s)

	fmt.Printf("[original] slice after extending: %v\n", s)
	fmt.Printf("[original] memory address of slice: %p, len: %d, cap: %d\n", &s, len(s), cap(s))
	fmt.Printf("[original] memory address of the 0th element of the backing array: %p\n", s)
}

/* output
[original] slice: [0 1 2 3 4]
[original] memory address of slice header: 0xc00000c030, len: 5, cap: 5
[original] memory address of the 0th element of the backing array: 0xc000014150

Inside extendSlice function:
[pass by reference] slice: [0 1 2 3 4]
[pass by reference] memory address of slice header: 0xc00000c030, len: 5, cap: 5
[pass by reference] memory address of the 0th element of the backing array: 0xc000014150

[pass by reference] memory address of slice header after extending: 0xc00000c030, len: 6, cap: 10
[pass by reference] memory address of the 0th element of the backing array: 0xc00001a050

[pass by reference] slice after extending: [0 1 2 3 4 5]

[original] slice after extending: [0 1 2 3 4 5]
[original] memory address of slice: 0xc00000c030, len: 6, cap: 10
[original] memory address of the 0th element of the backing array: 0xc00001a050
*/
