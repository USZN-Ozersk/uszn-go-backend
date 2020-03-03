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

	rows, err := s.Db.Query("SELECT page_id, page_name, page_menu FROM pages")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var result models.HalfPage

		if err := rows.Scan(&result.PageID, &result.PageName, &result.PageMenu); err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return &results, nil

}

// InsertPage ...
func InsertPage(s *store.Store, name string, text string, menu int) error {
	_, err := s.Db.Exec("INSERT INTO pages (page_name, page_text, page_menu) VALUES ($1, $2, $3)", name, text, menu)
	if err != nil {
		return err
	}
	return nil
}

// DeletePage ...
func DeletePage(s *store.Store, id int) error {
	_, err := s.Db.Exec("DELETE FROM pages WHERE page_id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

// UpdatePage ...
func UpdatePage(s *store.Store, id int, name string, text string, menu int) error {
	_, err := s.Db.Exec("UPDATE pages SET page_name=$2, page_text=$3, page_menu=$4 WHERE page_id=$1", id, name, text, menu)
	if err != nil {
		return err
	}
	return nil
}
