package model

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Payload struct for task 2
type CreatePersonRequest struct {
	Name    string `json:"name" binding:"required"`
	Age     int    `json:"age" binding:"required"`
	Number  string `json:"phone_number" binding:"required"`
	City    string `json:"city" binding:"required"`
	State   string `json:"state" binding:"required"`
	Street1 string `json:"street1" binding:"required"`
	Street2 string `json:"street2"`
	ZipCode string `json:"zip_code" binding:"required"`
}
