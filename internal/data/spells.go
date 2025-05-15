package data

import (
	"TormentaBot/internal/models"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var magias []models.Magia

func LoadSpells() error {
	// Caminho corrigido (assumindo que o arquivo está em assets/magias.json)
	path := filepath.Join("assets", "magias.json")
	log.Printf("Carregando magias de: %s", path) // Log do caminho

	file, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Erro ao ler arquivo: %v", err) // Log de erro
		return err
	}

	if err := json.Unmarshal(file, &magias); err != nil {
		log.Printf("Erro ao decodificar JSON: %v", err) // Log de erro
		return err
	}

	log.Printf("Magias carregadas: %d", len(magias)) // Confirmação
	return nil
}

func SearchSpells(query string) []models.Magia {
	var results []models.Magia
	query = strings.ToLower(query)

	for _, magia := range magias {
		if strings.Contains(strings.ToLower(magia.Nome), query) {
			results = append(results, magia)
		}
	}
	return results
}

func GetSpellByName(name string) (models.Magia, bool) {
	for _, magia := range magias {
		if strings.EqualFold(magia.Nome, name) {
			return magia, true
		}
	}
	return models.Magia{}, false
}
