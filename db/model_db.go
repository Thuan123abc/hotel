package db

type CustomerModel struct {
	Name        string
	Age         int64
	IdentityNum int64 `gorm:" primaryKey "`
	TimeRentA   int64
	TimeRentB   int64
	TimeRentC   int64
}

type Room struct {
	Price      int64
	NumberRoom int64
}
