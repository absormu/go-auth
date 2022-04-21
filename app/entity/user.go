package entity

type User struct {
	ID            int64  `json:"user_id,omitempty"`
	Name          string `json:"name,omitempty"`
	Email         string `json:"email,omitempty"`
	UserContactID int    `json:"user_contact_id,omitempty"`
	RoleID        int    `json:"role_id,omitempty"`
	Password      string `json:"password,omitempty"`
	IsDeleted     int    `json:"is_deleted,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	CreatedBy     string `json:"created_by,omitempty"`
	ModifiedAt    string `json:"modified_at,omitempty"`
	ModifiedBy    string `json:"modified_by,omitempty"`
}

type UserMenuPrivilege struct {
	ID        int64  `json:"user_id,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Active    *bool  `json:"active"`
	CreatedAt string `json:"created_at,omitempty"`
}

type UserData struct {
	ID         int64                   `json:"user_id,omitempty"`
	Name       string                  `json:"name,omitempty"`
	Email      string                  `json:"email,omitempty"`
	CompanyID  int                     `json:"company_id,omitempty"`
	SellerID   int                     `json:"seller_id,omitempty"`
	RoleID     int                     `json:"role_id,omitempty"`
	Password   string                  `json:"password,omitempty"`
	Active     *bool                   `json:"active"`
	MenuID     int                     `json:"menu_id,omitempty"`
	MenuUser   []UserMenuPrivilegeData `json:"menu,omitempty"`
	IsDeleted  int                     `json:"is_deleted,omitempty"`
	CreatedAt  string                  `json:"created_at,omitempty"`
	CreatedBy  string                  `json:"created_by,omitempty"`
	ModifiedAt string                  `json:"modified_at,omitempty"`
	ModifiedBy string                  `json:"modified_by,omitempty"`
}

type UserMenuPrivilegeData struct {
	MenuID int64 `json:"menu_id,omitempty"`
	Active bool  `json:"active"`
}
