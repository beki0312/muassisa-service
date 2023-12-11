package handlers

import (
	"github.com/gorilla/mux"
	response "muassisa-service/internal/models"
	"net/http"
	"strconv"
)

type lang struct {
	Id int64 `json:"id"`
}

func (ch Handler) GetCourse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)[`id`]
		num, _ := strconv.Atoi(id)

		course, err := ch.svc.RepositoryInstance().GetCourse(int64(num))
		if err != nil {
			//ch.Logger.Error("response", err)
			response.ToJson(w, http.StatusBadRequest, map[string]interface{}{
				"response": err,
			})
			return
		}
		//ch.Logger.Info("course: ", course)
		response.ToJson(w, http.StatusOK, map[string]interface{}{
			"response": course,
		})
	}
}
