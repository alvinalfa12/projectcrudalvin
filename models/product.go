package models

type Product struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Price      float64   `json:"price"`
	CategoryID int       `json:"category_id"`
	Category   *Category `json:"category,omitempty"`
}
