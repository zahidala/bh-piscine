package piscine

// package main

func StringToIntSlice(str string) []int {
	var intArray []int
	runeArray := []rune(str)
	for i := 0; i <= len(runeArray)-1; i++ {
		intArray = append(intArray, int(runeArray[i]))
	}
	return intArray
}

// func main() {
// 	fmt.Println(StringToIntSlice("A quick brown fox jumps over the lazy dog"))
// 	fmt.Println(StringToIntSlice("Converted this string into an int"))
// 	fmt.Println(StringToIntSlice("hello THERE"))
// }
