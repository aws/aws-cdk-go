//go:build !no_runtime_type_checking

package awscdkgluealpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

func validateDataCatalogEncryptionAtRest_KmsWithServiceRoleParameters(role awsiam.IRole) error {
	if role == nil {
		return fmt.Errorf("parameter role is required, but nil was provided")
	}

	return nil
}

