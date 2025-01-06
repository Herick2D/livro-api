package models

import "gorm.io/gorm"

type Livro struct {
	gorm.Model
	Titulo    string `json:"titulo" binding:"required"`
	Categoria string `json:"categoria" biding:"required"`
	Autor     string `json:"autor" binding:"required"`
	Sinopse   string `json:"sinopse" binding:"required"`
}
