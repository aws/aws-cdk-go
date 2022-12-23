//go:build no_runtime_type_checking

package awsiotactions

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IotRepublishMqttAction) validateBindParameters(rule awsiot.ITopicRule) error {
	return nil
}

func validateNewIotRepublishMqttActionParameters(topic *string, props *IotRepublishMqttActionProps) error {
	return nil
}

