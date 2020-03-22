package main

import (
	"fmt"
	"github.com/go-ffmt/ffmt"
	"strconv"
	"strings"
)

var (
	sta string
	stb string
)

func main() {
	sta = "hello word"
	stb = "hello word"

	fmt.Print(sta)
	fmt.Print(stb)

	if strings.HasPrefix(sta, "he") {
		ffmt.Print("pre true")
	}

	if strings.HasSuffix(stb, "r") {
		ffmt.Print("has suf true")
	} else {
		ffmt.Print("has suf false")
	}

	ind := strings.Index(sta,"w")
	ffmt.Print(ind)

	hasCon := strings.Contains(sta, "aw")
	ffmt.Print(hasCon)

	cou := strings.Count(sta,"l")
	ffmt.Print(cou)

	up := strings.ToUpper(sta)
	ffmt.Print(up)

	lo := strings.ToLower(up)
	ffmt.Print(lo)

	rep := strings.Replace(sta,"l","EE", -1)
	ffmt.Print(rep)

	repa := strings.Repeat(sta, 3)
	fmt.Println(repa)

	stc := "   hello world   "

	trS := strings.TrimSpace(stc)
	ffmt.Print(trS)

	tr := strings.Trim(stc," ")
	ffmt.Println(tr)

	trL := strings.TrimLeft(stc, " ")
	fmt.Printf(trL)

	fmt.Print("\n")
	trr := strings.TrimRight(stc, "")
	ffmt.Print(trr)

	std := " a b c d e "
	ste := "|a|b|c|d|e|f|g"

	trSpli := strings.Split(std," ")
	ffmt.Print(trSpli)

	trSpli = strings.Split(ste,"|")
	ffmt.Print(trSpli)

	trSpli = strings.Split(ste," ")
	ffmt.Print(trSpli)

	stf := 2
	fmt.Print(stf)
	ffmt.Print(stf)
	fmt.Print("\n")
	trStrCo := strconv.Itoa(stf)
	fmt.Print(trStrCo)
	ffmt.Print(trStrCo)

	trStrCos,_ := strconv.Atoi(trStrCo)
	ffmt.Print(trStrCos)

	sum := 4 + 5
	ffmt.Print(sum)

	stPin := trStrCo + " test"
	ffmt.Print(stPin)



}

