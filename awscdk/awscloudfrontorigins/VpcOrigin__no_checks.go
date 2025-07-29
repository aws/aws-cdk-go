//go:build no_runtime_type_checking

package awscloudfrontorigins

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VpcOrigin) validateBindParameters(scope constructs.Construct, options *awscloudfront.OriginBindOptions) error {
	return nil
}

func validateVpcOrigin_WithApplicationLoadBalancerParameters(alb awselasticloadbalancingv2.IApplicationLoadBalancer, props *VpcOriginWithEndpointProps) error {
	return nil
}

func validateVpcOrigin_WithEc2InstanceParameters(instance awsec2.IInstance, props *VpcOriginWithEndpointProps) error {
	return nil
}

func validateVpcOrigin_WithNetworkLoadBalancerParameters(nlb awselasticloadbalancingv2.INetworkLoadBalancer, props *VpcOriginWithEndpointProps) error {
	return nil
}

func validateVpcOrigin_WithVpcOriginParameters(origin awscloudfront.IVpcOrigin, props *VpcOriginProps) error {
	return nil
}

func validateNewVpcOriginParameters(domainName *string, props *VpcOriginProps) error {
	return nil
}

