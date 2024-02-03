package remove

import (
	"log/slog"
	"net/http"
	"url-shortener/internal/lib/api/response"
	"url-shortener/internal/lib/logger/sl"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type URLRemover interface {
	RemoveURL(alias string) error
}

func New(log *slog.Logger, urlRemover URLRemover) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.remove.New"

		log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("empty alias")

			render.JSON(w, r, response.Error("empty alias"))

			return
		}

		err := urlRemover.RemoveURL(alias)
		if err != nil {
			log.Info("error deleting url", sl.Err(err))

			render.JSON(w, r, response.Error("internal error"))

			return
		}

		log.Info("url deleted", slog.String("alias", alias))
		render.JSON(w, r, response.OK())
	}
}