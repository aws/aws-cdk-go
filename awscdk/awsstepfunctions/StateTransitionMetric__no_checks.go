//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func validateStateTransitionMetric_MetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateStateTransitionMetric_MetricConsumedCapacityParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateStateTransitionMetric_MetricProvisionedBucketSizeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateStateTransitionMetric_MetricProvisionedRefillRateParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateStateTransitionMetric_MetricThrottledEventsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

