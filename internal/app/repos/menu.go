package repos

import (
	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/models"
	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/store"
)

func GetMenus(s *store.Store) ([]models.Menu, error) {
	var results []models.Menu
	rows, err := s.Db.Query("SELECT * FROM menu")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			menuId     int
			menuItem   string
			menuParent int
		)
		if err := rows.Scan(&menuId, &menuItem, &menuParent); err != nil {
			return nil, err
		}

		results = append(results, models.Menu{MenuId: menuId, MenuItem: menuItem, MenuParent: menuParent})
	}
	return results, nil
}
