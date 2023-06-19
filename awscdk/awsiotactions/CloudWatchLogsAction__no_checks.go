//go:build no_runtime_type_checking

package awsiotactions

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudWatchLogsAction) validateBindParameters(rule awsiot.ITopicRule) error {
	return nil
}

func validateNewCloudWatchLogsActionParameters(logGroup awslogs.ILogGroup, props *CloudWatchLogsActionProps) error {
	return nil
}

