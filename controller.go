package hotel

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ListCustomers []Customer

func (c Customer) Input() {
	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Println("Nhap ten khach hang\n")
	name, _ := consoleReader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)
	c.Name = name

	fmt.Println("Nhap tuoi khach hang\n")
	age, _ := consoleReader.ReadString('\n')
	age = strings.Replace(age, "\n", "", -1)
	ageInt, _ := strconv.ParseInt(age, 10, 64)
	c.Age = ageInt

	fmt.Println("Nhap so cmnd cua khach hang\n")
	iden, _ := consoleReader.ReadString('\n')
	iden = strings.Replace(iden, "\n", "", -1)
	idenInt, _ := strconv.ParseInt(iden, 10, 64)
	c.IdentityNum = idenInt

	fmt.Println("Nhap so ngay luu tru cua khach hang\n")
	time, _ := consoleReader.ReadString('\n')
	time = strings.Replace(time, "\n", "", -1)
	timeInt, _ := strconv.ParseInt(time, 10, 64)
	c.TimeRent = timeInt

	fmt.Println("Nhap loai phong cua khach hang\n")
	ski, _ := consoleReader.ReadString('\n')
	if ski != A || ski != B || ski != C {
		fmt.Println("loai phong khong dung")
	}
	ski = strings.Replace(ski, "\n", "", -1)
	c.Skind = SkindType(ski)

	ListCustomers = append(ListCustomers, c)
}

func DeleteCus(id int64) {
	for i := 0; i < len(ListCustomers); i++ {
		cus := ListCustomers[i]
		if id == cus.IdentityNum {
			ListCustomers = append(ListCustomers[:i], ListCustomers[i+1:]...)
			break
		}
	}
}
func (c Customer) Money() int64 {
	pri := (SkindRoom[string(c.Skind)]).Price

	return c.TimeRent * pri
}
