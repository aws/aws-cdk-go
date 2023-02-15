//go:build no_runtime_type_checking

package awselasticloadbalancingv2

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationTargetGroup) validateAddLoadBalancerTargetParameters(props *LoadBalancerTargetProps) error {
	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateConfigureHealthCheckParameters(healthCheck *HealthCheck) error {
	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateEnableCookieStickinessParameters(duration awscdk.Duration) error {
	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateMetricHealthyHostCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateMetricHttpCodeTargetParameters(code HttpCodeTarget, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateMetricIpv6RequestCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateMetricRequestCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateMetricRequestCountPerTargetParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateMetricTargetConnectionErrorCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateMetricTargetResponseTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateMetricTargetTLSNegotiationErrorCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateMetricUnhealthyHostCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateRegisterConnectableParameters(connectable awsec2.IConnectable) error {
	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateRegisterListenerParameters(listener IApplicationListener) error {
	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateSetAttributeParameters(key *string) error {
	return nil
}

func validateApplicationTargetGroup_FromTargetGroupAttributesParameters(scope constructs.Construct, id *string, attrs *TargetGroupAttributes) error {
	return nil
}

func validateApplicationTargetGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationTargetGroup) validateSetHealthCheckParameters(val *HealthCheck) error {
	return nil
}

func validateNewApplicationTargetGroupParameters(scope constructs.Construct, id *string, props *ApplicationTargetGroupProps) error {
	return nil
}

