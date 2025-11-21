//go:build !no_runtime_type_checking

package awsecs

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsecs"
)

func (c *jsiiProxy_ClusterGrants) validateTaskProtectionParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateClusterGrants_FromClusterParameters(resource interfacesawsecs.IClusterRef) error {
	if resource == nil {
		return fmt.Errorf("parameter resource is required, but nil was provided")
	}

	return nil
}

