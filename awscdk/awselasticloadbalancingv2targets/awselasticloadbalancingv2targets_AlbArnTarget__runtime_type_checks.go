//go:build !no_runtime_type_checking

package awselasticloadbalancingv2targets

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
)

func (a *jsiiProxy_AlbArnTarget) validateAttachToNetworkTargetGroupParameters(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) error {
	if targetGroup == nil {
		return fmt.Errorf("parameter targetGroup is required, but nil was provided")
	}

	return nil
}

func validateNewAlbArnTargetParameters(albArn *string, port *float64) error {
	if albArn == nil {
		return fmt.Errorf("parameter albArn is required, but nil was provided")
	}

	if port == nil {
		return fmt.Errorf("parameter port is required, but nil was provided")
	}

	return nil
}

