package repos

import (
	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/models"
	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/store"
)

// GetMenus ...
func GetMenus(s *store.Store) ([]models.Menu, error) {
	var results []models.Menu
	rows, err := s.Db.Query("SELECT menu_id, menu_item, COALESCE(menu_parent, '') as menu_parent FROM menu")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var result models.Menu
<<<<<<< HEAD

		if err := rows.Scan(&result.MenuID, &result.MenuItem, &result.MenuParent); err != nil {
=======
		if err := rows.Scan(&result.MenuId, &result.MenuItem, &result.MenuParent); err != nil {
>>>>>>> 12072f2554d374c310f103e3c96a68212bf4dab6
			return nil, err
		}

		results = append(results, result)
	}
	return results, nil
}
