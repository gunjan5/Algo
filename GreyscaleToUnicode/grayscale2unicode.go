package main

import (
	"fmt"
	"os"
	_ "unicode/utf8"
)

var (
	uni = map[int]string{

		//0:  "\u2588",
		0:  "\u2588",
		1:  "\u2589",
		2:  "\u258A",
		3:  "\u258B",
		4:  "\u2593",
		5:  "\u25A9",
		6:  "\u25A6",
		7:  "\u25A4",
		8:  "\u25A7",
		9:  "\u25A8",
		10: "\u25C9",
		11: "\u25A3",
		12: "\u25C8",
		13: "\u2592",
		14: "\u2591",
		15: "$",
		16: "@",
		17: "B",
		18: "%",
		19: "8",
		20: "&",
		21: "W",
		22: "M",
		23: "#",
		24: "*",
		25: "o",
		26: "a",
		27: "h",
		28: "k",
		29: "b",
		30: "d",
		31: "p",
		32: "q",
		33: "w",
		34: "m",
		35: "Z",
		36: "O",
		37: "0",
		38: "Q",
		39: "L",
		40: "C",
		41: "J",
		42: "U",
		43: "Y",
		44: "X",
		45: "z",
		46: "c",
		47: "v",
		48: "u",
		49: "n",
		50: "x",
		51: "r",
		52: "j",
		53: "f",
		54: "t",
		55: "/",
		56: "\\",
		57: "|",
		58: "(",
		59: ")",
		60: "1",
		61: "{",
		62: "}",
		63: "[",
		64: "]",
		65: "?",
		66: "-",
		67: "_",
		68: "+",
		69: "~",
		70: "<",
		71: ">",
		72: "i",
		73: "!",
		74: "l",
		75: "I",
		76: ";",
		77: ":",
		78: ",",
		79: "\"",
		80: "^",
		81: "`",
		82: "'",
		83: ".",
		84: " ",
	}
)

func main() {

	input := 26
	output := Itou(input)

	fmt.Println("RESULT: ", output)

}

func Itou(i int) string {
	if i < 0 || i > 84 {
		fmt.Println("ERROR: Input must be between 0 and 84")
		os.Exit(-1)
	}
	return uni[i]

}
