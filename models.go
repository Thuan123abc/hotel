package hotel

type Customer struct {
	Name        string
	Age         int64
	IdentityNum int64
	TimeRent    int64
	Skind       SkindType
}

type SkindType string

const (
	A string = "A"
	B string = "B"
	C string = "C"
)

type Room struct {
	Price      int64
	NumberRoom int64
}

var SkindRoom = map[string]Room{
	A: Room{
		500,
		10,
	},
	B: Room{
		300,
		9,
	},
	C: Room{
		100,
		4,
	},
}
