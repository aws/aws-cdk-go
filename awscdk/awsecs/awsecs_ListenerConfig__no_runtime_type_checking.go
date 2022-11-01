//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_ListenerConfig) validateAddTargetsParameters(id *string, target *LoadBalancerTargetOptions, service BaseService) error {
	return nil
}

func validateListenerConfig_ApplicationListenerParameters(listener awselasticloadbalancingv2.ApplicationListener, props *awselasticloadbalancingv2.AddApplicationTargetsProps) error {
	return nil
}

func validateListenerConfig_NetworkListenerParameters(listener awselasticloadbalancingv2.NetworkListener, props *awselasticloadbalancingv2.AddNetworkTargetsProps) error {
	return nil
}

