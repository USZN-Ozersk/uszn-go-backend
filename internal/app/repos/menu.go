package repos

import (
	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/models"
	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/store"
)

// GetMenus ...
func GetMenus(s *store.Store) (*[]models.Menu, error) {
	var results []models.Menu
	rows, err := s.Db.Query("SELECT * FROM menu")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var result models.Menu

		if err := rows.Scan(&result.MenuID, &result.MenuItem, &result.MenuParent); err != nil {
			return nil, err
		}

		results = append(results, result)
	}
	return &results, nil
}

// InsertMenu ...
func InsertMenu(s *store.Store, name string, parent int) error {
	_, err := s.Db.Exec("INSERT INTO menu (menu_item, menu_parent) VALUES ($1, $2)", name, parent)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMenu ...
func DeleteMenu(s *store.Store, id int) error {
	_, err := s.Db.Exec("DELETE FROM menu WHERE menu_id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

// UpdateMenu ...
func UpdateMenu(s *store.Store, id int, name string, parent int) error {
	_, err := s.Db.Exec("UPDATE menu SET menu_item=$2, menu_parent=$3 WHERE menu_id=$1", id, name, parent)
	if err != nil {
		return err
	}
	return nil
}
