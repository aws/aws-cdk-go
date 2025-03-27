//go:build no_runtime_type_checking

package awsecspatterns

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationLoadBalancedFargateService) validateAddServiceAsTargetParameters(service awsecs.BaseService) error {
	return nil
}

func (a *jsiiProxy_ApplicationLoadBalancedFargateService) validateCreateAWSLogDriverParameters(prefix *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationLoadBalancedFargateService) validateGetDefaultClusterParameters(scope constructs.Construct) error {
	return nil
}

func validateApplicationLoadBalancedFargateService_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewApplicationLoadBalancedFargateServiceParameters(scope constructs.Construct, id *string, props *ApplicationLoadBalancedFargateServiceProps) error {
	return nil
}

