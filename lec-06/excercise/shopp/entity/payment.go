package entity

// import(
// 	"time"
// )

type Payment struct{
	ID int `json:"id"`
	Name string `json:"name"`
	// CreatedAt time.Time `json:"createAt"`
}


func (Payment) TableName() string {
    return "db_payment"
}
