//go:build no_runtime_type_checking

package awsroute53targets

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClassicLoadBalancerTarget) validateBindParameters(_record awsroute53.IRecordSet) error {
	return nil
}

func validateNewClassicLoadBalancerTargetParameters(loadBalancer awselasticloadbalancing.LoadBalancer) error {
	return nil
}

