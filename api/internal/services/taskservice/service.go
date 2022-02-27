package taskservice

import "github.com/wavelaunch/pikachu-sql/api/internal/ports"

type service struct {
	dbSettingsRepository ports.DBSettingsRepository
}

func New(dbSettingsRepository ports.DBSettingsRepository) *service {
	return &service{
		dbSettingsRepository: dbSettingsRepository,
	}
}
