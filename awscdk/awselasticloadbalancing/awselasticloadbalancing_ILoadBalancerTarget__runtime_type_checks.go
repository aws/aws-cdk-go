//go:build !no_runtime_type_checking

package awselasticloadbalancing

import (
	"fmt"
)

func (i *jsiiProxy_ILoadBalancerTarget) validateAttachToClassicLBParameters(loadBalancer LoadBalancer) error {
	if loadBalancer == nil {
		return fmt.Errorf("parameter loadBalancer is required, but nil was provided")
	}

	return nil
}

