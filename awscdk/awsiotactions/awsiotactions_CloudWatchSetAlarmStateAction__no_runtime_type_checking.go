//go:build no_runtime_type_checking

package awsiotactions

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudWatchSetAlarmStateAction) validateBindParameters(topicRule awsiot.ITopicRule) error {
	return nil
}

func validateNewCloudWatchSetAlarmStateActionParameters(alarm awscloudwatch.IAlarm, props *CloudWatchSetAlarmStateActionProps) error {
	return nil
}

