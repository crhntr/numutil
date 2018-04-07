package numutil

func ToWords(n int) string {
	if n < 0 {
		return "negative " + ToWords(-n)
	}
	if n < 20 {
		return upToNineteen[n]
	}
	if n < 100 {
		lowOrder := upToNineteen[:10][n-(n/10*10)]
		if lowOrder != "" {
			lowOrder = " " + lowOrder
		}
		return tens[n/10-2] + lowOrder
	}
	if n < 1000 {
		return upToNineteen[:10][n/100] + " hundred " + ToWords(n-(n/100*100))
	}
	if n < 1000000 {
		return ToWords(n/1000) + " thosand " + ToWords(n%1000)
	}
	if n < 1000000000 {
		return ToWords(n/1000000) + " billion " + ToWords(n%1000000)
	}
	return ToWords(n/1000000000) + " trillion " + ToWords(n%1000000000)
}

var upToNineteen = [...]string{
	"", // zero
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
}

var tens = [...]string{
	"twenty",
	"thirty",
	"forty",
	"fifty",
	"sixty",
	"seventy",
	"eighty",
	"ninety",
}
