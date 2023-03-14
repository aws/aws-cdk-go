//go:build no_runtime_type_checking

package awsiotactions

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KinesisPutRecordAction) validateBindParameters(rule awsiot.ITopicRule) error {
	return nil
}

func validateNewKinesisPutRecordActionParameters(stream awskinesis.IStream, props *KinesisPutRecordActionProps) error {
	return nil
}

