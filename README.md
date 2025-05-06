# go_grpc_source_core
# hrm_hr_service

# generate protoc
protoc --go_out=. --go-grpc_out=. proto/hr.proto

sed -i 's|"hrm/|"github.com/huynhthanhthao/hrm_hr_service/|g' $(find . -name '*.go')