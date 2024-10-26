// package iteration

// const repeatedCount = 5

// func Repeat(character string) string {
// 	var repeated string
// 	for i := 0; i < repeatedCount; i++ {
// 		repeated += character
// 	}
// 	return repeated
// }

package iteration

func Repeat(character string, count int) string {
	repeated := ""
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}
