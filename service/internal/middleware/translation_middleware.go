package middleware

import "net/http"

type TranslationMiddleware struct {
}

func NewTranslationMiddleware() *TranslationMiddleware {
	return &TranslationMiddleware{}
}

func (m *TranslationMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}
