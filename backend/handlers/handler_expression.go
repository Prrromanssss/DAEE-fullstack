package handlers

import (
	"Prrromanssss/DAEE/config"
	"Prrromanssss/DAEE/internal/database"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func HandlerCreateExpression(w http.ResponseWriter, r *http.Request, apiCfg *config.ApiConfig) {
	type parametrs struct {
		Data string `json:"data"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parametrs{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
	}

	expression, err := apiCfg.DB.CreateExpression(r.Context(),
		database.CreateExpressionParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Data:      params.Data,
			ParseData: "",
			Status:    "ready for computation",
		})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create expression: %v", err))
		return
	}

	respondWithJson(w, 201, database.DatabaseExpressionToExpression(expression))
}

func HandlerGetExpressionByID(w http.ResponseWriter, r *http.Request, apiCfg *config.ApiConfig) {
	expressionIDString := chi.URLParam(r, "expressionID")
	expressionID, err := uuid.Parse(expressionIDString)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse expression id: %v", err))
		return
	}
	expression, err := apiCfg.DB.GetExpressionByID(r.Context(), expressionID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get expression: %v", err))
		return
	}
	respondWithJson(w, 200, database.DatabaseExpressionToExpression(expression))
}

func HandlerGetExpressions(w http.ResponseWriter, r *http.Request, apiCfg *config.ApiConfig) {
	expressions, err := apiCfg.DB.GetExpressions(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get expressions: %v", err))
		return
	}
	respondWithJson(w, 200, database.DatabaseExpressionsToExpressions(expressions))
}