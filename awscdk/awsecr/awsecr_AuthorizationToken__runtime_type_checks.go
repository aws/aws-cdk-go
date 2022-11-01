//go:build !no_runtime_type_checking

package awsecr

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

func validateAuthorizationToken_GrantReadParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

