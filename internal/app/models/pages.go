package models

// Page ...
type Page struct {
	PageID   int    `json:"page_id"`
	PageName string `json:"page_name"`
	PageText string `json:"page_text"`
	PageMenu int    `json:"page_menu"`
}

// HalfPage ...
type HalfPage struct {
	PageID   int    `json:"page_id"`
	PageMenu int    `json:"page_menu"`
	PageName string `json:"page_name"`
}
