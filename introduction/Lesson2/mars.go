package main

import (
	"fmt"
)

func main() {
	mars()
	tryFmt()
}

func mars() {
	const weight = 64 * 0.3783
	const age = 30 * 365 / 687
	fmt.Print("火星の表面では、とべの体重は", weight, "Kgに、")
	fmt.Println("年齢は", age, "歳になるでしょう。")
}

func tryFmt() {

}
