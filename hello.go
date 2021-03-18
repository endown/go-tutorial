package main

import (
	"fmt"
	// "strconv"
	// "strings"
)

type character struct {
	name        string
	yearOfBirth int
	birthPlace  string
}

type performance struct {
	character
	drama string
}

func (p performance) profile() {
	fmt.Println(p.name, "\n出生年:", p.yearOfBirth, "\n出身地:", p.birthPlace, "\n出演ドラマ:", p.drama)
}

func main() {
	p1 := performance{
		character: character{
			"新垣結衣",
			1988,
			"沖縄県",
		},
		drama: "逃げるは恥だが役に立つ",
	}
	p2 := performance{
		character: character{
			"本田翼",
			1992,
			"東京都",
		},
		drama: "ラジエーションハウス",
	}
	fmt.Println(p1)
	fmt.Println(p2)
	p1.profile()
	p2.profile()
}
