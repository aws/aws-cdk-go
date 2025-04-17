//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IVpnConnection) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IVpnConnection) validateMetricTunnelDataInParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IVpnConnection) validateMetricTunnelDataOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IVpnConnection) validateMetricTunnelStateParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

