//go:build no_runtime_type_checking

package awscertificatemanager

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ICertificate) validateMetricDaysToExpiryParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

