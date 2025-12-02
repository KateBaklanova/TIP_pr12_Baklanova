// Package main Notes API server.
//
// @title           Notes API
// @version         1.0
// @description     Учебный REST API для заметок (CRUD).
// @contact.name    Backend Course
// @contact.email   example@university.ru
// @BasePath        /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Введите токен: Bearer <token>

package main

import (
	"log"
	"net/http"

	_ "Kate.com/notes-api/docs"
	"Kate.com/notes-api/internal/httpx"
	"Kate.com/notes-api/internal/httpx/handlers"
	"Kate.com/notes-api/internal/repo"
)

func main() {
	repo := repo.NewNoteRepoMem()
	h := &handlers.Handler{Repo: repo}
	r := httpx.NewRouter(h)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
