package repos

import (
	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/models"
	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/store"
)

func GetMenus(s *store.Store) ([]models.Menu, err) {
	var result []models.Menu
	if rows, err := s.Db.Query("SELECT * FROM menus"); err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		//result
	}
}
