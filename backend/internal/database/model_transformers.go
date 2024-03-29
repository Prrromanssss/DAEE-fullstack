package database

import (
	"time"

	"github.com/google/uuid"
)

type ExpressionTransformed struct {
	ID        uuid.UUID        `json:"id"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	Data      string           `json:"data"`
	Status    ExpressionStatus `json:"status"`
	ParseData string           `json:"parse_data"`
	Result    int32            `json:"result"`
	IsReady   bool             `json:"is_ready"`
}

func DatabaseExpressionToExpression(dbExpr Expression) ExpressionTransformed {
	return ExpressionTransformed{
		ID:        dbExpr.ID,
		CreatedAt: dbExpr.CreatedAt,
		UpdatedAt: dbExpr.UpdatedAt,
		Data:      dbExpr.Data,
		Status:    dbExpr.Status,
		ParseData: dbExpr.ParseData,
		Result:    dbExpr.Result,
		IsReady:   dbExpr.IsReady,
	}
}

func DatabaseExpressionsToExpressions(dbExprs []Expression) []ExpressionTransformed {
	exprs := []ExpressionTransformed{}
	for _, dbExpr := range dbExprs {
		exprs = append(exprs, DatabaseExpressionToExpression(dbExpr))
	}
	return exprs
}

type OperationTransformed struct {
	ID            uuid.UUID `json:"id"`
	OperationType string    `json:"operation_type"`
	ExecutionTime int32     `json:"execution_time"`
}

func DatabaseOperationToOperation(dbOper Operation) OperationTransformed {
	return OperationTransformed(dbOper)
}

func DatabaseOperationsToOperations(dbOpers []Operation) []OperationTransformed {
	opers := []OperationTransformed{}
	for _, dbOper := range dbOpers {
		opers = append(opers, DatabaseOperationToOperation(dbOper))
	}
	return opers
}

type AgentTransformed struct {
	ID                           uuid.UUID   `json:"id"`
	NumberOfParallelCalculations int32       `json:"number_of_parallel_calculations"`
	LastPing                     time.Time   `json:"last_ping"`
	Status                       AgentStatus `json:"status"`
	CreatedAt                    time.Time   `json:"created_at"`
}

func DatabaseAgentToAgent(dbAgent Agent) AgentTransformed {
	return AgentTransformed(dbAgent)
}

func DatabaseAgentsToAgents(dbAgents []Agent) []AgentTransformed {
	agents := []AgentTransformed{}
	for _, dbAgent := range dbAgents {
		agents = append(agents, DatabaseAgentToAgent(dbAgent))
	}
	return agents
}
