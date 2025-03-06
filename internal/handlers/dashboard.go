package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/qwaq-dev/culina/internal/repository"
	"github.com/qwaq-dev/culina/structures"
)

type DashboardHandler struct {
	repo repository.DashboardRepository
	log  *slog.Logger
}

func NewDashboardHandler(repo repository.DashboardRepository, log *slog.Logger) *DashboardHandler {
	return &DashboardHandler{repo: repo, log: log}
}

/*
FORM-DATA{
	"name":"",
	"descr":"",
	"diff":"",
	"filters":["", ""],
	"imgs":"{auto}",
	"authorID":"from token",
	"ingredients":{"first":"ingr", },
	"steps":{"first":"step"},
}
*/

func (h *DashboardHandler) CreateRecipe(c *fiber.Ctx) error {
	name := c.FormValue("name")
	descr := c.FormValue("descr")
	diff := c.FormValue("diff")
	authorId, _ := strconv.Atoi(c.FormValue("authorID"))

	var filters []string
	ingredients := make(map[string]string)
	steps := make(map[string]string)

	if err := json.Unmarshal([]byte(c.FormValue("filters")), &filters); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Ivalid filers format",
		})
	}

	if err := json.Unmarshal([]byte(c.FormValue("ingredients")), &ingredients); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Ivalid ingredients format",
		})
	}

	if err := json.Unmarshal([]byte(c.FormValue("steps")), &steps); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Ivalid ingredients format",
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid form data",
		})
	}

	files, ok := form.File["images"]
	if !ok || len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No images uploaded",
		})
	}

	imgs := make(map[string]string)

	if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
		err := os.Mkdir("./uploads", os.ModePerm)
		if err != nil {
			log.Println("Ошибка при создании папки:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create uploads directory",
			})
		}
	}

	for i, file := range files {
		if i+1 >= 4 {
			break
		}

		filename := fmt.Sprintf("./uploads/%d_%s", time.Now().Unix(), file.Filename)

		// Сохраняем файл
		if err := c.SaveFile(file, filename); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to upload images",
			})
		}

		imgs[strconv.Itoa(i+1)] = filename
	}

	recipe := structures.Recipes{
		Name:        name,
		Descr:       descr,
		Diff:        diff,
		Filters:     filters,
		Imgs:        imgs,
		AuthorID:    authorId,
		Ingredients: ingredients,
		Steps:       steps,
	}

	id, err := h.repo.CreateRecipe(recipe, h.log)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to save recipe"})
	}

	recipe.Id = id
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Recipe was upload successfully",
		"recipe":  recipe,
	})
}

func (h *DashboardHandler) AddReview(c *fiber.Ctx) error {
	return nil
}

func (h *DashboardHandler) Filter(c *fiber.Ctx) error {
	return nil
}

func (h *DashboardHandler) SortBy(c *fiber.Ctx) error {
	return nil
}

func (h *DashboardHandler) SearchByTypesense(c *fiber.Ctx) error {
	return nil
}

func (h *DashboardHandler) AllRecipes(c *fiber.Ctx) error {
	return nil
}

func (h *DashboardHandler) RecipeById(c *fiber.Ctx) error {
	return nil
}
