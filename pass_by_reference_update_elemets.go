package main

import "fmt"

func updateElements(s *[]int) {
	fmt.Println("Inside updateElements function:")
	fmt.Printf("[pass by reference] slice: %v\n", *s)
	fmt.Printf("[pass by reference] memory address of slice header: %p, len: %d, cap: %d\n", s, len(*s), cap(*s))
	fmt.Printf("[pass by reference] memory address of the 0th element of the backing array: %p\n", *s)

	for i := 0; i < len(*s); i++ {
		(*s)[i] = i
	}

	fmt.Printf("[pass by reference] slice after update elements: %v\n\n", *s)
}

func main() {
	s := make([]int, 5, 5)
	fmt.Printf("[original] slice: %v\n", s)
	fmt.Printf("[original] memory address of slice header: %p, len: %d, cap: %d\n", &s, len(s), cap(s))
	fmt.Printf("[original] memory address of the 0th element of the backing array: %p\n\n", s)

	updateElements(&s)

	fmt.Printf("[original] slice after update elements: %v\n", s)
	fmt.Printf("[original] memory address of slice header: %p, len: %d, cap: %d\n", &s, len(s), cap(s))
	fmt.Printf("[original] memory address of the 0th element of the backing array: %p\n", s)
}

/* output
[original] slice: [0 0 0 0 0]
[original] memory address of slice header: 0xc00011a018, len: 5, cap: 5
[original] memory address of the 0th element of the backing array: 0xc000120060

Inside updateElements function:
[pass by reference] slice: [0 0 0 0 0]
[pass by reference] memory address of slice header: 0xc00011a018, len: 5, cap: 5
[pass by reference] memory address of the 0th element of the backing array: 0xc000120060
[pass by reference] slice after update elements: [0 1 2 3 4]

[original] slice after update elements: [0 1 2 3 4]
[original] memory address of slice header: 0xc00011a018, len: 5, cap: 5
[original] memory address of the 0th element of the backing array: 0xc000120060
*/
