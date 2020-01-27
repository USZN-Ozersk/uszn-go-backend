package models

// Menu ...
type Menu struct {
	MenuID     int    `json:"menu_id"`
	MenuItem   string `json:"menu_item"`
	MenuParent int    `json:"menu_parent"`
	MenuPage   int    `json:"menu_page"`
}
