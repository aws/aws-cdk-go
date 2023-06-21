//go:build no_runtime_type_checking

package awscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_Metric) validateAttachToParameters(scope constructs.IConstruct) error {
	return nil
}

func (m *jsiiProxy_Metric) validateCreateAlarmParameters(scope constructs.Construct, id *string, props *CreateAlarmOptions) error {
	return nil
}

func (m *jsiiProxy_Metric) validateWithParameters(props *MetricOptions) error {
	return nil
}

func validateMetric_GrantPutMetricDataParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateNewMetricParameters(props *MetricProps) error {
	return nil
}

