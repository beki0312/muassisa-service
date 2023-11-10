package repository

import "muassisa-service/internal/models"

func (r repository) UpdateStatus(transId string) (errResponse models.ErrorResponse) {
	db := r.Postgres.GetPostgresConnection()
	sqlQuery := `update pos_terminal_transaction set status = 'COMPLETE' where transaction_id=?;`
	if err := db.Exec(sqlQuery, transId).Error; err != nil {
		return models.SetDbError(err)
	}
	return
}
