//go:build no_runtime_type_checking

package awsglobalacceleratorendpoints

// Building without runtime type checking enabled, so all the below just return nil

func validateNewNetworkLoadBalancerEndpointParameters(loadBalancer awselasticloadbalancingv2.INetworkLoadBalancer, options *NetworkLoadBalancerEndpointProps) error {
	return nil
}

