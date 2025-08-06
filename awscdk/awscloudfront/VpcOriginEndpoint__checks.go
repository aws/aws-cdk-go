//go:build !no_runtime_type_checking

package awscloudfront

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
)

func validateVpcOriginEndpoint_ApplicationLoadBalancerParameters(alb awselasticloadbalancingv2.IApplicationLoadBalancer) error {
	if alb == nil {
		return fmt.Errorf("parameter alb is required, but nil was provided")
	}

	return nil
}

func validateVpcOriginEndpoint_Ec2InstanceParameters(instance awsec2.IInstance) error {
	if instance == nil {
		return fmt.Errorf("parameter instance is required, but nil was provided")
	}

	return nil
}

func validateVpcOriginEndpoint_NetworkLoadBalancerParameters(nlb awselasticloadbalancingv2.INetworkLoadBalancer) error {
	if nlb == nil {
		return fmt.Errorf("parameter nlb is required, but nil was provided")
	}

	return nil
}

