package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Prrromanssss/DAEE-fullstack/internal/storage"
	"github.com/Prrromanssss/DAEE-fullstack/internal/storage/postgres"
)

func HandlerGetOperations(w http.ResponseWriter, r *http.Request, dbCfg *storage.Storage) {
	operations, err := dbCfg.DB.GetOperations(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get operations: %v", err))
		return
	}
	respondWithJson(w, 200, postgres.DatabaseOperationsToOperations(operations))
}

func HandlerUpdateOperation(w http.ResponseWriter, r *http.Request, dbCfg *storage.Storage) {
	type parametrs struct {
		OperationType string `json:"operation_type"`
		ExecutionTime int32  `json:"execution_time"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parametrs{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
	}

	operation, err := dbCfg.DB.UpdateOperationTime(r.Context(), postgres.UpdateOperationTimeParams{
		OperationType: params.OperationType,
		ExecutionTime: params.ExecutionTime,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't update operation: %v", err))
		return
	}

	respondWithJson(w, 200, postgres.DatabaseOperationToOperation(operation))
}