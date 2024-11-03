package model

import "time"

type Commission struct {
	ID                 int64     `db:"id" json:"id"`
	InsuranceID        int64     `db:"insurance_id" json:"insurance_id"`
	CommissionAmount   int64     `db:"commission_amount" json:"commission_amount"`
	ProductName        string    `db:"product_name" json:"product_name"`
	InsuranceAgentName string    `db:"insurance_agent_name" json:"insurance_agent_name"`
	ApplicantName      string    `db:"applicant_name" json:"applicant_name"`
	InsuredPersonName  string    `db:"insured_person_name" json:"insured_person_name"`
	CreatedAt          time.Time `db:"created_at" json:"created_at"`
}
