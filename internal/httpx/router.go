package httpx

import (
	"net/http"
	"strings"

	"Kate.com/notes-api/internal/httpx/handlers"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(h *handlers.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/docs/*", httpSwagger.WrapHandler)

	r.Route("/api/v1/notes", func(r chi.Router) {
		r.Post("/", h.CreateNote)      // CREATE
		r.Get("/", h.GetNotes)         // READ ALL
		r.Get("/{id}", h.GetNote)      // READ ONE
		r.Patch("/{id}", h.UpdateNote) // UPDATE

		// фэйк проверка токена
		// любой токен, начинающийся с Bearer
		r.Group(func(r chi.Router) {
			r.Use(func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					authHeader := r.Header.Get("Authorization")
					if !strings.HasPrefix(authHeader, "Bearer ") {
						http.Error(w, `{"error": "Authorization failed, use: Bearer <...>"}`, http.StatusUnauthorized)
						return
					}

					next.ServeHTTP(w, r)
				})
			})

			r.Delete("/{id}", h.DeleteNote) // DELETE
		})
	})

	return r
}
