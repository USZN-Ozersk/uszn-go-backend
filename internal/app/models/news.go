package models

// News ...
type News struct {
	NewsID   int    `json:"news_id"`
	NewsName string `json:"news_name"`
	NewsDate string `json:"news_date"`
	NewsText string `json:"news_text"`
}
