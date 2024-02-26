package main

import "fmt"

func main() {

	// possível cricar slice atraves do make
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	// mesmo atingindo a capacidade do slice o array aumenta sozinho
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))

}
