//go:build !no_runtime_type_checking

package awselasticloadbalancingv2targets

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
)

func (a *jsiiProxy_AlbListenerTarget) validateAttachToNetworkTargetGroupParameters(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) error {
	if targetGroup == nil {
		return fmt.Errorf("parameter targetGroup is required, but nil was provided")
	}

	return nil
}

func validateNewAlbListenerTargetParameters(albListener awselasticloadbalancingv2.ApplicationListener) error {
	if albListener == nil {
		return fmt.Errorf("parameter albListener is required, but nil was provided")
	}

	return nil
}

