//go:build no_runtime_type_checking

package awscdkneptunealpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IDatabaseInstance) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

