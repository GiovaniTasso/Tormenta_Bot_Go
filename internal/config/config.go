package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	TelegramToken string `mapstructure:"telegram_token"`
	DataPath      string `mapstructure:"data_path"`
	Debug         bool   `mapstructure:"debug"`
}

func Load() *Config {
	viper.SetConfigName("config") // Nome do arquivo (config.yaml)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // Procura o arquivo no diretório raiz

	// Tenta ler o arquivo
	if err := viper.ReadInConfig(); err != nil {
		panic("Erro ao ler o arquivo de configuração: " + err.Error())
	}

	// Carrega variáveis de ambiente (opcional)
	viper.AutomaticEnv()

	// Mapeia para a struct
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic("Erro ao decodificar configurações: " + err.Error())
	}

	return &cfg
}
