//go:build !no_runtime_type_checking

package awscodedeploy

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancing"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
)

func validateLoadBalancer_ApplicationParameters(albTargetGroup awselasticloadbalancingv2.IApplicationTargetGroup) error {
	if albTargetGroup == nil {
		return fmt.Errorf("parameter albTargetGroup is required, but nil was provided")
	}

	return nil
}

func validateLoadBalancer_ClassicParameters(loadBalancer awselasticloadbalancing.LoadBalancer) error {
	if loadBalancer == nil {
		return fmt.Errorf("parameter loadBalancer is required, but nil was provided")
	}

	return nil
}

func validateLoadBalancer_NetworkParameters(nlbTargetGroup awselasticloadbalancingv2.INetworkTargetGroup) error {
	if nlbTargetGroup == nil {
		return fmt.Errorf("parameter nlbTargetGroup is required, but nil was provided")
	}

	return nil
}

