//go:build no_runtime_type_checking

package awsroute53targets

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LoadBalancerTarget) validateBindParameters(_record awsroute53.IRecordSet) error {
	return nil
}

func validateNewLoadBalancerTargetParameters(loadBalancer awselasticloadbalancingv2.ILoadBalancerV2) error {
	return nil
}

