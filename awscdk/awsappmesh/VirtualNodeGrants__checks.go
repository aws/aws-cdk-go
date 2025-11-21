//go:build !no_runtime_type_checking

package awsappmesh

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappmesh"
)

func (v *jsiiProxy_VirtualNodeGrants) validateStreamAggregatedResourcesParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateVirtualNodeGrants_FromVirtualNodeParameters(resource interfacesawsappmesh.IVirtualNodeRef) error {
	if resource == nil {
		return fmt.Errorf("parameter resource is required, but nil was provided")
	}

	return nil
}

