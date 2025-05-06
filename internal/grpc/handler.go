package hrGrpcServer

import (
	"context"

	"github.com/google/uuid"
	"github.com/huynhthanhthao/hrm_hr_service/ent"
	"github.com/huynhthanhthao/hrm_hr_service/ent/company"
	"github.com/huynhthanhthao/hrm_hr_service/generate"
)

type HRServer struct {
	generate.UnimplementedValidateCompanyServiceServer
	client *ent.Client
}

func NewHRServer(client *ent.Client) *HRServer {
	return &HRServer{client: client}
}

func (s *HRServer) ValidateCompany(ctx context.Context, req *generate.ValidateCompanyRequest) (*generate.ValidateCompanyResponse, error) {
	// Parse company_id thành UUID
	companyID, err := uuid.Parse(req.CompanyId)
	if err != nil {
		return &generate.ValidateCompanyResponse{
			Exists: false,
			Error:  "Invalid Company ID format",
		}, nil
	}

	// Kiểm tra sự tồn tại của company_id trong database
	exists, err := s.client.Company.
		Query().
		Where(company.ID(companyID)).
		Exist(ctx)

	if err != nil {
		return &generate.ValidateCompanyResponse{
			Exists: false,
			Error:  err.Error(),
		}, nil
	}

	return &generate.ValidateCompanyResponse{
		Exists: exists,
		Error:  "",
	}, nil
}
