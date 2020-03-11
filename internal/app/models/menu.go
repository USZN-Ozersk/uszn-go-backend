package models

// Menu ...
type Menu struct {
	MenuID          int    `json:"menu_id"`
	MenuItem        string `json:"menu_item"`
	MenuParent      int    `json:"menu_parent"`
	CustomLink      bool   `json:"custom_link"`
	CustomLinkValue string `json:"custom_link_value"`
}
