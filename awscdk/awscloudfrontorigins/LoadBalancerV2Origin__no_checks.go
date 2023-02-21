//go:build no_runtime_type_checking

package awscloudfrontorigins

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LoadBalancerV2Origin) validateBindParameters(_scope constructs.Construct, options *awscloudfront.OriginBindOptions) error {
	return nil
}

func validateNewLoadBalancerV2OriginParameters(loadBalancer awselasticloadbalancingv2.ILoadBalancerV2, props *LoadBalancerV2OriginProps) error {
	return nil
}

