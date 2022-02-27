package dbsettingsrepository

import "github.com/wavelaunch/pikachu-sql/api/internal/domain"

type jsonConfigFileLoader struct {
	configFilePath string
}

func NewJsonConfigFileLoader(configFilePath string) *jsonConfigFileLoader {
	return &jsonConfigFileLoader{
		configFilePath: configFilePath,
	}
}

func (l *jsonConfigFileLoader) Find(id string) (domain.DatabaseConnection, error) {
	return domain.DatabaseConnection{}, nil
}

func (l *jsonConfigFileLoader) FindAll() ([]domain.DatabaseConnection, error) {
	return []domain.DatabaseConnection{}, nil
}
