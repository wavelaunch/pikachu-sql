package ports

import "github.com/wavelaunch/pikachu-sql/api/internal/domain"

type DBSettingsRepository interface {
	Find(id string) (domain.DatabaseConnection, error)
	FindAll() ([]domain.DatabaseConnection, error)
}
