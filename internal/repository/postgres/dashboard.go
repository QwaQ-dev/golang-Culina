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

func (p *PostgresDashboardRepository) CreateRecipe(recipe structures.Recipes, log *slog.Logger) (int, error) {
	var recipeId int

	ingredientsJSON, _ := json.Marshal(recipe.Ingredients)
	stepsJSON, _ := json.Marshal(recipe.Steps)
	imagesJSON, _ := json.Marshal(recipe.Imgs)
	filtersJSON, _ := json.Marshal(recipe.Filters)

	query := `INSERT INTO recipes (name, descr, diff, filters, ingredients, steps, authorid, imgs) 
          VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	err := p.DB.QueryRow(query, recipe.Name, recipe.Descr, recipe.Diff, string(filtersJSON), string(ingredientsJSON), string(stepsJSON), recipe.AuthorID, string(imagesJSON)).Scan(&recipeId)
	if err != nil {
		log.Error("Error with inserting data", sl.Err(err))
		return 0, err
	}

	recipe.Id = recipeId

	log.Info("Recipe was upload to db", slog.Any("recipe", recipe))

	return recipeId, nil
}

func (p *PostgresDashboardRepository) GetAllRecipes(log *slog.Logger) ([]structures.Recipes, error) {
	var recipes []structures.Recipes

	rows, err := p.DB.Query("SELECT * FROM recipes")
	if err != nil {
		log.Error("Error with selecting recipes", sl.Err(err))
		return nil, nil
	}
	defer rows.Close()

	for rows.Next() {
		var recipe structures.Recipes
		var filtersJSON, imgsJSON, ingredientsJSON, stepsJSON []byte

		err := rows.Scan(&recipe.Id, &recipe.Name, &recipe.Descr, &recipe.Diff, &filtersJSON, &imgsJSON, &recipe.AuthorID, &ingredientsJSON, &stepsJSON)
		if err != nil {
			log.Error("Error scanning row", sl.Err(err))
			continue
		}

		json.Unmarshal(filtersJSON, &recipe.Filters)
		json.Unmarshal(imgsJSON, &recipe.Imgs)
		json.Unmarshal(ingredientsJSON, &recipe.Ingredients)
		json.Unmarshal(stepsJSON, &recipe.Steps)

		recipes = append(recipes, recipe)
	}

	if err := rows.Err(); err != nil {
		log.Error("Error after iterating rows", sl.Err(err))
		return nil, err
	}

	return recipes, nil
}
