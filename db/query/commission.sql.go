// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: commission.sql

package query

import (
	"context"
)

const getCommissionsByAgentName = `-- name: GetCommissionsByAgentName :many
    SELECT id, insurance_id, commission, product_name, insurance_agent_name, applicant_name, insured_person_name FROM commissions WHERE insurance_agent_name = $1
`

func (q *Queries) GetCommissionsByAgentName(ctx context.Context, insuranceAgentName string) ([]Commission, error) {
	rows, err := q.db.Query(ctx, getCommissionsByAgentName, insuranceAgentName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Commission
	for rows.Next() {
		var i Commission
		if err := rows.Scan(
			&i.ID,
			&i.InsuranceID,
			&i.Commission,
			&i.ProductName,
			&i.InsuranceAgentName,
			&i.ApplicantName,
			&i.InsuredPersonName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertCommission = `-- name: InsertCommission :exec
    INSERT INTO commissions (insurance_id, commission, product_name, insurance_agent_name, applicant_name, insured_person_name)
    VALUES ($1, $2, $3, $4, $5, $6)
`

type InsertCommissionParams struct {
	InsuranceID        int64
	Commission         int64
	ProductName        string
	InsuranceAgentName string
	ApplicantName      string
	InsuredPersonName  string
}

func (q *Queries) InsertCommission(ctx context.Context, arg InsertCommissionParams) error {
	_, err := q.db.Exec(ctx, insertCommission,
		arg.InsuranceID,
		arg.Commission,
		arg.ProductName,
		arg.InsuranceAgentName,
		arg.ApplicantName,
		arg.InsuredPersonName,
	)
	return err
}
