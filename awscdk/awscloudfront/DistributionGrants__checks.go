//go:build !no_runtime_type_checking

package awscloudfront

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudfront"
)

func (d *jsiiProxy_DistributionGrants) validateCreateInvalidationParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateDistributionGrants_FromDistributionParameters(resource interfacesawscloudfront.IDistributionRef) error {
	if resource == nil {
		return fmt.Errorf("parameter resource is required, but nil was provided")
	}

	return nil
}

