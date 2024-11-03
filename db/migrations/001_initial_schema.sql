CREATE TABLE commissions (
                             id SERIAL PRIMARY KEY,
                             insurance_id BIGINT NOT NULL,
                             commission BIGINT NOT NULL,
                             product_name TEXT NOT NULL,
                             insurance_agent_name TEXT NOT NULL,
                             applicant_name TEXT NOT NULL,
                             insured_person_name TEXT NOT NULL
);