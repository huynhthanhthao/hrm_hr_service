package hrGrpcServer

import (
	"context"
	"fmt"
	"github.com/huynhthanhthao/hrm_hr_service/ent"
	"github.com/huynhthanhthao/hrm_hr_service/ent/company"
	"github.com/huynhthanhthao/hrm_hr_service/generate"

	"github.com/google/uuid"
)

type HRServer struct {
	generate.ValidateCompanyRequest
	generate.UnimplementedValidateCompanyServiceServer
	client *ent.Client
}

func NewHRServer(client *ent.Client) *HRServer {
	return &HRServer{client: client}
}

func (s *HRServer) ValidateCompany(ctx context.Context, req *generate.ValidateCompanyRequest) (*generate.ValidateCompanyResponse, error) {
	// Convert int32 CompanyId to a UUID string if applicable
	companyID, err := uuid.FromBytes([]byte(fmt.Sprintf("%032d", req.CompanyId)))

	if err != nil {
		return &generate.ValidateCompanyResponse{Exists: false, Error: "Invalid Company ID"}, nil
	}

	exists, err := s.client.Company.
		Query().
		Where(company.ID(companyID)).
		Exist(ctx)

	if err != nil {
		return &generate.ValidateCompanyResponse{Exists: false, Error: err.Error()}, nil
	}

	return &generate.ValidateCompanyResponse{Exists: exists, Error: ""}, nil
}
