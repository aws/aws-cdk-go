//go:build no_runtime_type_checking

package awsglobalacceleratorendpoints

// Building without runtime type checking enabled, so all the below just return nil

func validateNewApplicationLoadBalancerEndpointParameters(loadBalancer awselasticloadbalancingv2.IApplicationLoadBalancer, options *ApplicationLoadBalancerEndpointOptions) error {
	return nil
}

