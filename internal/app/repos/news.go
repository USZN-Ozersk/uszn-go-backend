package repos

import (
	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/models"
	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/store"
)

// GetNews ...
func GetNews(r *store.Store) (*[]models.News, error) {
	var results []models.News
	rows, err := r.Db.Query("SELECT * FROM news")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var result models.News

		if err := rows.Scan(&result.NewsID, &result.NewsName, &result.NewsDate, &result.NewsText); err != nil {
			return nil, err
		}
		result.CutNewsText(150)
		results = append(results, result)
	}

	return &results, nil
}

// GetOneNews ...
func GetOneNews(r *store.Store, id string) (*models.News, error) {
	var result models.News
	if err := r.Db.QueryRow("SELECT * FROM news WHERE news_id = $1", id).Scan(&result.NewsID, &result.NewsName, &result.NewsDate, &result.NewsText); err != nil {
		return nil, err
	}

	return &result, nil
}
