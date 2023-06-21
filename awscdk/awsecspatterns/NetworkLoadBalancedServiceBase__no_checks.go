//go:build no_runtime_type_checking

package awsecspatterns

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkLoadBalancedServiceBase) validateAddServiceAsTargetParameters(service awsecs.BaseService) error {
	return nil
}

func (n *jsiiProxy_NetworkLoadBalancedServiceBase) validateCreateAWSLogDriverParameters(prefix *string) error {
	return nil
}

func (n *jsiiProxy_NetworkLoadBalancedServiceBase) validateGetDefaultClusterParameters(scope constructs.Construct) error {
	return nil
}

func validateNetworkLoadBalancedServiceBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewNetworkLoadBalancedServiceBaseParameters(scope constructs.Construct, id *string, props *NetworkLoadBalancedServiceBaseProps) error {
	return nil
}

