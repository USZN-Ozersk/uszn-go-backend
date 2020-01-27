package models

// Page ...
type Page struct {
	PageID   string `json:"page_id"`
	PageName string `json:"page_name"`
	PageText string `json:"page_text"`
	PageMenu string `json:"page_menu"`
}
