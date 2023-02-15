//go:build no_runtime_type_checking

package awsecspatterns

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationLoadBalancedServiceBase) validateAddServiceAsTargetParameters(service awsecs.BaseService) error {
	return nil
}

func (a *jsiiProxy_ApplicationLoadBalancedServiceBase) validateCreateAWSLogDriverParameters(prefix *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationLoadBalancedServiceBase) validateGetDefaultClusterParameters(scope constructs.Construct) error {
	return nil
}

func validateApplicationLoadBalancedServiceBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewApplicationLoadBalancedServiceBaseParameters(scope constructs.Construct, id *string, props *ApplicationLoadBalancedServiceBaseProps) error {
	return nil
}

