//go:build no_runtime_type_checking

package awsecspatterns

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationLoadBalancedEc2Service) validateAddServiceAsTargetParameters(service awsecs.BaseService) error {
	return nil
}

func (a *jsiiProxy_ApplicationLoadBalancedEc2Service) validateCreateAWSLogDriverParameters(prefix *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationLoadBalancedEc2Service) validateGetDefaultClusterParameters(scope constructs.Construct) error {
	return nil
}

func validateApplicationLoadBalancedEc2Service_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewApplicationLoadBalancedEc2ServiceParameters(scope constructs.Construct, id *string, props *ApplicationLoadBalancedEc2ServiceProps) error {
	return nil
}

