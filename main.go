package testingreviewdog

import "fmt"

// golangci-lint run --out-format line-number

func main() {
	RandomStuff()
	RandomWackyStuff()
}

// RandomStuff comment
func RandomStuff() {
	fmt.Println("Print from RandomStuff")
}

func randomWackyStuff() bool {
	return true
}

// RandomWackyStuffV2 comment
func RandomWackyStuffV2() {
	fmt.Println("Print from RandomWackyStuff")
}

func RandomWackyStuffV3() {
	fmt.Println("Print from RandomWackyStuff")
}

// RandomWackyStuff comment
func RandomWackyStuff() {
	fmt.Println("Print from RandomWackyStuff")
}
