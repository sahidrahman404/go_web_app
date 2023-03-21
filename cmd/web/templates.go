package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/sahidrahman404/snippetbox/pkg/models"
)

type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

