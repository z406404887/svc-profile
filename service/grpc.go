// Code generated by protoc-gen-gomeet-service. DO NOT EDIT.
// source: pb/profile.proto
package service

import (
	"github.com/jinzhu/gorm"

	pb "github.com/gomeet-examples/svc-profile/pb"
	// SUB-SERVICES DEFINITION : import-pb
	// svc{{SubServiceNamePascalCase}}Pb "github.com/gomeet-examples/svc-{{SubServiceNameKebabCase}}/pb"
	// END SUB-SERVICES DEFINITION : import-pb
)

// Implements of profileServer
type profileServer struct {
	jwtSecret     string
	caCertificate string
	certificate   string
	privateKey    string

	// EXTRA : var
	mysqlDataSourceName string
	mysqlHandle         *gorm.DB
	// END EXTRA : var

	// SUB-SERVICES DEFINITION : var-address
	// svc{{SubServiceNamePascalCase}}Address string
	// END SUB-SERVICES DEFINITION : var-address

	// SUB-SERVICES DEFINITION : var-client
	// svc{{SubServiceNamePascalCase}}GrpcClient svc{{SubServiceNamePascalCase}}Pb.{{SubServiceNamePascalCase}}Client
	// END SUB-SERVICES DEFINITION : var-client
}

func newProfileServer() pb.ProfileServer {
	return new(profileServer)
}
