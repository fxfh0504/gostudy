package main

import (
	"golang.org/x/tour/wc"
	"unicode"
	"strings"
)

func WordCount(s string) map[string]int {
	f := func(c rune) bool {
		return unicode.IsSpace(c)
	}
	var result=make( map[string]int)
	fieldsFunc := strings.FieldsFunc(s, f)
	for i:=0; i<len(fieldsFunc);i++ {
		key := fieldsFunc[i]
		println(result[key],key,i)
		result[key]+=1
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
