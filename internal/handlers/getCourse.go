package handlers

import (
	"encoding/json"
	response "muassisa-service/internal/models"
	"net/http"
)

type lang struct {
	Id int64 `json:"id"`
}

func (ch Handler) GetCourse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var language lang
		err := json.NewDecoder(r.Body).Decode(&language)
		if err != nil {
			ch.Logger.Error("request", response.SetError(err))
			response.ToJson(w, http.StatusBadRequest, map[string]interface{}{
				"response": err,
			})
			return
		}
		course, err := ch.svc.RepositoryInstance().GetCourse(language.Id)
		if err != nil {
			ch.Logger.Error("response", err)
			response.ToJson(w, http.StatusBadRequest, map[string]interface{}{
				"response": err,
			})
			return
		}
		ch.Logger.Info("course: ", course)
		response.ToJson(w, http.StatusOK, map[string]interface{}{
			"response": course,
		})
	}
}
