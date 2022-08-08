package hotel

import "fmt"

func main() {
	cus1 := Customer{
		Name:        "Thuan",
		Age:         25,
		IdentityNum: 2343442,
		TimeRent:    4,
		Skind:       "A",
	}
	fmt.Printf("so tien cu khac phai tra la:%v\n", cus1.Money())
}
