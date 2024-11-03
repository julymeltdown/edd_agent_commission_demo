package service

import (
	"context"
	"edd_agent_commission/db/query"
	"edd_agent_commission/internal/event"
	"edd_agent_commission/internal/repository"
)

type CommissionService struct {
	repo *repository.CommissionRepository
}

func NewCommissionService(repo *repository.CommissionRepository) *CommissionService {
	return &CommissionService{repo: repo}
}

func (s *CommissionService) GetCommissionsByAgentName(ctx context.Context, agentName string) ([]query.Commission, error) {
	return s.repo.GetCommissionsByAgentName(ctx, agentName)
}

func (s *CommissionService) ProcessCommission(ctx context.Context, event event.InsuranceApplicationAcceptedEvent) error {
	commission := query.Commission{
		InsuranceID:        event.InsuranceId,
		Commission:         event.Commission,
		ProductName:        event.ProductName,
		InsuranceAgentName: event.InsuranceAgentName,
		ApplicantName:      event.ApplicantName,
		InsuredPersonName:  event.InsuredPersonName,
	}
	return s.repo.InsertCommission(ctx, commission)
}

// ... (생략) ... - 이벤트 구조체 정의
