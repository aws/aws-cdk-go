//go:build no_runtime_type_checking

package awsecspatterns

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkLoadBalancedFargateService) validateAddServiceAsTargetParameters(service awsecs.BaseService) error {
	return nil
}

func (n *jsiiProxy_NetworkLoadBalancedFargateService) validateCreateAWSLogDriverParameters(prefix *string) error {
	return nil
}

func (n *jsiiProxy_NetworkLoadBalancedFargateService) validateGetDefaultClusterParameters(scope awscdk.Construct) error {
	return nil
}

func (n *jsiiProxy_NetworkLoadBalancedFargateService) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (n *jsiiProxy_NetworkLoadBalancedFargateService) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateNetworkLoadBalancedFargateService_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewNetworkLoadBalancedFargateServiceParameters(scope constructs.Construct, id *string, props *NetworkLoadBalancedFargateServiceProps) error {
	return nil
}

