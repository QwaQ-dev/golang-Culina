package postgres

import (
	"database/sql"
)

type PostgresDashboardRepository struct {
	DB *sql.DB
}

func (r *PostgresDashboardRepository) GetAllRecipes() (string, error) {
	return "", nil
}
