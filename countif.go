package piscine

// package main

// func IsNumeric(s string) bool {
// 	for i := range s {
// 		if s[i] < '0' || s[i] > '9' {
// 			return false
// 		}
// 	}
// 	return true
// }

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, v := range tab {
		if f(v) {
			count = count + 1
		}
	}
	return count
}

// func main() {
// 	tab1 := []string{"Hello", "how", "are", "you"}
// 	tab2 := []string{"This", "1", "is", "4", "you"}
// 	answer1 := CountIf(IsNumeric, tab1)
// 	answer2 := CountIf(IsNumeric, tab2)
// 	fmt.Println(answer1)
// 	fmt.Println(answer2)
// }
