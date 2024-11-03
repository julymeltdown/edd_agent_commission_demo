-- name: GetCommissionsByAgentName :many
SELECT * FROM commissions WHERE insurance_agent_name = $1;

-- name: InsertCommission :exec
INSERT INTO commissions (insurance_id, commission, product_name, insurance_agent_name, applicant_name, insured_person_name)
VALUES ($1, $2, $3, $4, $5, $6);