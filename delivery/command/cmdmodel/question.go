package cmdmodel

type Question struct {
	ID       uint64   `json:"id"`
	Question string   `json:"question"`
	Answer   []Answer `json:"answers"`
	Required bool     `json:"required"`
}
