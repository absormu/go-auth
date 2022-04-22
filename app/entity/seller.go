package entity

type Seller struct {
	ID         int64    `json:"id"`
	Name       string   `json:"name"`
	City       CityData `json:"city_id"`
	Email      string   `json:"email"`
	Telephone  *string  `json:"telephone"`
	Password   string   `json:"password,omitempty"`
	Address    *string  `json:"address"`
	User       User     `json:"user"`
	Active     int      `json:"active"`
	IsDeleted  int      `json:"is_deleted"`
	CreatedAt  string   `json:"created_at"`
	CreatedBy  string   `json:"created_by"`
	ModifiedAt *string  `json:"modified_at"`
	ModifiedBy *string  `json:"modified_by"`
}

type SellerData struct {
	ID         int64    `json:"id"`
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	Password   string   `json:"password,omitempty"`
	Telephone  *string  `json:"telephone"`
	Address    *string  `json:"address"`
	City       CityData `json:"city"`
	Active     bool     `json:"active"`
	IsDeleted  int      `json:"is_deleted"`
	CreatedAt  string   `json:"created_at"`
	CreatedBy  string   `json:"created_by"`
	ModifiedAt *string  `json:"modified_at"`
	ModifiedBy *string  `json:"modified_by"`
}
