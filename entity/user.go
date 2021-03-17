package entity

type User struct{
	ID string `json:"id"`
	Name string `json:"name"`
	// Addr string `json:"addr"`
}

type Usertb struct{
	// ID string `gorm:`
	Name string `json:"name"`
	// Addr string `json:"addr"`
}

type Owner struct{
	Id int `json:"id"`
	Name string
	Books []Book
}

type Book struct{
	ID int
	BName string
	OwnerID uint
}