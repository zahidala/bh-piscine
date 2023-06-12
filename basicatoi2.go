package piscine

// package main

func BasicAtoi2(s string) int {
	numberArray := []rune(s)
	num := 0

	for i := 0; i < len(numberArray); i++ {
		if numberArray[i] < '0' || numberArray[i] > '9' {
			return 0
		}
		num = num*10 + (int(numberArray[i]) - 48)
	}
	return num
}

// func main() {
// 	fmt.Println(BasicAtoi2("12345"))
// 	fmt.Println(BasicAtoi2("0000000012345"))
// 	fmt.Println(BasicAtoi2("012 345"))
// 	fmt.Println(BasicAtoi2("Hello World!"))
// }
