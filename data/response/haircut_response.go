package response

type HaircutResponse struct {
	Id         int     `json:"id"`
	NameClient string  `json:"client_name"`
	Price      float64 `json:"price"`
	Date       string  `json:"date"`
}
