//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsiotactions

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FirehosePutRecordAction) validateBindParameters(rule awsiot.ITopicRule) error {
	return nil
}

func validateNewFirehosePutRecordActionParameters(stream awskinesisfirehose.IDeliveryStream, props *FirehosePutRecordActionProps) error {
	return nil
}

