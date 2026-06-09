package main
import "fmt"

type Fees struct {
	buyFee int
	sellFee int
	transferFee int
}

func main() {
	var x Fees = Fees{2, 2, 2}
	y := Fees{3, 3, 3}

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(y.buyFee, y.sellFee, y.transferFee)
}

// go run app.go
