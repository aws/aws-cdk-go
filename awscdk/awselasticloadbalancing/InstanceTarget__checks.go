//go:build !no_runtime_type_checking

package awselasticloadbalancing

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

func (i *jsiiProxy_InstanceTarget) validateAttachToClassicLBParameters(loadBalancer LoadBalancer) error {
	if loadBalancer == nil {
		return fmt.Errorf("parameter loadBalancer is required, but nil was provided")
	}

	return nil
}

func validateNewInstanceTargetParameters(instance awsec2.Instance) error {
	if instance == nil {
		return fmt.Errorf("parameter instance is required, but nil was provided")
	}

	return nil
}

