package middleware

import (
	"net/http"
	"task-management/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := utils.ExtractBearerToken(r)
		if tokenString == "" {
			utils.WriteErrorResponse(w, http.StatusUnauthorized, "Please, login to continue!")
			return
		}
		_, err := utils.VerifyToken(tokenString)

		if err != nil {
			utils.WriteErrorResponse(w, http.StatusUnauthorized, "Please, login to continue!")
			return
		}

		next.ServeHTTP(w, r)
	})
}
