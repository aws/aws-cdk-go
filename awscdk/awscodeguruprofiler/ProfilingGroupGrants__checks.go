//go:build !no_runtime_type_checking

package awscodeguruprofiler

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscodeguruprofiler"
)

func (p *jsiiProxy_ProfilingGroupGrants) validatePublishParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_ProfilingGroupGrants) validateReadParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateProfilingGroupGrants_FromProfilingGroupParameters(resource interfacesawscodeguruprofiler.IProfilingGroupRef) error {
	if resource == nil {
		return fmt.Errorf("parameter resource is required, but nil was provided")
	}

	return nil
}

