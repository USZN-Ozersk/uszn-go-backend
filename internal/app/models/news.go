package models

type News struct {
	NewsId   int    `json:"news_id"`
	NewsName string `json:"news_name"`
	//NewsDate time.Date `json:"news_date"`
	NewsText string `json:"news_text"`
}
