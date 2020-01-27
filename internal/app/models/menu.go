package models

// Menu ...
type Menu struct {
	MenuID     string `json:"menu_id"`
	MenuItem   string `json:"menu_item"`
	MenuParent string `json:"menu_parent"`
}
