package main

import (
	"fmt"
	"math/rand"
)

func main() {
	mars()
	lightspeed()
	shortcut()
	makeRandomNum()
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

func shortcut() {
	var weight = 149.0
	weight = weight * 0.3783
	weight *= 0.3783

	var age = 41
	age = age + 1
	age += 1
	age++
	// ++ageはサポートしない
}

func makeRandomNum() {
	var num = rand.Intn(10) + 1
	fmt.Println(num)
	num = rand.Intn(10) + 1
	fmt.Println(num)
}
