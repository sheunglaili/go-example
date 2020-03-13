package iteration

import "strings"

//Repeat repeat the character provided in 'count' times and return a concatanated string 
func Repeat(character string , count int) ( string) {
	return strings.Repeat(character,count)
}
