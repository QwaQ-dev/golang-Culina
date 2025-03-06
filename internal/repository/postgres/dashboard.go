package postgres

import (
	"database/sql"
	"encoding/json"
	"log/slog"

	"github.com/qwaq-dev/culina/pkg/logger/sl"
	"github.com/qwaq-dev/culina/structures"
)

type PostgresDashboardRepository struct {
	DB *sql.DB
}

func (p *PostgresDashboardRepository) GetAllRecipes() (string, error) {
	return "", nil
}

func (p *PostgresDashboardRepository) CreateRecipe(recipe structures.Recipes, log *slog.Logger) (int, error) {
	var recipeId int

	ingredientsJSON, _ := json.Marshal(recipe.Ingredients)
	stepsJSON, _ := json.Marshal(recipe.Steps)
	imagesJSON, _ := json.Marshal(recipe.Imgs)
	filtersJSON, _ := json.Marshal(recipe.Filters)

	query := `INSERT INTO recipes (name, descr, diff, filters, ingredients, steps, author, imgs) 
          VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	log.Info("Serialized JSON before DB insert", slog.String("steps", string(stepsJSON)), slog.String("ingredients", string(ingredientsJSON)), slog.String("images", string(imagesJSON)))

	err := p.DB.QueryRow(query, recipe.Name, recipe.Descr, recipe.Diff, string(filtersJSON), string(ingredientsJSON), string(stepsJSON), recipe.AuthorID, string(imagesJSON)).Scan(&recipeId)
	if err != nil {
		log.Error("Error with inserting data", sl.Err(err))
		return 0, err
	}

	return recipeId, nil
}
