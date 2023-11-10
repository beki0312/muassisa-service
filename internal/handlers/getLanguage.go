package handlers

import (
	response "muassisa-service/internal/models"
	"net/http"
)

func (ch Handler) GetLanguage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		language, err := ch.svc.RepositoryInstance().GeLanguage()
		if err != nil {
			ch.Logger.Error("response", err)
			response.ToJson(w, http.StatusBadRequest, map[string]interface{}{
				"response": err,
			})
			return
		}
		ch.Logger.Info("language: ", language)
		response.ToJson(w, http.StatusOK, map[string]interface{}{
			"response": language,
		})
	}
}
