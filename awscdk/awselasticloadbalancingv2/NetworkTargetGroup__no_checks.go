//go:build no_runtime_type_checking

package awselasticloadbalancingv2

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkTargetGroup) validateAddLoadBalancerTargetParameters(props *LoadBalancerTargetProps) error {
	return nil
}

func (n *jsiiProxy_NetworkTargetGroup) validateConfigureHealthCheckParameters(healthCheck *HealthCheck) error {
	return nil
}

func (n *jsiiProxy_NetworkTargetGroup) validateMetricHealthyHostCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NetworkTargetGroup) validateMetricUnHealthyHostCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NetworkTargetGroup) validateRegisterListenerParameters(listener INetworkListener) error {
	return nil
}

func (n *jsiiProxy_NetworkTargetGroup) validateSetAttributeParameters(key *string) error {
	return nil
}

func validateNetworkTargetGroup_FromTargetGroupAttributesParameters(scope constructs.Construct, id *string, attrs *TargetGroupAttributes) error {
	return nil
}

func validateNetworkTargetGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkTargetGroup) validateSetHealthCheckParameters(val *HealthCheck) error {
	return nil
}

func validateNewNetworkTargetGroupParameters(scope constructs.Construct, id *string, props *NetworkTargetGroupProps) error {
	return nil
}

