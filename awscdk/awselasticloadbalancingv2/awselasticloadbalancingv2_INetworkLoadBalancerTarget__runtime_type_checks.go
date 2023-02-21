//go:build !no_runtime_type_checking

package awselasticloadbalancingv2

import (
	"fmt"
)

func (i *jsiiProxy_INetworkLoadBalancerTarget) validateAttachToNetworkTargetGroupParameters(targetGroup INetworkTargetGroup) error {
	if targetGroup == nil {
		return fmt.Errorf("parameter targetGroup is required, but nil was provided")
	}

	return nil
}

