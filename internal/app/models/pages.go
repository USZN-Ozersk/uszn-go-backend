package models

type Page struct {
	PageId   int    `json:"page_id"`
	PageName int    `json:"page_name"`
	PageText string `json:"page_text"`
	PageMenu int    `json:"page_menu"`
}
