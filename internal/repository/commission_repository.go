package repository

import (
	"context"
	"edd_agent_commission/db/query"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CommissionRepository struct {
	pool    *pgxpool.Pool
	queries *query.Queries
}

func NewCommissionRepository(pool *pgxpool.Pool) *CommissionRepository {
	return &CommissionRepository{
		pool:    pool,
		queries: query.New(pool),
	}
}

func (r *CommissionRepository) GetCommissionsByAgentName(ctx context.Context, agentName string) ([]query.Commission, error) {
	return r.queries.GetCommissionsByAgentName(ctx, agentName)
}

func (r *CommissionRepository) InsertCommission(ctx context.Context, commission query.Commission) error {
	params := query.InsertCommissionParams{
		InsuranceID:        commission.InsuranceID,
		Commission:         commission.Commission,
		ProductName:        commission.ProductName,
		InsuranceAgentName: commission.InsuranceAgentName,
		ApplicantName:      commission.ApplicantName,
		InsuredPersonName:  commission.InsuredPersonName,
	}
	return r.queries.InsertCommission(ctx, params) // params 구조체 전달
}
