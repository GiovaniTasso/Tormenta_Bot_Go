package models

type Aprimoramento struct {
	Custo     string `json:"custo"`
	Descricao string `json:"descrição"`
}

type Magia struct {
	Nome           string          `json:"nome"`
	Nivel          int             `json:"nivel"`
	Escola         string          `json:"escola"`
	Tipo           string          `json:"tipo"`
	Execucao       string          `json:"execução"`
	Alcance        string          `json:"alcance"`
	Alvo           string          `json:"alvo"`
	Duracao        string          `json:"duração"`
	Resistencia    string          `json:"resistencia"`
	Descricao      string          `json:"descrição"`
	Aprimoramentos []Aprimoramento `json:"aprimoramentos"`
}
