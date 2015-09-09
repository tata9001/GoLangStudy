// bigdigts
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	length := len(os.Args)
	if length == 1 {
		fmt.Printf("usage: %s <whole-number>\n", filepath.Base(os.Args[0]))
		os.Exit(0)
	}
	if length > 1 && ("--help" == os.Args[1] || "-h" == os.Args[1]) {
		fmt.Println("usage: bigdigits [-b|--bar] <whole-number>")
		fmt.Println("-b --bar draw an underbar and an overbar")
		os.Exit(0)
	}
	needbar := length > 1 && ("--bar" == os.Args[1] || "-b" == os.Args[1])

	str := os.Args[1]
	if needbar {
		if length == 2 {
			fmt.Printf("usage: %s [-b|--bar] <whole-number>\n", filepath.Base(os.Args[0]))
			os.Exit(0)
		}
		str = os.Args[2]
	}

	for row := range bigDigits[0] {
		line := ""
		for column := range str {
			digit := str[column] - '0'
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + " "
			} else {
				log.Fatal(" 无效的输入")
			}
		}
		if needbar && row == 0{
			fmt.Println(strings.Repeat("*",len(line)))
			
		}
		
		fmt.Println(line)
		if needbar && row + 1 == len(bigDigits[0]){
			fmt.Println(strings.Repeat("*",len(line)))
		}
	}

	

}


var bigDigits = [][]string{
	{"  000  ",
		" 0   0 ",
		"0     0",
		"0     0",
		"0     0",
		" 0   0 ",
		"  000  "},
	{" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
	{" 222 ", "2   2", "   2 ", "  2  ", " 2   ", "2    ", "22222"},
	{" 333 ", "3   3", "    3", "  33 ", "    3", "3   3", " 333 "},
	{"   4  ", "  44  ", " 4 4  ", "4  4  ", "444444", "   4  ",
		"   4  "},
	{"55555", "5    ", "5    ", " 555 ", "    5", "5   5", " 555 "},
	{" 666 ", "6    ", "6    ", "6666 ", "6   6", "6   6", " 666 "},
	{"77777", "    7", "   7 ", "  7  ", " 7   ", "7    ", "7    "},
	{" 888 ", "8   8", "8   8", " 888 ", "8   8", "8   8", " 888 "},
	{" 9999", "9   9", "9   9", " 9999", "    9", "    9", "    9"},
}
