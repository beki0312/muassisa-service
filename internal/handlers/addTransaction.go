package handlers

//
//import (
//	"encoding/json"
//	"mintrans-integration/internal/models"
//	response "mintrans-integration/internal/models"
//	"net/http"
//)
//
//func (ch TransactionHandler) AddTransaction() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		apiKey := r.Header.Get("API-Key")
//		if apiKey != ch.svc.ConfigInstance().GetString("api.api-key.API-Key") {
//			response.ToJson(w, http.StatusUnauthorized, map[string]interface{}{
//				"response": "Неверный API ключ",
//			})
//			return
//		}
//		var transactionReq []*models.AddTransaction
//		var transaction models.AddTransaction
//		var newTransaction models.AddedNewTransaction
//
//		err := json.NewDecoder(r.Body).Decode(&newTransaction)
//		if err != nil {
//			ch.Logger.Error("request", response.SetError(err))
//			response.ToJson(w, http.StatusBadRequest, map[string]interface{}{
//				"response": err,
//			})
//			return
//		}
//		transaction.OrderId = newTransaction.OrderId
//		transaction.TransactionId = newTransaction.TransactionId
//		transaction.TerminalNumber = newTransaction.TerminalNumber
//		transaction.Amount = newTransaction.Amount
//		transaction.Status = newTransaction.Status
//		transaction.Type = newTransaction.Type
//		transaction.DateCreteCFT = newTransaction.DateCreteCFT
//		if newTransaction.Inn != ch.svc.ConfigInstance().GetString("api.merchant.merchant_inn") {
//			ch.Logger.Error("Платеж не принадлежить нужному мерчанту", newTransaction)
//			response.ToJson(w, http.StatusSeeOther, map[string]interface{}{
//				"response": "Платеж не принадлежить нужному мерчанту",
//			})
//			return
//		}
//		transactionReq = append(transactionReq, &transaction)
//		errr := ch.svc.TransactionRepositoryInstance().AddTransactionsPos(transactionReq)
//		if errr.ErrorCode != 0 {
//			ch.Logger.Error("response", errr)
//			response.ToJson(w, http.StatusBadRequest, map[string]interface{}{
//				"response": errr,
//			})
//			return
//		}
//		ch.Logger.Info("transaction:", transaction)
//		response.ToJson(w, http.StatusOK, map[string]interface{}{
//			"response": "success",
//		})
//	}
//}
