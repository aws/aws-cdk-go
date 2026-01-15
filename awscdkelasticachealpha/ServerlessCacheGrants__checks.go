//go:build !no_runtime_type_checking

package awscdkelasticachealpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawselasticache"
)

func (s *jsiiProxy_ServerlessCacheGrants) validateConnectParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateServerlessCacheGrants_FromServerlessCacheParameters(resource interfacesawselasticache.IServerlessCacheRef) error {
	if resource == nil {
		return fmt.Errorf("parameter resource is required, but nil was provided")
	}

	return nil
}

