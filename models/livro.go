package models

import "gorm.io/gorm"

type Livro struct {
	gorm.Model
	Titulo    string `json:"titulo"`
	Categoria string `json:"categoria"`
	Autor     string `json:"autor"`
	Sinopse   string `json:"sinopse"`
}
