//go:build no_runtime_type_checking

package awsiotactions

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudWatchPutMetricAction) validateBindParameters(rule awsiot.ITopicRule) error {
	return nil
}

func validateNewCloudWatchPutMetricActionParameters(props *CloudWatchPutMetricActionProps) error {
	return nil
}

