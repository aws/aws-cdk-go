//go:build no_runtime_type_checking

package awslogs

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ILogGroup) validateAddMetricFilterParameters(id *string, props *MetricFilterOptions) error {
	return nil
}

func (i *jsiiProxy_ILogGroup) validateAddStreamParameters(id *string, props *StreamOptions) error {
	return nil
}

func (i *jsiiProxy_ILogGroup) validateAddSubscriptionFilterParameters(id *string, props *SubscriptionFilterOptions) error {
	return nil
}

func (i *jsiiProxy_ILogGroup) validateExtractMetricParameters(jsonField *string, metricNamespace *string, metricName *string) error {
	return nil
}

func (i *jsiiProxy_ILogGroup) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_ILogGroup) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

