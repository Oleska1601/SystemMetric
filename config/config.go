package config

import "github.com/ilyakaznacheev/cleanenv"

type (
	Config struct {
		App      `yaml:"app"`
		HTTP     `yaml:"http"`
		Logger   `yaml:"logger"`
		Web      `yaml:"web"`
		Postgres `yaml:"postgres"`
	}

	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}

	HTTP struct {
		Port string `yaml:"port"`
	}

	Logger struct {
		Level string `yaml:"level"`
	}

	Web struct {
		Path string `yaml:"path"`
	}

	Postgres struct {
		PoolMax int    `yaml:"pool_max"`
		PgUrl   string `yaml:"pg_url"`
	}
)

func New() (*Config, error) {
	cfg := &Config{}
	//читаем конфигурацию из файла
	//`cleanenv` автоматически сопоставит значения из YAML-файла полям структуры
	err := cleanenv.ReadConfig("config/config.yml", cfg)
	if err != nil {
		return nil, err
	}
	//переопределить значения полей структуры `cfg` значениями из переменных окружения.
	//Переменные окружения имеют приоритет над значениями
	//затем переопределять отдельные значения с помощью переменных окружения
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
