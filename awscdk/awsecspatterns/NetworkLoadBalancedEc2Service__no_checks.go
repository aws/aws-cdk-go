//go:build no_runtime_type_checking

package awsecspatterns

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkLoadBalancedEc2Service) validateAddServiceAsTargetParameters(service awsecs.BaseService) error {
	return nil
}

func (n *jsiiProxy_NetworkLoadBalancedEc2Service) validateCreateAWSLogDriverParameters(prefix *string) error {
	return nil
}

func (n *jsiiProxy_NetworkLoadBalancedEc2Service) validateGetDefaultClusterParameters(scope constructs.Construct) error {
	return nil
}

func validateNetworkLoadBalancedEc2Service_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewNetworkLoadBalancedEc2ServiceParameters(scope constructs.Construct, id *string, props *NetworkLoadBalancedEc2ServiceProps) error {
	return nil
}

