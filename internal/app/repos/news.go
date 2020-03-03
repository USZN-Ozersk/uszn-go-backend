package repos

import (
	"strconv"

	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/models"
	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/store"
)

// GetFirstNews ...
func GetFirstNews(r *store.Store, count string) (*[]models.News, error) {
	var results []models.News
	rows, err := r.Db.Query("SELECT * FROM news order by news_id desc fetch first $1 rows only", count)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var result models.News

		if err := rows.Scan(&result.NewsID, &result.NewsName, &result.NewsDate, &result.NewsText, &result.NewsImg); err != nil {
			return nil, err
		}
		result.CutNewsText(200)
		results = append(results, result)
	}

	return &results, nil
}

// GetOneNews ...
func GetOneNews(r *store.Store, id string) (*models.News, error) {
	var result models.News
	if err := r.Db.QueryRow("SELECT * FROM news WHERE news_id = $1", id).Scan(&result.NewsID, &result.NewsName, &result.NewsDate, &result.NewsText, &result.NewsImg); err != nil {
		return nil, err
	}

	return &result, nil
}

// CountNews ...
func CountNews(r *store.Store) (int, error) {
	var count int
	if err := r.Db.QueryRow("SELECT COUNT(*) FROM news").Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

// GetPageOfNews ...
func GetPageOfNews(r *store.Store, page string) (*[]models.News, error) {
	var results []models.News
	c, err := strconv.Atoi(page)
	rows, err := r.Db.Query("SELECT * FROM news ORDER BY news_id DESC LIMIT 10 OFFSET $1", c*10)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var result models.News

		if err := rows.Scan(&result.NewsID, &result.NewsName, &result.NewsDate, &result.NewsText, &result.NewsImg); err != nil {
			return nil, err
		}
		result.CutNewsText(150)
		results = append(results, result)
	}

	return &results, nil

}

// InsertNews ...
func InsertNews(s *store.Store, name string, text string, img string) error {
	_, err := s.Db.Exec("INSERT INTO news (news_name, news_text, news_img) VALUES ($1, $2, $3)", name, text, img)
	if err != nil {
		return err
	}
	return nil
}

// DeleteNews ...
func DeleteNews(s *store.Store, id int) error {
	_, err := s.Db.Exec("DELETE FROM news WHERE news_id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

// UpdateNews ...
func UpdateNews(s *store.Store, id int, name string, text string, img string) error {
	_, err := s.Db.Exec("UPDATE news SET news_name=$2, news_text=$3, news_img=$4 WHERE news_id=$1", id, name, text, img)
	if err != nil {
		return err
	}
	return nil
}
