package repository

import (
	"gorm.io/gorm/clause"
	"muassisa-service/internal/models"
)

func (r repository) AddTransactionsPos(transactions []*models.AddTransaction) (errResponse models.ErrorResponse) {

	db := r.Postgres.GetPostgresConnection()
	err := db.Table("pos_terminal_transaction").Clauses(clause.OnConflict{DoNothing: true}).Create(transactions)
	if err.Error != nil {
		errResponse = models.SetDbError(err.Error)
		return errResponse
	}
	return
}
