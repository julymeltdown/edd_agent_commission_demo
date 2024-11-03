package event

type InsuranceApplicationAcceptedEvent struct {
	InsuranceId        int64  `json:"insuranceId"`
	Commission         int64  `json:"commission"`
	ProductName        string `json:"productName"`
	InsuranceAgentName string `json:"insuranceAgentName"`
	ApplicantName      string `json:"applicantName"`
	InsuredPersonName  string `json:"insuredPersonName"`
}
