//go:build !no_runtime_type_checking

package awselasticloadbalancingv2targets

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
)

func (a *jsiiProxy_AlbTarget) validateAttachToNetworkTargetGroupParameters(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) error {
	if targetGroup == nil {
		return fmt.Errorf("parameter targetGroup is required, but nil was provided")
	}

	return nil
}

func validateNewAlbTargetParameters(alb awselasticloadbalancingv2.ApplicationLoadBalancer, port *float64) error {
	if alb == nil {
		return fmt.Errorf("parameter alb is required, but nil was provided")
	}

	if port == nil {
		return fmt.Errorf("parameter port is required, but nil was provided")
	}

	return nil
}

