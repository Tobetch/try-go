package main

import (
	"fmt"
)

func main() {
	mars()
	lightspeed()
}

func mars() {
	const weight = 64 * 0.3783
	const age = 30 * 365 / 687
	fmt.Print("火星の表面では、とべの体重は", weight, "Kgに、")
	fmt.Println("年齢は", age, "歳になるでしょう。")
	fmt.Printf("火星の表面では、とべの体重はだいたい%dKgになります。\n", 24)
}

func lightspeed() {
	const lightSpeed = 299792 // km
	var distance = 56000000   // km
	fmt.Println(distance/lightSpeed, "秒")

	distance = 401000000
	fmt.Println(distance/lightSpeed, "秒")
}
