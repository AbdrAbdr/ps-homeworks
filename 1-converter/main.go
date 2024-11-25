package main

import "fmt"

func main() {
	const usdEur float64 = 0.95
	const usdRub float64 = 103
	const eurRub float64 = usdRub / usdEur
	fmt.Println(eurRub)

}
