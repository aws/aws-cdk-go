//go:build no_runtime_type_checking

package awscloudfront

// Building without runtime type checking enabled, so all the below just return nil

func validateVpcOriginEndpoint_ApplicationLoadBalancerParameters(alb awselasticloadbalancingv2.IApplicationLoadBalancer) error {
	return nil
}

func validateVpcOriginEndpoint_Ec2InstanceParameters(instance awsec2.IInstance) error {
	return nil
}

func validateVpcOriginEndpoint_NetworkLoadBalancerParameters(nlb awselasticloadbalancingv2.INetworkLoadBalancer) error {
	return nil
}

