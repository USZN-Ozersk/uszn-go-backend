package models

// Menu ...
type Menu struct {
	MenuId     int    `json:"menu_id"`
	MenuItem   string `json:"menu_item"`
	MenuParent int    `json:"menu_parent"`
}
