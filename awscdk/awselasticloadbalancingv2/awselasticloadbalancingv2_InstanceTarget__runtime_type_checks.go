//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awselasticloadbalancingv2

import (
	"fmt"
)

func (i *jsiiProxy_InstanceTarget) validateAttachToApplicationTargetGroupParameters(targetGroup IApplicationTargetGroup) error {
	if targetGroup == nil {
		return fmt.Errorf("parameter targetGroup is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_InstanceTarget) validateAttachToNetworkTargetGroupParameters(targetGroup INetworkTargetGroup) error {
	if targetGroup == nil {
		return fmt.Errorf("parameter targetGroup is required, but nil was provided")
	}

	return nil
}

func validateNewInstanceTargetParameters(instanceId *string) error {
	if instanceId == nil {
		return fmt.Errorf("parameter instanceId is required, but nil was provided")
	}

	return nil
}

