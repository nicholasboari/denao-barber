package request

type CreateHaircutRequest struct {
	NameClient string  `json:"client_name"`
	Price      float64 `json:"price"`
	Date       string  `json:"date"`
}
