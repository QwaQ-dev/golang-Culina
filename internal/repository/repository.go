package repository

import (
	"log/slog"

	"github.com/qwaq-dev/culina/structures"
)

type UserRepository interface {
	CreateUser(user *structures.User, log *slog.Logger) (int, error)
	GetUser(username string, log *slog.Logger) (*structures.User, error)
}

type DashboardRepository interface {
}

type ProfileRepository interface {
	ChangeProfileData(column, newData, userId string, log *slog.Logger) (*structures.User, error)
	GetUserRecipes(user *structures.User, log *slog.Logger) (*structures.Recipes, error)
}
