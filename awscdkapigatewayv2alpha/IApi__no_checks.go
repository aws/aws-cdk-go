//go:build no_runtime_type_checking

package awscdkapigatewayv2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IApi) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

