// hello.go
package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {

	s := unicode.IsUpper('a')
	fmt.Print(s)
	fmt.Println(IsHexDigit('8'), IsHexDigit('x'), IsHexDigit('X'),
		IsHexDigit('b'), IsHexDigit('B'))

}

func IsHexDigit(char rune) bool {
	return unicode.Is(unicode.ASCII_Hex_Digit, char)
}

func Humanize(amount float64, width, decimals int,
	pad, separator rune) string {
	dollars, cents := math.Modf(amount)
	whole := fmt.Sprintf("%+.0f", dollars)[1:] // Strip "±"
	fraction := ""
	if decimals > 0 {
		fraction = fmt.Sprintf("%+.*f", decimals, cents)[2:] // Strip "±0"
	}
	sep := string(separator)
	for i := len(whole) - 3; i > 0; i -= 3 {
		whole = whole[:i] + sep + whole[i:]
	}
	if amount < 0.0 {
		whole = "-" + whole
	}
	number := whole + fraction
	gap := width - utf8.RuneCountInString(number)
	if gap > 0 {
		return strings.Repeat(string(pad), gap) + number
	}
	return number
}

func test() {

	line := "rå tørt\u2028vær"
	i := strings.IndexFunc(line, unicode.IsSpace) // i == 3
	firstWord := line[:i]
	j := strings.LastIndexFunc(line, unicode.IsSpace) // j == 9
	_, size := utf8.DecodeRuneInString(line[j:])      // size == 3
	lastWord := line[j+size:]                         // j + size == 12
	fmt.Println(firstWord, lastWord)                  // Prints: rå vær

	who := "world"
	if len(os.Args) > 1 {
		who = strings.Join(os.Args[1:], " ")
	}

	fmt.Println("Hello ", who)
	æs := ""
	for _, char := range []rune{'æ', 0xE6, 0346, 230, '\xE6', '\u00E6'} {
		fmt.Printf("[0x%X '%c'] ", char, char)
		æs += string(char)
	}
	fmt.Println(æs)

	phrase := "vått og tørt王"
	fmt.Printf("string: \"%s\"\n", phrase)
	fmt.Println("index rune char bytes", len([]rune(phrase)))
	for index, char := range phrase {
		fmt.Println(char)
		fmt.Printf("%-2d %U '%c' % X\n",
			index, char, char,
			[]byte(string(char)))
	}

	fmt.Printf("|%b|%9b|%-9b|%09b|% 9b|\n", 37, 37, 37, 37, 37)
	fmt.Printf("|%o|%#o|%# 8o|%#+ 8o|%+08o|\n", 41, 41, 41, 41, -41)
	//	i := 3931
	//	fmt.Printf("|%x|%X|%8x|%08x|%#04X|0x%04X|\n", i, i, i, i, i, i)
	fmt.Printf("%d %#04x %U '%c'\n", 0x3A6, 934, '\u03A6', '\U000003A6')

	s := "Dare to be naïve"
	fmt.Printf("|%22s|%-22s|%10s|\n", s, s, s)

	//	i := strings.Index(s, "n")
	//	fmt.Printf("|%.10s|%.*s|%-22.10s|%s|\n", s, i, s, s, s)

	slogan := "End Óréttlæti♥"
	fmt.Printf("%s\n%q\n%+q\n%#q\n", slogan, slogan, slogan, slogan)

	//for _, x := range []float64{-.258, 7194.84, -60897162.0218, 1.500089e-8} {
	//	fmt.Printf("|%20.5e|%20.5f|%s|\n", x, x, Humanize(x, 20, 5, '*', ','))
	//}

	names := " Antônio\tAndré\tFriedrich\t\t\tJean\t\tÉlisabeth\tIsabella \t"
	names = strings.Replace(names, "\t", " ", -1)
	fmt.Printf("|%s|\n", names)
	fmt.Println(strings.Join(strings.Fields(names), " "))

	asciiOnly := func(char rune) rune {
		if char > 127 {
			return '?'
		}
		return char
	}

	fmt.Println(strings.Map(asciiOnly, names))
}
