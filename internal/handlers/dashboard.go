package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/qwaq-dev/culina/internal/repository"
)

type DashboardHandler struct {
	repo repository.DashboardRepository
	log  *slog.Logger
}

func NewDashboardHandler(repo repository.DashboardRepository, log *slog.Logger) *DashboardHandler {
	return &DashboardHandler{repo: repo, log: log}
}

func (h *DashboardHandler) CreateRecipe(c *fiber.Ctx) error {
	return nil
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
