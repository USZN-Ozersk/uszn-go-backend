package repos

import (
	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/models"
	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/store"
)

// GetPage ...
func GetPage(s *store.Store, id string) (*models.Page, error) {
	var result models.Page

	if err := s.Db.QueryRow("SELECT * FROM pages WHERE page_menu = $1", id).Scan(&result.PageID, &result.PageName, &result.PageText, &result.PageMenu); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetAllPages ...
func GetAllPages(s *store.Store) (*[]models.HalfPage, error) {
	var results []models.HalfPage

	rows, err := s.Db.Query("SELECT page_id, page_name FROM pages")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var result models.HalfPage

		if err := rows.Scan(&result.PageID, &result.PageName); err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return &results, nil

}
