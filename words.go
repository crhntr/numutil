package numutil

import "strings"

func ToWords(n int) string {
	switch {
	case n < 0:
		return "negative " + ToWords(-n)
	case n < 20:
		return upToNineteen[n]
	case n < 100:
		return strings.TrimSpace(tens[n/10-2] + " " + upToNineteen[n-(n/10*10)])
	case n < 1000:
		return strings.TrimSpace(upToNineteen[n/100] + " hundred " + ToWords(n-(n/100*100)))
	case n < 1000000:
		return strings.TrimSpace(ToWords(n/1000) + " thousand " + ToWords(n%1000))
	case n < 1000000000:
		return strings.TrimSpace(ToWords(n/1000000) + " billion " + ToWords(n%1000000))
	default:
		return strings.TrimSpace(ToWords(n/1000000000) + " trillion " + ToWords(n%1000000000))
	}
}

var upToNineteen = [...]string{
	"" /*"zero"*/, "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
}

var tens = [...]string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
