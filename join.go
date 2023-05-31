package piscine

// package main

func Join(strs []string, sep string) string {
	var words string
	for i := 0; i <= len(strs)-1; i++ {
		words = words + strs[i] + sep
	}
	return words[:len(words)-1]
}

// func main() {
// 	toConcat := []string{"Hello!", " How", " are", " you?"}
// 	fmt.Println(Join(toConcat, ":"))
// }
