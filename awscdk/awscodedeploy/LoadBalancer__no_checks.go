//go:build no_runtime_type_checking

package awscodedeploy

// Building without runtime type checking enabled, so all the below just return nil

func validateLoadBalancer_ApplicationParameters(albTargetGroup awselasticloadbalancingv2.IApplicationTargetGroup) error {
	return nil
}

func validateLoadBalancer_ClassicParameters(loadBalancer awselasticloadbalancing.LoadBalancer) error {
	return nil
}

func validateLoadBalancer_NetworkParameters(nlbTargetGroup awselasticloadbalancingv2.INetworkTargetGroup) error {
	return nil
}

